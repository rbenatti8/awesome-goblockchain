package core

import (
	"encoding/binary"
	"io"
)

type Transaction struct {
	Data []byte
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, t.Data); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) DecodeBinary(r io.Reader) error {

	if err := binary.Read(r, binary.LittleEndian, &t.Data); err != nil {
		return err
	}
	return nil
}
