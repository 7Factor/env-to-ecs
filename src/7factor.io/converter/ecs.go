package converter

import (
	"bytes"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

type Pair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func ConvertInputToJson(contents []string) (string, error) {
	if len(contents) == 0 {
		return `[]`, errors.New("contents cannot be empty")
	}
	itemsToBeParsed := handleInputSlice(contents)

	pairs := splitOnEquals(itemsToBeParsed)
	return pairsToString(pairs)
}

var assignmentRegex = regexp.MustCompile("\\w+ *= *(?:'[^']*'|\"[^\"]*\"|[^\\s=]*)=?")

func handleInputSlice(contents []string) []string {
	var itemsToBeParsed []string
	for _, line := range contents {
		line = stripComments(line)

		if hasEmptyString(line) {
			continue
		}

		items := assignmentRegex.FindAllString(line, -1)

		itemsToBeParsed = append(itemsToBeParsed, items...)
	}

	return itemsToBeParsed
}

var commentRegex = regexp.MustCompile("^(?m)\\s*#.*\\n*$")

func stripComments(line string) string {
	return commentRegex.ReplaceAllLiteralString(line, "")
}

func hasEmptyString(line string) bool {
	return len(line) == 0 || line == ""
}

func splitOnEquals(slice []string) []Pair {
	var splitOnEquals []string
	var pairs []Pair

	for _, item := range slice {
		splitOnEquals = strings.SplitN(item, "=", 2)
		pairs = append(pairs, Pair{
			Name:  strings.TrimSpace(splitOnEquals[0]),
			Value: trimQuotes(strings.TrimSpace(splitOnEquals[1])),
		})
	}

	return pairs
}

func pairsToString(pairs []Pair) (string, error) {
	buffer := bytes.Buffer{}
	err := json.NewEncoder(&buffer).Encode(pairs)

	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(buffer.String()), nil
	}
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
