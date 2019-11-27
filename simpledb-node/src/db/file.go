package db

import (
	"encoding/binary"
	"os"
)

func writeNewFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	numBytes, err := f.Write(data)
	if err != nil {
		return err
	}
	if numBytes != len(data) {
		return ErrWriteUnexpectedBytes(filename)
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}

func mmap(filename string) (entries []*LSMDataEntry, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	dataSize, _, _, _, err := readHeader(f)
	if err != nil {
		return nil, err
	}
	data := make([]byte, dataSize)
	numBytes, err := f.ReadAt(data, headerSize)
	if err != nil {
		return nil, err
	}
	if numBytes != len(data) {
		return nil, ErrReadUnexpectedBytes("SST File")
	}
	for i := 0; i < len(data); i += BlockSize {
		block := data[i : i+BlockSize]
		j := 0
		for j+8 <= len(block) {
			seqID := binary.LittleEndian.Uint64(block[j : j+8])
			j += 8
			if j >= len(block) {
				break
			}
			keySize := uint8(block[j])
			j++
			if j+int(keySize)-1 >= len(block) {
				break
			}
			key := string(block[j : j+int(keySize)])
			j += int(keySize)
			if j >= len(block) {
				break
			}
			valueType := uint8(block[j])
			j++
			if j+1 >= len(block) {
				break
			}
			valueSize := binary.LittleEndian.Uint16(block[j : j+2])
			j += 2
			if j+int(valueSize)-1 >= len(block) {
				break
			}
			value := block[j : j+int(valueSize)]
			j += int(valueSize)

			if key != "" {
				entries = append(entries, &LSMDataEntry{
					seqID:     seqID,
					keySize:   keySize,
					key:       key,
					valueType: valueType,
					valueSize: valueSize,
					value:     value,
				})
			}
		}
	}
	return entries, nil
}

// recoverFile reads a file and returns key range, bloom filter, and total size of the file
func recoverFile(filename string) (keyRange *KeyRange, bloom *Bloom, size int, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer f.Close()

	if err != nil {
		return nil, nil, 0, err
	}

	dataSize, indexSize, bloomSize, keyRangeSize, err := readHeader(f)
	if err != nil {
		return nil, nil, 0, err
	}

	bitsAndKeyRange := make([]byte, bloomSize+keyRangeSize)

	numBytes, err := f.ReadAt(bitsAndKeyRange, int64(headerSize+dataSize+indexSize))
	if err != nil {
		return nil, nil, 0, err
	}
	if numBytes != len(bitsAndKeyRange) {
		return nil, nil, 0, ErrWriteUnexpectedBytes(filename)
	}

	bits := bitsAndKeyRange[:bloomSize]
	keyRangeBytes := bitsAndKeyRange[bloomSize:]

	bloom = RecoverBloom(bits)
	keyRange = parseKeyRangeEntry(keyRangeBytes)

	return keyRange, bloom, int(dataSize + indexSize + bloomSize + keyRangeSize), nil
}

func fileFind(filename, key string) (entry *LSMDataEntry, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	dataSize, indexSize, _, _, err := readHeader(f)
	if err != nil {
		return nil, err
	}

	index := make([]byte, indexSize)
	numBytes, err := f.ReadAt(index, headerSize+int64(dataSize))
	if err != nil {
		return nil, err
	}
	if numBytes != int(indexSize) {
		return nil, ErrReadUnexpectedBytes("SST File, Index Block")
	}

	blockIndex, err := findDataBlock(key, index)
	if err != nil {
		return nil, err
	}

	block := make([]byte, BlockSize)
	numBytes, err = f.ReadAt(block, headerSize+int64(BlockSize*int(blockIndex)))
	if err != nil {
		return nil, err
	}
	if numBytes != BlockSize {
		return nil, ErrReadUnexpectedBytes("SST File, Data Block")
	}

	entry, err = findKeyInBlock(key, block)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func fileRange(filename string, keyRange *KeyRange) (entries []*LSMDataEntry, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	dataSize, indexSize, _, _, err := readHeader(f)
	if err != nil {
		return nil, err
	}

	index := make([]byte, indexSize)
	numBytes, err := f.ReadAt(index, headerSize+int64(dataSize))
	if err != nil {
		return nil, err
	}
	if numBytes != int(indexSize) {
		return nil, ErrReadUnexpectedBytes("SST File, Index Block")
	}

	startBlock, endBlock := rangeDataBlocks(keyRange.startKey, keyRange.endKey, index)
	size := int(endBlock-startBlock+1) * BlockSize
	blocks := make([]byte, size)

	numBytes, err = f.ReadAt(blocks, headerSize+int64(BlockSize*int(startBlock)))
	if err != nil {
		return nil, err
	}
	if numBytes != size {
		return nil, ErrReadUnexpectedBytes("SST File, Data Block")
	}

	entries = findKeysInBlocks(keyRange.startKey, keyRange.endKey, blocks)
	return entries, nil
}

func findDataBlock(key string, index []byte) (uint32, error) {
	i := 0
	for i < len(index) {
		size := uint8(index[i])
		i++
		if i+int(size)-1 >= len(index) {
			break
		}
		indexKey := string(index[i : i+int(size)])
		i += int(size)
		if i+3 >= len(index) {
			break
		}
		if key <= indexKey {
			return binary.LittleEndian.Uint32(index[i : i+4]), nil
		}
		i += 4
	}
	return 0, ErrKeyNotFound()
}

func findKeyInBlock(key string, block []byte) (*LSMDataEntry, error) {
	i := 0
	for i+8 <= len(block) {
		seqID := binary.LittleEndian.Uint64(block[i : i+8])
		i += 8
		if i >= len(block) {
			break
		}
		keySize := uint8(block[i])
		i++
		if i+int(keySize)-1 >= len(block) {
			break
		}
		foundKey := string(block[i : i+int(keySize)])
		i += int(keySize)
		if i >= len(block) {
			break
		}
		valueType := uint8(block[i])
		i++
		if i+1 >= len(block) {
			break
		}
		valueSize := binary.LittleEndian.Uint16(block[i : i+2])
		i += 2
		if i+int(valueSize)-1 >= len(block) {
			break
		}
		value := block[i : i+int(valueSize)]
		i += int(valueSize)

		if key == foundKey {
			return &LSMDataEntry{
				seqID:     seqID,
				keySize:   keySize,
				key:       key,
				valueType: valueType,
				valueSize: valueSize,
				value:     value,
			}, nil
		}
	}
	return nil, ErrKeyNotFound()
}

func rangeDataBlocks(startKey, endKey string, index []byte) (startBlock, endBlock uint32) {
	foundStartBlock := false
	block := uint32(0)
	i := 0
	for i < len(index) {
		size := uint8(index[i])
		i++
		if i+int(size)-1 >= len(index) {
			break
		}
		indexKey := string(index[i : i+int(size)])
		i += int(size)
		if i+3 >= len(index) {
			break
		}
		block = binary.LittleEndian.Uint32(index[i : i+4])
		i += 4
		if !foundStartBlock && startKey <= indexKey {
			startBlock = block
			foundStartBlock = true
		}
		if endKey <= indexKey {
			return startBlock, block
		}
	}
	return startBlock, block
}

func findKeysInBlocks(startKey, endKey string, data []byte) (entries []*LSMDataEntry) {
	for i := 0; i < len(data); i += BlockSize {
		block := data[i : i+BlockSize]
		j := 0
		for j+8 < len(block) {
			seqID := binary.LittleEndian.Uint64(block[j : j+8])
			j += 8
			if j >= len(block) {
				break
			}
			keySize := uint8(block[j])
			j++
			if j+int(keySize)-1 >= len(block) {
				break
			}
			key := string(block[j : j+int(keySize)])
			j += int(keySize)
			if j >= len(block) {
				break
			}
			valueType := uint8(block[j])
			j++
			if j+1 >= len(block) {
				break
			}
			valueSize := binary.LittleEndian.Uint16(block[j : j+2])
			j += 2
			if j+int(valueSize)-1 >= len(block) {
				break
			}
			value := block[j : j+int(valueSize)]
			j += int(valueSize)

			if startKey <= key && key <= endKey {
				entries = append(entries, &LSMDataEntry{
					seqID:     seqID,
					keySize:   keySize,
					key:       key,
					valueType: valueType,
					valueSize: valueSize,
					value:     value,
				})
			} else if key > endKey {
				return entries
			}
		}
	}
	return entries
}
