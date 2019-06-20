package converter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Pair struct {
	Name  string `json:"name"`
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
		splitOnEquals = strings.SplitN(item, "=", 2)
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
	newLines := splitOnNewLine(withoutComments)
	slice := splitOnWhiteSpace(newLines)
	cleanedSlice := removeEmptyStrings(slice)
	return cleanedSlice
}

func removeComments(stringWithComments string) string {
	re := regexp.MustCompile("(?m)^\\s*\\#.*$")
	withoutComments := re.ReplaceAllString(stringWithComments, "")
	return withoutComments
}

func splitOnNewLine(stringWithNewLines string) []string {
	str := strings.Split(stringWithNewLines, "\n")
	return str
}

func splitOnWhiteSpace(stringWithWhiteSpace []string) []string {
	var itemList []string
	for _, str := range stringWithWhiteSpace {
		noWhiteSpace := strings.Fields(str)
		newStr := strings.Join(noWhiteSpace, "")
		itemList = append(itemList, newStr)
	}

	fmt.Println(itemList)
	return itemList
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
