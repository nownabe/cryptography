package main

import (
	"crypto/aes"
	"crypto/cipher"
)

type Decryptor struct {
	Block cipher.Block
}

func NewDecryptor(key string) (*Decryptor, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &Decryptor{Block: block}, nil
}

func (e *Decryptor) Exec(input []byte) (output []byte, err error) {
	output = make([]byte, len(input[aes.BlockSize:]))
	iv := input[:aes.BlockSize]
	stream := cipher.NewCTR(e.Block, iv)
	stream.XORKeyStream(output, input[aes.BlockSize:])
	err = nil
	return
}
