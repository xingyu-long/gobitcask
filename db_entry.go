package gobitcask

import "encoding/binary"

const entryHeaderSize = 10
const (
	PUT uint16 = iota
	DELETE
)

type Entry struct {
	Key       []byte
	Value     []byte
	KeySize   uint32
	ValueSize uint32
	Mark      uint16 // operation: PUT or DELETE
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

func (e *Entry) GetSize() int64 {
	return int64(entryHeaderSize + e.KeySize + e.ValueSize)
}

/*
Little Endian vs Big Endian (https://www.youtube.com/watch?v=T1C9Kj_78ek)

MSB(most significant bytes)
LSB(least significant bytes)

    MSB                       LSB
| byte 3 | byte 2 | byte 1 | byte 0 |

Little Endian: Byte0, Byte1, Byte2, Byte3
| Byte 3 |  <--- high address
| Byte 2 |
| Byte 1 |
| Byte 0 |  <--- low address

Big Endian: Byte3, Byte2, Byte1, Byte0
| Byte 0 |  <--- high address
| Byte 1 |
| Byte 2 |
| Byte 3 |  <--- low address
*/
// convert Entry object to bytes
func (e *Entry) Encode() ([]byte, error) {
	buffer := make([]byte, e.GetSize())
	// why 0:4; 4:8?
	// byte -> uint8
	// uint32 = 4 * byte
	binary.BigEndian.PutUint32(buffer[0:4], e.KeySize)
	binary.BigEndian.PutUint32(buffer[4:8], e.ValueSize)
	binary.BigEndian.PutUint16(buffer[8:10], e.Mark)

	copy(buffer[entryHeaderSize:entryHeaderSize+e.KeySize], e.Key)
	copy(buffer[entryHeaderSize+e.KeySize:], e.Value)

	return buffer, nil
}

// convert bytes to Entry object
func Decode(buf []byte) (*Entry, error) {
	keySize := binary.BigEndian.Uint32(buf[0:4])
	valueSize := binary.BigEndian.Uint32(buf[4:8])
	mark := binary.BigEndian.Uint16(buf[8:10])
	// why don't we read the content here?
	return &Entry{KeySize: keySize, ValueSize: valueSize, Mark: mark}, nil
}
