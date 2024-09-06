package gobitcask

type GoBitcask struct {
	indexes map[string]int64 // 内存中的索引信息 key -> latest offset
	dbFile  *DBFile          // 数据文件
	dirPath string           // file path for DB
}

func Open(dirPath string) (*GoBitcask, error) {
	return nil, nil
}

func (db *GoBitcask) Put(key []byte, value []byte) error {
	return nil
}

func (db *GoBitcask) Get(key []byte) (value []byte, err error) {
	return nil, nil
}

func (db *GoBitcask) Merge() error {
	return nil
}
