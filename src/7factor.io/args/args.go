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

	return ArgConfig{}, nil
}

func init() {
	setInfileFlag()
}

func setInfileFlag() {
	const (
		defaultInfile = ""
		usage         = "The infile to parse."
	)
	flag.StringVar(&argConfig.InFile, "i", defaultInfile, usage)
	flag.StringVar(&argConfig.InFile, "infile", defaultInfile, usage)
}
