package converter

import "errors"

func Transform(contents string) (string, error) {
	if contents == "" {
		return `[]`, errors.New("contents cannot be empty")
	}

	return contents, nil
}