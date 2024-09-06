package gobitcask

import "os"

type DBFile struct {
	File   *os.File
	Offset int64
}

// create new DB file and return it
func newDBFile(fileName string) (*DBFile, error) {
	return &DBFile{}, nil
}

// write Entry to db file
func (df *DBFile) Write(e *Entry) (err error) {
	return nil
}

// read Entry from db file with offset
func (df *DBFile) Read(offset int64) (e *Entry, err error) {
	return &Entry{}, nil
}
