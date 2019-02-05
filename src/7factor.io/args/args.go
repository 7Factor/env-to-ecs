package args

import (
	"errors"
	"fmt"
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
Usage: env_to_ecs  [INFILE] [-o]

Process INFILE and converts it to a new file type.

Arguments:
  INFILE        Required input file.
  OUTFILE       Optional output file.

Options:
  -o --output       verbose mode
`

func GetArguments() (Config, error) {
	args, err := docopt.Parse(docString, os.Args, true, "", false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return Config{}, errors.New("error parsing args")
	}

	if args["INFILE"] == nil {
		return Config{}, errors.New("did not find file to parse")
	}

	if hasOutputFlag(args) {
		return Config{}, errors.New("must specify outfile")
	}

	envFile := args["INFILE"].(string)

	return Config{InFile: envFile}, nil
}

func hasOutputFlag(args map[string]interface{}) bool {
	return args["-o"] != nil || args["--output"] != false
}
