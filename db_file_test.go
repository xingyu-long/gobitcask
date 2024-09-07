package gobitcask

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteReadEntry(t *testing.T) {
	dirPath := "/tmp/gobitcask_entry/"

	createFolder(dirPath)
	df, err := NewDBFile(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	assert := assert.New(t)
	var offset int64 = 0
	for i := 0; i < 5; i++ {
		entry := NewEntry([]byte(fmt.Sprint("key", i)), []byte(fmt.Sprint("value", i)), PUT)
		err = df.Write(entry)
		if err != nil {
			t.Fatal(err)
		}
		readEntry, err := df.Read(offset)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(entry.Key, readEntry.Key)
		assert.Equal(entry.Value, readEntry.Value)
		assert.Equal(entry.Mark, readEntry.Mark)
		offset += readEntry.GetSize()
	}

	deleteFolder(dirPath)
}
