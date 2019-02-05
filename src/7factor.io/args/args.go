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
	EnvironmentFile string
	OutputFile string
}

const docString =
`
Usage: env_to_ecs [FILE] [--flags]

Process FILE and convert to a new file type.

Arguments:
  FILE        optional input file
`

func GetArguments() (Config, error) {
	args, err := docopt.Parse(docString, os.Args, true, "", false)
	if err != nil {
		fmt.Fprintln(os.Stderr)
		return Config{}, errors.New("error parsing args")
	}

	if args["FILE"] == nil {
		return Config{}, errors.New("did not find file to parse")
	}

	envFile := args["FILE"].(string)

	return Config{EnvironmentFile:envFile}, nil
}
