package converter

import (
	"bytes"
	"encoding/json"
	"errors"
	"regexp"
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

	cleanedContents := cleanContents(contents)

	pairs := transform(cleanedContents)

	return translate(pairs)
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

func cleanContents(dirtyString string) []string {
	withoutComments := removeComments(dirtyString)
	noWhiteSpace := removeWhiteSpace(withoutComments)
	cleanedSlice := removeEmptyStrings(noWhiteSpace)

	return cleanedSlice
}

func removeComments(stringWithComments string) string {
	re := regexp.MustCompile("(?m)[\r\n]+^.*#.*$")
	withoutComments := re.ReplaceAllString(stringWithComments, "")

	return withoutComments
}

func removeWhiteSpace(stringWithWhiteSpace string) []string {
	noWhiteSpace := strings.Fields(stringWithWhiteSpace)

	return noWhiteSpace
}

func removeEmptyStrings(sliceWithEmptyStrings []string) []string {
	var noEmptyStrings []string
	for _, str := range sliceWithEmptyStrings {
		if str != "" {
			noEmptyStrings = append(noEmptyStrings, str)
		}
	}

	return noEmptyStrings
}
