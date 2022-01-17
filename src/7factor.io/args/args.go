package args

import (
	"github.com/jessevdk/go-flags"
)

type ArgConfig struct {
	InFile     string
	OutFile    string
	JsonParent *string
	Variables  []string
}

func GetArguments() (ArgConfig, error) {
	_, err := flags.Parse(&opt)

	return ArgConfig{opt.Infile, opt.Outfile, opt.JsonParent, opt.Variables}, err
}
