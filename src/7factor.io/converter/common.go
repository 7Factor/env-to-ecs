package converter

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndConvert(path string) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf( "File not found.")
	}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Unable to read file. Catestrophic error!")
	} else {
		return string(contents), nil
	}
}