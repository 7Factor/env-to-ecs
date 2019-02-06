package converter

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndConvert(inFile string, outFile string) (string, error) {
	// verify inFile exists
	_, err := os.Stat(inFile)
	if err != nil {
		return "", fmt.Errorf("INFILE not found")
	}

	// parse inFile
	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		return "", fmt.Errorf("unable to read INFILE, catestrophic error")
	}

	// transform contents
	transformedContents, err := TransformAndTranslate(string(contents))

	// write to outFile
	if outFile == "stdout" {
		fmt.Fprint(os.Stdout, transformedContents)
	} else {
		// WriteFile will create the file if it does not exist
		if err != nil {
			return "", fmt.Errorf("error while transforming contents")
		}
		n1 := []byte(transformedContents)
		err = ioutil.WriteFile(outFile, n1, 0644)
	}

	return outFile, nil
}
