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

func TransformAndTranslate(contents string) (string, error) {
	if contents == "" {
		return `[]`, errors.New("contents cannot be empty")
	}

	cleanedSlice := removeEmptyStrings(strings.Split(contents,"\n"))

	pairs := transform(cleanedSlice)

	return translate(pairs)
}

func removeEmptyStrings(slice []string) []string {
	var cleanedSlice []string
	for _, str := range slice {
		if str != "" {
			cleanedSlice = append(cleanedSlice, str)
		}
	}
	return cleanedSlice
}

func transform(slice []string) []Pair {
	var splitOnEquals []string
	var pairs []Pair

	for _, item := range slice {
		splitOnEquals = strings.Split(item,"=")
		pairs = append(pairs, Pair{
			Name:  splitOnEquals[0],
			Value: splitOnEquals[1],
		})
	}

	return pairs
}

func translate(pairs []Pair) (string, error) {
	buffer := bytes.Buffer{}
	err := json.NewEncoder(&buffer).Encode(pairs)

	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(buffer.String()), nil
	}
}
