package parse

import (
	"encoding/binary"
	"math"
	"reflect"
)

func (Bool)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	if valueByte[0] == 1 {
		return true, typeByte[1:], nameByte, valueByte[1:]
	}
	return false, typeByte[1:], nameByte, valueByte[1:]
}

func (Int)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int(binary.BigEndian.Uint32(valueByte[0:4])), typeByte[1:], nameByte, valueByte[4:]
}

func (Int8)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int8(valueByte[0]), typeByte[1:], nameByte,valueByte[1:]
}

func (Int16)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int16(binary.BigEndian.Uint16(valueByte[0:2])), typeByte[1:], nameByte,valueByte[2:]
}

func (Int32)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int32(binary.BigEndian.Uint32(valueByte[0:4])), typeByte[1:], nameByte, valueByte[4:]
}

func (Int64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int64(binary.BigEndian.Uint64(valueByte[0:8])),typeByte[1:], nameByte, valueByte[8:]
}

func (Uint)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return uint(binary.BigEndian.Uint32(valueByte[0:4])), typeByte[1:], nameByte, valueByte[4:]
}


func (Uint8)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return uint8(valueByte[0]), typeByte[1:], nameByte,valueByte[1:]
}

func (Uint16)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return binary.BigEndian.Uint16(valueByte[0:2]), typeByte[1:], nameByte,valueByte[2:]
}


func (Uint32)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return binary.BigEndian.Uint32(valueByte[0:4]), typeByte[1:], nameByte, valueByte[4:]
}

func (Uint64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return uint64(binary.BigEndian.Uint64(valueByte[0:8])), typeByte[1:], nameByte, valueByte[8:]
}

func (Float32)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return math.Float32frombits(binary.BigEndian.Uint32(valueByte[0:4])),
		typeByte[1:],
		nameByte,
		valueByte[4:]
}

func (Float64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return math.Float64frombits(binary.BigEndian.Uint64(valueByte[0:8])),
		typeByte[1:],
		nameByte,
		valueByte[8:]
}

func (Complex64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	realVar := math.Float32frombits(binary.BigEndian.Uint32(valueByte[0:4]))
	imagVar := math.Float32frombits(binary.BigEndian.Uint32(valueByte[4:8]))
	return complex(realVar,imagVar), typeByte[1:],nameByte, valueByte[8:]
}

func (Complex128)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	realVar := math.Float64frombits(binary.BigEndian.Uint64(valueByte[0:8]))
	imagVar := math.Float64frombits(binary.BigEndian.Uint64(valueByte[8:16]))
	return complex(realVar,imagVar), typeByte[1:],nameByte, valueByte[16:]
}

func (String) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	var (
		i int
		v byte
	)
	for i, v = range valueByte {
		if v == us {
			break
		}
	}
	return string(valueByte[:i]), typeByte[1:], nameByte, valueByte[i+1:]
}

func (Slice)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte)  {
	length := binary.BigEndian.Uint32(valueByte[:4])
	valueByte = valueByte[4:]
	resValue := make([]interface{}, length)

	itemType := reflect.Kind(typeByte[1])
	for i := uint32(0); i<length; i++{
		resValue[i], _, nameByte,valueByte = MapParse[itemType].Decode(typeByte[1:],nameByte, valueByte)
	}

	return resValue, typeByte[2:], nameByte, valueByte
}

func (Uintptr)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}

func (UnsafePointer)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}

func (Ptr)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}

func (Struct) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte){
	numField := binary.BigEndian.Uint32(typeByte[1:5])
	var resValue = make(map[string]interface{})
	typeByte = typeByte[5:]
	var (
		k int
		v byte
		value interface{}
	)
	for i:=uint32(0);i<numField;i++ {
		for k,v = range nameByte {
			if v == us {
				break
			}
		}
		name := string(nameByte[:k])
		value, typeByte, nameByte, valueByte = MapParse[reflect.Kind(typeByte[0])].Decode(typeByte, nameByte[k+1:], valueByte)
		resValue[name] = value
	}
	return resValue, typeByte, nameByte, valueByte
}

func (Map) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	typeByte = typeByte[1:]
	// map长度
	length := binary.BigEndian.Uint32(typeByte[:4])
	typeByte = typeByte[4:]
	var (
		k int
		v byte
		i uint32
		value interface{}
		resValue map[string]interface{} = make(map[string]interface{})
	)
	for i = 0;i<length;i++ {
		for k,v = range nameByte {
			if v == us {
				break
			}
		}
		name := string(nameByte[:k])
		value, typeByte, nameByte, valueByte = MapParse[reflect.Kind(typeByte[0])].Decode(typeByte, nameByte[k+1:], valueByte)
		resValue[name] = value
	}
	return resValue, typeByte, nameByte, valueByte
}

func (Empty) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}