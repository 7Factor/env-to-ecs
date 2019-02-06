package args

import (
	"errors"
	"github.com/docopt/docopt-go"
	"os"
)

// Default return should at a minimum have the
// InFile value populated with something.
type Config struct {
	InFile  string
	OutFile string
}

// docopt expects this to be in a very specify format, edit with caution
const docString = `
Usage: env_to_ecs [INFILE] [-oh] [OUTFILE]

Process INFILE and convert it to a new file type.

Arguments:
  INFILE            Required input file.
  OUTFILE           Optional output file.

Options:
  -o --output       Specify output file.
  -h --help         Display cli info.
`

func GetArguments() (Config, error) {
	// setup docString
	args, err := docopt.Parse(docString, os.Args[1:], true, "", false)
	if err != nil {
		return Config{}, errors.New("error parsing args")
	}

	inFile, err := parseArgOrError(args["INFILE"])
	if err != nil {
		return Config{}, errors.New("INFILE cannot be empty")
	}

	var outFile string
	if hasOutputFlag(args) {
		outFile, err = parseArgOrError(args["OUTFILE"])
		if err != nil {
			return Config{}, err
		}
	} else {
		outFile = "stdout"
	}

	return Config{InFile: inFile, OutFile: outFile}, nil
}

func parseArgOrError(arg interface{}) (string, error) {
	if arg == nil {
		return "", errors.New("arg cannot be nil")
	} else {
		return arg.(string), nil
	}
}

func hasOutputFlag(args map[string]interface{}) bool {
	return args["-o"] != nil || args["--output"] != false
}
