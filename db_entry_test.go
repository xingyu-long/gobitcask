package gobitcask

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntryEncodeDecode(t *testing.T) {
	for i := 0; i < 10; i++ {
		entry := NewEntry([]byte(fmt.Sprint("key", i)), []byte(fmt.Sprint("value", i)), PUT)
		enc, err := entry.Encode()
		if err != nil {
			t.Fatal(err)
		}
		dec, err := Decode(enc)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(entry.KeySize, dec.KeySize)
		assert.Equal(entry.ValueSize, dec.ValueSize)
		assert.Equal(entry.Mark, dec.Mark)
	}
}
