package logic

import (
	"fblviewer/binlog"
	"fblviewer/logdata"
)

type logic struct {
	handle *binlog.Handle
}

func NewLogic(filePath string) (*logic, error) {
	logic := &logic{}
	var err error
	logic.handle, err = binlog.NewHandle(filePath)
	if err != nil {
		panic(err)
	}
	return logic, nil
}

func (l *logic)List() []*logdata.Log {

	res := make([]*logdata.Log, 0)
	for  {
		log, err := l.handle.Next()
		if err != nil {
			break
		}
		res = append(res, log)
	}
	return res
}