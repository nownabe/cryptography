package main

type Cryptographer interface {
	Exec([]byte) ([]byte, error)
}
