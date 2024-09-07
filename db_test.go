package gobitcask

import (
	"bytes"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

const dirPath = "/tmp/gobitcask"

func destroyDB(db *GoBitcask) {
	_ = db.Close()
	_ = os.RemoveAll(db.dirPath)
}

func TestOpen(t *testing.T) {
	db, err := Open(dirPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(db)
}

func TestDB_PutGet(t *testing.T) {
	db, err := Open(dirPath)
	if err != nil {
		t.Error(err)
	}
	defer destroyDB(db)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	keyPrefix := "test_key_"
	valPrefix := "test_val_"

	for i := 0; i < 10000; i++ {
		key := []byte(keyPrefix + strconv.Itoa(i%5))
		val := []byte(valPrefix + strconv.FormatInt(r.Int63(), 10))
		err = db.Put(key, val)
		if err != nil {
			t.Error(err)
		}
		ret_val, err := db.Get(key)
		if !bytes.Equal(val, ret_val) {
			t.Error(err)
		}
	}

	_, err = db.Get([]byte("test_key_5"))
	t.Log(err)
	if err == ErrKeyNotFound {
		t.Log("test_key_5 does not exist as expected")
	}

	if err != nil {
		t.Log(err)
	}
}
func TestDB_Delete(t *testing.T) {
	db, err := Open(dirPath)
	if err != nil {
		t.Error(err)
	}
	defer destroyDB(db)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	keyPrefix := "test_key_"
	valPrefix := "test_val_"

	for i := 0; i < 5; i++ {
		key := []byte(keyPrefix + strconv.Itoa(i%5))
		val := []byte(valPrefix + strconv.FormatInt(r.Int63(), 10))
		err = db.Put(key, val)
		if err != nil {
			t.Fatal(err)
		}
		if i%2 == 0 {
			if err = db.Delete(key); err != nil {
				t.Fatal(err)
			}
			if _, err = db.Get(key); err != nil && err == ErrKeyNotFound {
				t.Logf("%s not found as expected", string(key))
			} else {
				t.Fatal(err)
			}
		}
	}
}

func TestDB_Merge(t *testing.T) {
	db, err := Open(dirPath)
	if err != nil {
		t.Error(err)
	}
	defer destroyDB(db)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	keyPrefix := "test_key_"
	valPrefix := "test_val_"

	for i := 0; i < 5; i++ {
		key := []byte(keyPrefix + strconv.Itoa(i%5))
		val := []byte(valPrefix + strconv.FormatInt(r.Int63(), 10))
		err = db.Put(key, val)
		if err != nil {
			t.Fatal(err)
		}
		if i%2 == 0 {
			if err = db.Delete(key); err != nil {
				t.Fatal(err)
			}
		}
	}

	db.Merge()

	for i := 0; i < 5; i++ {
		key := []byte(keyPrefix + strconv.Itoa(i%5))
		val, err := db.Get(key)
		if i%2 == 0 {
			if err != nil && err == ErrKeyNotFound {
				t.Logf("%s not found as expected", string(key))
			} else {
				t.Fatal(err)
			}
		} else {
			t.Logf("key = %s, val = %s", string(key), string(val))
		}
	}
}
