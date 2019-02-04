package args

import (
	"errors"
	"github.com/pborman/getopt/v2"
)

// Default return should at a minimum have the
// InFile value populated with something.
type Config struct {
	EnvironmentFile string
}

func GetArguments() (Config, error) {
	getopt.Parse()
	args := getopt.Args()

	if len(args) <= 0 {
		return Config{}, errors.New("did not find file to parse")
	}

	return Config{EnvironmentFile:args[0]}, nil
}
