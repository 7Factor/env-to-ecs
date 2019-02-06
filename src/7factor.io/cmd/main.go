package main

import (
	"7factor.io/args"
	"7factor.io/converter"
	"fmt"
	"os"
)

func main() {
	config, err := args.GetArguments()

	if err != nil {
		fmt.Fprintf(os.Stderr, "received error while getting args: %v\n", err)
		os.Exit(1)
	}

	_, err = converter.ReadAndConvert(config.InFile, config.OutFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "recieved error while reading and converting args: %v\n", err)
		os.Exit(1)
	}
}
