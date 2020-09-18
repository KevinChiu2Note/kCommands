package protocol

import (
	"bufio"
	"bytes"
	"testing"
)

func TestCode(t *testing.T) {
	encode, err := Encode("hello")
	if err != nil {
		t.Fatal(err)
	}
	decode, err := Decode(bufio.NewReader(bytes.NewBuffer(encode)))
	if err != nil {
		t.Fatal(err)
	}
	if decode != "hello" {
		t.Fatal("encode or decode has problem")
	}
}
