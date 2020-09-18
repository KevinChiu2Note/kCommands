package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(msg string) ([]byte, error) {
	length := int32(len(msg))
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lenBytes, err := reader.Peek(4)
	if err != nil {
		return "", err
	}
	lenBuf := bytes.NewBuffer(lenBytes)
	var length int32
	err = binary.Read(lenBuf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
