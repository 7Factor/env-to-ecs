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
		return "", fmt.Errorf("file not found")
	}

	// parse inFile
	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		return "", fmt.Errorf("unable to read file, catestrophic error")
	}

	// create outFile if it does not exist
	_, err = os.Stat(outFile)
	if os.IsNotExist(err) {
		os.Create(outFile)
	}

	return TransformAndTranslate(string(contents))

}
