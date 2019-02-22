package args

import (
	"errors"
	"flag"
)

type ArgConfig struct {
	InFile  string
	OutFile string
}

var argConfig ArgConfig

func GetArguments() (ArgConfig, error) {
	flag.Parse()

	if argConfig.InFile == "" {
		return ArgConfig{}, errors.New("infile cannot be empty")
	}

	return ArgConfig{argConfig.InFile, argConfig.OutFile}, nil
}

func init() {
	setFlag(&argConfig.InFile, "i", "infile", "", "The infile to parse")
	setFlag(&argConfig.OutFile, "o", "outfile", "stdout", "The outfile to write to.")
}

func setFlag(flagVar *string, shortFlag string, longFlag string, defaultValue string, usage string) {
	flag.StringVar(flagVar, shortFlag, defaultValue, usage)
	flag.StringVar(flagVar, longFlag, defaultValue, usage)
}
