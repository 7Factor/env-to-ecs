package args

import (
	"errors"
	"flag"
)

type ArgConfig struct {
	InFile   string
	OutFile  string
	Variable string
}

var argConfig ArgConfig

func GetArguments() (ArgConfig, error) {
	flag.Parse()

	if argConfig.InFile == "" {
		return ArgConfig{}, errors.New("infile cannot be empty")
	}

	return ArgConfig{argConfig.InFile, argConfig.OutFile, argConfig.Variable}, nil
}

func init() {
	parseArgsModelAndSetFlag(&argConfig.InFile, inFileArgs)
	parseArgsModelAndSetFlag(&argConfig.OutFile, outFileArgs)
	parseArgsModelAndSetFlag(&argConfig.Variable, variableArgs)
}

func parseArgsModelAndSetFlag(flagVar *string, argsModel map[string]string) {
	flag.StringVar(flagVar, argsModel["shortFlag"], argsModel["defaultValue"], argsModel["usage"])
	flag.StringVar(flagVar, argsModel["longFlag"], argsModel["defaultValue"], argsModel["usage"])
}
