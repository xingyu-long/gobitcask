# gobitcask

This is an implementation of [Bitcask](https://riak.com/assets/bitcask-intro.pdf) (key/value database) in Golang based on [mini-bitcask](https://github.com/rosedblabs/mini-bitcask/tree/main)

## Usage
```go
package main

import (
	"bytes"
	"fmt"

	"github.com/xingyu-long/gobitcask"
)

func main() {
	db, err := gobitcask.Open("/tmp/gobitcask_main")
	if err != nil {
		panic(err)
	}

	key := []byte("test_key")
	value := []byte("test_value")

	err = db.Put(key, value)
	if err != nil {
		panic(err)
	}
	fmt.Printf("1. Run %s: %s\n", "Put", fmt.Sprintf("put key = %s, value = %s into kv", string(key), string(value)))

	ret_val, err := db.Get(key)
	if err != nil {
		panic(err)
	}
	if bytes.Equal(ret_val, value) {
		fmt.Printf("2. Run %s: %s\n", "Get", fmt.Sprintf("value = %s", string(ret_val)))
	} else {
		panic(fmt.Sprintf("Expected value = %s, Actual value = %s\n", string(value), string(ret_val)))
	}

	err = db.Delete(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("3. Run %s: %s\n", "Delete", fmt.Sprintf("key = %s", string(key)))

	db.Merge()
	fmt.Printf("4. Run %s: %s\n", "Merge", "merge data and create new DBFile")

	db.Close()
	fmt.Printf("5. Run %s: %s\n", "Close", "close DB")
}
```