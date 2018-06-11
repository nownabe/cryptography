package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	ExitCodeOK = iota
	ExitCodeNG

	Name = "cryptography"
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		enc     bool
		dec     bool
		in      string
		out     string
		key     string
		version bool
	)

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = cli.printHelp
	flags.BoolVar(&enc, "enc", false, "Encrypt")
	flags.BoolVar(&dec, "dec", false, "Decrypt")
	flags.StringVar(&in, "in", "", "Input file path")
	flags.StringVar(&out, "out", "", "Output file path")
	flags.StringVar(&key, "key", "", "Encryption key string")
	flags.BoolVar(&version, "version", false, "Print version")
	flags.BoolVar(&version, "v", false, "Print version")

	if len(args) < 2 {
		cli.printHelp()
		return ExitCodeNG
	}

	mode := args[1]
	err := flags.Parse(args[2:])
	cli.chkErr(err)

	if version {
		fmt.Fprintln(cli.outStream, Version)
		return ExitCodeOK
	}

	var cryptographer Cryptographer

	switch mode {
	case "enc":
		cryptographer, err = NewEncryptor(key)
	case "dec":
		cryptographer, err = NewDecryptor(key)
	default:
		fmt.Fprintln(cli.errStream, "Invalid command")
		return ExitCodeNG
	}
	cli.chkErr(err)

	input, err := ioutil.ReadFile(in)
	cli.chkErr(err)

	result, err := cryptographer.Exec(input)
	cli.chkErr(err)

	err = ioutil.WriteFile(out, result, 0644)
	cli.chkErr(err)

	return ExitCodeOK
}

func (cli *CLI) chkErr(err error) {
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		os.Exit(ExitCodeNG)
	}
}

func (cli *CLI) printHelp() {
	fmt.Fprintln(cli.errStream, usage)
}

var usage = `Usage: cryptography (enc|dec) -in input_path -out output_path -key encryption_key`
