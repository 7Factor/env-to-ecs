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
		fmt.Fprint(os.Stderr, err)
	}

	converter.ReadAndConvert(config.InFile, config.OutFile)
}
