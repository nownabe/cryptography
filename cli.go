package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
)

const (
	Name = "cryptography"

	ExitCodeOK = iota
	ExitCodeNG
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		in      string
		out     string
		key     string
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprintln(cli.errStream, usage)
	}
	flags.StringVar(&in, "in", "", "Input file path")
	flags.StringVar(&out, "out", "", "Output file path")
	flags.StringVar(&key, "key", "", "Encryption key string")
	flags.BoolVar(&version, "version", false, "Print version")
	flags.BoolVar(&version, "v", false, "Print version")

	if err := flags.Parse(args[2:]); err != nil {
		return ExitCodeNG
	}

	if version {
		fmt.Fprintln(cli.outStream, Version)
		return ExitCodeOK
	}

	var cryptographer Cryptographer
	var err error

	switch args[1] {
	case "enc":
		cryptographer, err = NewEncryptor(key)
	case "dec":
		cryptographer, err = NewDecryptor(key)
	default:
		fmt.Fprintln(cli.errStream, "Invalid command "+args[1])
		return ExitCodeNG
	}

	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeNG
	}

	input, err := ioutil.ReadFile(in)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeNG
	}

	result, err := cryptographer.Exec(input)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeNG
	}

	err = ioutil.WriteFile(out, result, 0644)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeNG
	}

	return ExitCodeOK
}

var usage = `Usage: cryptography (enc|dec) -in input_path -out output_path -key encryption_key`
