package binlog

import (
	"encoding/binary"
)

type header struct {
	headerByte []byte
	dataOffset    uint64 //位置0
	generalOffset uint64 //位置8
}

func newHeader() *header {
	header := &header{
		headerByte: make([]byte, 16),
	}
	return header
}

func (f *header)ParseByte()  {
	f.dataOffset = binary.BigEndian.Uint64(f.headerByte[0:8])
	if f.dataOffset == 0 {
		f.dataOffset = 16
	}
	f.generalOffset = binary.BigEndian.Uint64(f.headerByte[8:16])
	if f.generalOffset == 0 {
		f.generalOffset = 16
	}
}

