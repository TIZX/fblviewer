package binlog

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

type logFile struct {
	header        *header
	file          *os.File
	size          int64 // 文件大小，所占字节数
	readSeek int64
	buf                *bufio.Reader
	generalMap map[uint32]*general
}

func NewLogFile(filePath string) (*logFile, error) {
	var err error
	file := &logFile{}
	file.file, err = os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	file.buf = bufio.NewReader(file.file)
	file.generalMap  = make(map[uint32]*general)
	return file, nil
}

func (l *logFile)Init() error {
	l.header = newHeader()
	_, err := l.buf.Read(l.header.headerByte)
	if err != nil {
		return err
	}

	l.header.ParseByte()
	_, _ = l.file.Seek(int64(l.header.generalOffset), io.SeekStart)
	for {
		var generalLen = make([]byte, 4)
		_, err := l.file.Read(generalLen)
		if err != nil && err == io.EOF {
			break
		}
		length := binary.BigEndian.Uint32(generalLen)
		var itemByte = make([]byte, length)
		_ ,err = l.file.Read(itemByte[4:])
		if err != nil && err == io.EOF {
			break
		}
		for k,v := range generalLen {
			itemByte[k] = v
		}
		item := &general{}
		item.Decode(itemByte)
		l.generalMap[item.ID] = item
	}

	_, _ = l.file.Seek(int64(l.header.dataOffset), io.SeekStart)
	l.readSeek = int64(l.header.dataOffset)
	return nil
}


// 获取单条log的byte切片
func (l *logFile)Read() ([]byte, error) {
	if uint64(l.readSeek) >= l.header.generalOffset {
		return nil, io.EOF
	}
	var lengthByte []byte = make([]byte, 4)
	n, err := l.buf.Read(lengthByte)
	if err != nil && n != 4 {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthByte)
	var data = make([]byte, length)

	for k,v := range lengthByte {
		data[k] = v
	}
	_, err = l.buf.Read(data[4:])
	if err != nil {
		return nil, err
	}
	l.readSeek = l.readSeek + int64(length)
	return data, nil
}

func (l *logFile) close() {
	l.file.Close()
}

