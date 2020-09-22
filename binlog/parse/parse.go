package parse

import (
	"encoding/binary"
	"reflect"
)

// 解码
func Decode(typeName, valueByte []byte) map[string]interface{} {
	typeLen := binary.BigEndian.Uint32(typeName[0:4])
	typeByte := typeName[4:typeLen+4]

	nameByte := typeName[typeLen+4:]

	value, _,_,_ := MapParse[reflect.Map].Decode(typeByte, nameByte, valueByte)
	return value.(map[string]interface{})
}