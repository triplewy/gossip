package db

import (
	"fmt"
	"path/filepath"
	"sync"
)

// LSM is struct for all levels in an LSM
type LSM struct {
	levels []*Level
	fm     *FileManager
}

// NewLSM instatiates all levels for a new LSM tree
func NewLSM(directory string) (*LSM, error) {
	fm := NewFileManager()
	levels := []*Level{}
	for i := 0; i < 7; i++ {
		level, err := NewLevel(i, directory, fm)
		if err != nil {
			return nil, err
		}
		if i > 0 {
			above := levels[i-1]
			above.below = level
			level.above = above
		}
		levels = append(levels, level)
	}

	return &LSM{
		levels: levels,
		fm:     fm,
	}, nil
}

// Append takes data blocks, an index block, and a key range as input and writes an SST File to level 0.
// It then adds this new file to level 0's manifest
func (lsm *LSM) Append(blocks, index []byte, bloom *Bloom, keyRange *KeyRange) error {
	level := lsm.levels[0]
	fileID := level.getUniqueID()
	filename := filepath.Join(level.directory, fileID+".sst")

	keyRangeEntry := createKeyRangeEntry(keyRange)
	header := createHeader(len(blocks), len(index), len(bloom.bits), len(keyRangeEntry))
	data := append(header, append(append(append(blocks, index...), bloom.bits...), keyRangeEntry...)...)

	err := lsm.fm.Write(filename, data)
	if err != nil {
		return err
	}

	level.NewSSTFile(fileID, keyRange, bloom)

	return nil
}

// Find goes through each level of the LSM tree and returns if a result is found for the given key.
// If no result is found, Find throws a KeyNotFound error
func (lsm *LSM) Find(key string, ts uint64) (*LSMDataEntry, error) {
	for _, level := range lsm.levels {
		entry, err := level.find(key, ts)
		if err != nil {
			switch err.(type) {
			case *ErrKeyNotFound:
				continue
			default:
				return nil, err
			}
		} else {
			return entry, nil
		}
	}
	return nil, NewErrKeyNotFound()
}

// Range concurrently finds all keys in the LSM tree that fall within the range query.
// Concurrency is achieved by going through each level on its own goroutine
func (lsm *LSM) Range(keyRange *KeyRange, ts uint64) ([]*LSMDataEntry, error) {
	replyChan := make(chan []*LSMDataEntry)
	errChan := make(chan error)
	result := []*LSMDataEntry{}
	errs := make(map[string]int)
	var wg sync.WaitGroup

	wg.Add(7)
	for _, level := range lsm.levels {
		go func(level *Level) {
			entries, err := level.Range(keyRange, ts)
			if err != nil {
				errChan <- err
			} else {
				replyChan <- entries
			}
		}(level)
	}

	go func() {
		for {
			select {
			case reply := <-replyChan:
				result = append(result, reply...)
				wg.Done()
			case err := <-errChan:
				if _, ok := errs[err.Error()]; !ok {
					errs[err.Error()] = 1
				} else {
					errs[err.Error()]++
				}
				wg.Done()
			}
		}
	}()
	wg.Wait()

	if len(errs) > 0 {
		return result, fmt.Errorf("Errors during LSM range query: %v", errs)
	}
	return result, nil
}

// CheckPrimaryKey traverses through LSM and checks if key exists
// func (lsm *LSM) CheckPrimaryKey(key string, levelNum int) (bool, error) {
// 	if levelNum > 6 {
// 		return false, nil
// 	}

// 	level := lsm.levels[levelNum]

// 	filenames := level.FindSSTFile(key)
// 	if len(filenames) == 0 {
// 		return lsm.CheckPrimaryKey(key, levelNum+1)
// 	}

// 	replyChan := make(chan *LSMDataEntry)
// 	errChan := make(chan error)

// 	replies := []*LSMDataEntry{}

// 	var wg sync.WaitGroup
// 	var errs []error

// 	wg.Add(len(filenames))
// 	for _, filename := range filenames {
// 		go func(filename string) {
// 			entry, err := lsm.fm.Find(filename, key,)
// 			if err != nil {
// 				errChan <- err
// 			} else {
// 				replyChan <- entry
// 			}
// 		}(filename)
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case reply := <-replyChan:
// 				replies = append(replies, reply)
// 				wg.Done()
// 			case err := <-errChan:
// 				errs = append(errs, err)
// 				wg.Done()
// 			}
// 		}
// 	}()

// 	wg.Wait()

// 	if len(replies) > 0 {
// 		return true, nil
// 	}

// 	for _, err := range errs {
// 		if _, ok := err.(*ErrKeyNotFound); ok {
// 			continue
// 		} else if strings.HasSuffix(err.Error(), "no such file or directory") {
// 			fmt.Println(key, err)
// 		} else {
// 			return true, err
// 		}
// 	}

// 	return lsm.CheckPrimaryKey(key, levelNum+1)
// }

// Close closes all levels in the LSM
func (lsm *LSM) Close() {
	for _, level := range lsm.levels {
		level.Close()
	}
}
