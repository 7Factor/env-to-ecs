package converter

import (
	. "7factor.io/args"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndConvert(config ArgConfig) (string, error) {
	contents, err := parseInfileOrPanic(config.InFile)
	if err != nil {
		return "", err
	}

	withExtraVars := concatExtraVars(contents, config.Variables)

	transformedContents, err := ConvertInputToJson(withExtraVars)
	if err != nil {
		return "", fmt.Errorf("caught error while attempting to transform contents: %s\n", err)
	}

	err = writeToOutFile(config.OutFile, transformedContents)
	if err != nil {
		return "", fmt.Errorf("caught error while attempting to tranform contents: %s\n", err)
	}

	return config.OutFile, nil
}

func parseInfileOrPanic(infile string) (string, error) {
	_, err := os.Stat(infile)
	if err != nil {
		return "", fmt.Errorf("caught error while looking up file: %s\n", err)
	}
	contents, err := ioutil.ReadFile(infile)
	if err != nil {
		return "", fmt.Errorf("catestrophic faliure while attempting to read infile: %s\n", err)
	}
	return string(contents), nil
}

func concatExtraVars(toConcat string, extraVars []string) string {
	var concatedString = toConcat
	for i := range extraVars {
		concatedString += "\n" + extraVars[i]
	}
	return concatedString
}

func writeToOutFile(outFile string, transformedContents string) error {
	var err error
	if outFile == "stdout" {
		fmt.Fprint(os.Stdout, transformedContents)
	} else {
		n1 := []byte(transformedContents)
		err = ioutil.WriteFile(outFile, n1, 0644)
	}
	return err
}
