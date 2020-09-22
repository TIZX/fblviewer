package binlog

import (
	"time"
	"fblviewer/binlog/parse"
	"fblviewer/logdata"
)

type Index uint32

type Handle struct {
	file *logFile

}

func NewHandle(filePath string) (*Handle, error) {
	handle := &Handle{}
	var err error
	handle.file, err  = NewLogFile(filePath)
	if err != nil {
		return nil, err
	}
	err = handle.file.Init()
	if err != nil {
		return nil, err
	}
	return handle, err
}

// 解析日志
func (h *Handle) Next() (*logdata.Log, error)  {

	dataByte, err := h.file.Read()
	if err != nil {
		return nil, err
	}
	return h.parse(dataByte)
}

func (h *Handle)parse(dataByte []byte) (*logdata.Log, error) {
	data := &logData{}
	data.Decode(dataByte)
	general := h.file.generalMap[data.generalID]
	log := &logdata.Log{}
	log.File = general.File
	log.Level = general.Level
	log.Time = time.Unix(int64(data.time), 0)
	log.Line = general.Line
	log.Message = general.Message
	log.Fields = parse.Decode([]byte(general.TypeNameByte), data.dataByte)
	return log, nil
}
