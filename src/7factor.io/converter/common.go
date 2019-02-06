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
		return "", fmt.Errorf("inFile not found")
	}

	// parse inFile
	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		return "", fmt.Errorf("unable to read inFile, catestrophic error")
	}

	// write to outFile
	// WriteFile will create the file if it does not exist
	transformedContents, err := TransformAndTranslate(string(contents))
	if err != nil {
		return "", fmt.Errorf("error while transforming contents")
	}
	n1 := []byte(transformedContents)
	err = ioutil.WriteFile(outFile, n1, 0644)

	return outFile, nil
}
