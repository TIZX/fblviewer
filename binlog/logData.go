package binlog

import "encoding/binary"

// log data
type logData struct {
	length uint32
	generalID uint32
	time uint32
	dataByte []byte
}


func (d *logData)Decode(dataByte []byte) {
	d.length = binary.BigEndian.Uint32(dataByte[0:4])
	d.generalID = binary.BigEndian.Uint32(dataByte[4:8])
	d.time = binary.BigEndian.Uint32(dataByte[8:12])
	d.dataByte = dataByte[12:]
}
