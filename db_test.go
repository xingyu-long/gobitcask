package gobitcask

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const dirPath = "/tmp/gobitcask"

func TestOpen(t *testing.T) {
	db, err := Open(dirPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(db)
}

func TestDB_Put(t *testing.T) {
	db, err := Open(dirPath)
	if err != nil {
		t.Error(err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	keyPrefix := "test_key_"
	valPrefix := "test_val_"

	for i := 0; i < 50; i++ {
		key := []byte(keyPrefix + strconv.Itoa(i%5))
		val := []byte(valPrefix + strconv.FormatInt(r.Int63(), 10))
		err = db.Put(key, val)
	}

	if err != nil {
		t.Log(err)
	}
}

func TestDB_Get(t *testing.T) {
	db, err := Open(dirPath)
	if err != nil {
		t.Error(err)

	}

	getVal := func(key []byte) (err error) {
		val, err := db.Get(key)
		if err != nil {
			t.Error("error val: ", err)
		} else {
			t.Logf("key = %v, val=%v", string(key), string(val))
		}
		return
	}

	getVal([]byte("test_key_0"))
	getVal([]byte("test_key_1"))
	getVal([]byte("test_key_2"))
	getVal([]byte("test_key_3"))
	getVal([]byte("test_key_4"))

	_, err = db.Get([]byte("test_key_5"))
	t.Log(err)
	if err == ErrKeyNotFound {
		t.Log("test_key_5 does not exist as expected")
	}
}

func TestDB_Delete(t *testing.T) {

}
