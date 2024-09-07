package gobitcask

import (
	"io"
	"path/filepath"
	"sync"
)

type GoBitcask struct {
	indexes map[string]int64 // 内存中的索引信息 key -> latest offset
	dbFile  *DBFile          // 数据文件
	dirPath string           // file path for DB
	mu      sync.RWMutex
}

func Open(dirPath string) (*GoBitcask, error) {
	err := createFolder(dirPath)
	if err != nil {
		return nil, err
	}

	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	dbFile, err := NewDBFile(absDirPath)
	if err != nil {
		return nil, err
	}
	db := &GoBitcask{
		indexes: make(map[string]int64),
		dbFile:  dbFile,
		dirPath: absDirPath,
	}

	db.readIndexesFromDBFile()
	return db, nil
}

func (db *GoBitcask) Put(key []byte, value []byte) (err error) {
	if len(key) == 0 {
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()

	offset := db.dbFile.Offset
	entry := NewEntry(key, value, PUT)
	err = db.dbFile.Write(entry)

	db.indexes[string(key)] = offset
	return
}

func (db *GoBitcask) exist(key []byte) (int64, error) {
	offset, ok := db.indexes[string(key)]
	if !ok {
		return 0, ErrKeyNotFound
	}
	return offset, nil
}

func (db *GoBitcask) Get(key []byte) (value []byte, err error) {
	if len(key) == 0 {
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()

	// 1. use hashmap to get offset for latest entry
	offset, err := db.exist(key)
	if err == ErrKeyNotFound {
		return
	}

	var e *Entry
	// 2. read dbFile with offset
	e, err = db.dbFile.Read(offset)
	if err != nil && err != io.EOF {
		return
	}
	if e != nil {
		value = e.Value
	}
	return
}

func (db *GoBitcask) readIndexesFromDBFile() (err error) {
	if db.dbFile == nil {
		return
	}
	var offset int64 = 0
	for {
		e, err := db.dbFile.Read(offset)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		db.indexes[string(e.Key)] = offset
		if e.Mark == DELETE {
			delete(db.indexes, string(e.Key))
		}
		offset += e.GetSize()
	}

	return nil
}

func (db *GoBitcask) Merge() error {
	return nil
}

func (db *GoBitcask) Close() error {
	if db.dbFile == nil {
		return ErrInvalidDBFile
	}
	return db.dbFile.File.Close()
}
