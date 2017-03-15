package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Encryptor struct {
	Block cipher.Block
}

func NewEncryptor(key string) (*Encryptor, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &Encryptor{Block: block}, nil
}

func (e *Encryptor) Exec(input []byte) (output []byte, err error) {
	output = make([]byte, aes.BlockSize+len(input))
	iv := output[:aes.BlockSize]
	res := output[aes.BlockSize:]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return
	}
	stream := cipher.NewCTR(e.Block, iv)
	stream.XORKeyStream(res, input)
	return
}
