package converter

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

type Pair struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

func Transform(contents string) (string, error) {
	if contents == "" {
		return `[]`, errors.New("contents cannot be empty")
	}

	var pairs []Pair
	split := strings.Split(contents,"=")
	pairs = append(pairs, Pair{
		Name: split[0],
		Value: split[1],
	})

	return Translate(pairs)
}

func Translate(pairs []Pair) (string, error) {
	buffer := bytes.Buffer{}
	err := json.NewEncoder(&buffer).Encode(pairs)

	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(buffer.String()), nil
	}
}