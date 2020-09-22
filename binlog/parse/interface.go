package parse

type Decoder interface {
	Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte)
}

