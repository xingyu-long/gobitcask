package gobitcask

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

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
	keyPrefix := "test_key_"
	valPrefix := "test_val_"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10000; i++ {
		key := []byte(keyPrefix + strconv.Itoa(i%5))
		val := []byte(valPrefix + strconv.FormatInt(r.Int63(), 10))
		entry := NewEntry(key, val, PUT)
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
