package converter

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAndConvert(path string) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("file not found")
	}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("unable to read file, catestrophic error")
	} else {
		return TransformAndTranslate(string(contents))
	}
}
