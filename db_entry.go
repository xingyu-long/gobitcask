package gobitcask

const entryHeaderSize = 10
const (
	PUT uint16 = iota
	DEL
)

type Entry struct {
	Key       []byte
	Value     []byte
	KeySize   uint32
	ValueSize uint32
	Mark      uint16 // what's this?
}

func NewEntry(key, value []byte, mark uint16) *Entry {
	return &Entry{
		Key:       key,
		Value:     value,
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		Mark:      mark,
	}
}

// convert Entry object to bytes
func (e *Entry) Encode() ([]byte, error) {
	return nil, nil
}

// convert bytes to Entry object
func (e *Entry) Decode(buf []byte) (*Entry, error) {
	return nil, nil
}
