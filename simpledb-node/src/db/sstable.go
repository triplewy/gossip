package db

import (
	"encoding/binary"
	"errors"
	"os"
)

type SSTable struct {
	level0 *Level
	level1 *Level
}

type FindReply struct {
	offset uint64
	size   uint32
	err    error
}

func NewSSTable() (*SSTable, error) {
	level0, err := NewLevel(0)
	if err != nil {
		return nil, err
	}

	level1, err := NewLevel(1)
	if err != nil {
		return nil, err
	}

	return &SSTable{
		level0: level0,
		level1: level1,
	}, nil
}

func (table *SSTable) Append(blocks, index []byte, startKey, endKey string) error {
	var f *os.File
	var err error

	filename := table.level0.getUniqueId()

	f, err = os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()

	if err != nil {
		return err
	}

	filler := make([]byte, indexBlockSize-len(index))
	index = append(index, filler...)

	if len(index) != indexBlockSize {
		return errors.New("LSM index block does not match 16 KB")
	}

	_, err = f.Write(append(blocks, index...))
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	err = table.level0.NewSSTFile(filename, startKey, endKey)
	if err != nil {
		return err
	}

	return nil
}

func (table *SSTable) Find(key string) ([]*FindReply, error) {
	filenames := table.level0.FindSSTFile(key)
	if len(filenames) == 0 {
		return nil, errors.New("No SSTables in Level 0 contain key")
	}

	replyChan := make(chan *FindReply, len(filenames))

	for _, filename := range filenames {
		go func(filename, key string) {
			table.find(filename, key, replyChan)
		}(filename, key)
	}

	replies := []*FindReply{}

	for reply := range replyChan {
		if reply.err != nil {
			return nil, reply.err
		}
		replies = append(replies, reply)
	}

	return replies, nil
}

func (table *SSTable) find(filename, key string, replyChan chan *FindReply) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer f.Close()

	if err != nil {
		replyChan <- &FindReply{
			offset: 0,
			size:   0,
			err:    err,
		}
		return
	}

	index := make([]byte, indexBlockSize)
	numBytes, err := f.ReadAt(index, int64(128*blockSize))

	if err != nil {
		replyChan <- &FindReply{
			offset: 0,
			size:   0,
			err:    err,
		}
		return
	}

	if numBytes != indexBlockSize {
		replyChan <- &FindReply{
			offset: 0,
			size:   0,
			err:    errors.New("Did not read correct amount of bytes for index"),
		}
		return
	}

	blockIndex := findDataBlock(key, index)

	block := make([]byte, blockSize)
	numBytes, err = f.ReadAt(block, int64(blockSize*blockIndex))
	if err != nil {
		replyChan <- &FindReply{
			offset: 0,
			size:   0,
			err:    err,
		}
		return
	}
	if numBytes != blockSize {
		replyChan <- &FindReply{
			offset: 0,
			size:   0,
			err:    errors.New("Did not read correct amount of bytes for data block"),
		}
		return
	}

	offset, size, err := findKeyInBlock(key, block)
	replyChan <- &FindReply{
		offset: offset,
		size:   size,
		err:    err,
	}
}

func findDataBlock(key string, index []byte) uint16 {
	i := 0
	for i < len(index) {
		size := uint8(index[i])
		i++
		indexKey := string(index[i : i+int(size)])

		i += int(size)
		if key < indexKey {
			return binary.LittleEndian.Uint16(index[i:i+2]) - 1
		}
		i += 2
	}

	return uint16(l0Size - 1)
}

func findKeyInBlock(key string, block []byte) (offset uint64, size uint32, err error) {
	i := 0
	for i < blockSize {
		size := uint8(block[i])
		i++
		foundKey := string(block[i : i+int(size)])
		i += int(size)
		if key == foundKey {
			valueOffset := binary.LittleEndian.Uint64(block[i : i+8])
			i += 8
			valueSize := binary.LittleEndian.Uint32(block[i : i+4])
			i += 4
			return valueOffset, valueSize, nil
		}
		i += 12
	}

	return 0, 0, errors.New("Key not found")
}