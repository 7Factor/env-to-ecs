package converter

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndConvert(inFile string, outFile string) (string, error) {
	_, err := os.Stat(inFile)
	if err != nil {
		return "", fmt.Errorf("INFILE not found")
	}

	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		return "", fmt.Errorf("unable to read INFILE, catestrophic error")
	}

	transformedContents, err := TransformAndTranslate(string(contents))
	if err != nil {
		return "", fmt.Errorf("error while transforming contents")
	}

	err = writeToOutFile(outFile, transformedContents)
	if err != nil {
		return "", fmt.Errorf("unable to write to OUTFILE")
	}

	return outFile, nil
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
