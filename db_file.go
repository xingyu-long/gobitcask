package gobitcask

import (
	"os"
	"path/filepath"
	"sync"
)

type DBFile struct {
	File          *os.File
	Offset        int64
	HeaderBufPool *sync.Pool
}

const FileName = "gobitcask.data"
const MergeFileName = "gobitcask.data.merge"

// create new DB file and return it
func newInternal(filePath string) (*DBFile, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	stat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	// what's this?
	pool := &sync.Pool{New: func() interface{} {
		return make([]byte, entryHeaderSize)
	}}

	return &DBFile{Offset: stat.Size(), File: file, HeaderBufPool: pool}, nil
}

func NewDBFile(path string) (*DBFile, error) {
	filePath := filepath.Join(path, FileName)
	return newInternal(filePath)
}

func NewDBMergeFile(path string) (*DBFile, error) {
	filePath := filepath.Join(path, MergeFileName)
	return newInternal(filePath)
}

// write Entry to db file
func (df *DBFile) Write(e *Entry) (err error) {
	enc, err := e.Encode()
	if err != nil {
		return nil
	}

	_, err = df.File.WriteAt(enc, df.Offset)
	df.Offset += e.GetSize()
	return
}

// read Entry from db file with offset
func (df *DBFile) Read(offset int64) (e *Entry, err error) {
	// https://hackernoon.com/go-and-syncpool
	buffer := df.HeaderBufPool.Get().([]byte)
	defer df.HeaderBufPool.Put(buffer)

	// check if we can read buffer from offset position
	if _, err = df.File.ReadAt(buffer, offset); err != nil {
		return
	}

	if e, err = Decode(buffer); err != nil {
		return
	}

	offset += entryHeaderSize
	if e.KeySize > 0 {
		key := make([]byte, e.KeySize)
		if _, err = df.File.ReadAt(key, offset); err != nil {
			return
		}
		e.Key = key
	}

	offset += int64(e.KeySize)
	if e.ValueSize > 0 {
		value := make([]byte, e.ValueSize)
		if _, err = df.File.ReadAt(value, offset); err != nil {
			return
		}
		e.Value = value
	}
	return
}
