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


func TransformAndTranslate(contents string) (string, error) {
	if contents == "" {
		return `[]`, errors.New("contents cannot be empty")
	}

	var itemsToBeParsed []string
	for _, line := range strings.Split(contents, "\n") {
		if hasComment(line) || hasEmptyString(line) {
			continue
		}
		items := parse(line)
		

		itemsToBeParsed = append(itemsToBeParsed, items...)
	}

	pairs := transform(itemsToBeParsed)
	return translate(pairs)
}

var commentRegex = regexp.MustCompile("(?m)^\\s*\\#.*$")
func hasComment(line string) bool {
	return commentRegex.MatchString(line)
}

func hasEmptyString(line string) bool {
	return len(line) == 0 || line == ""
}

var singleValueContextRegex = regexp.MustCompile("^.+?=.+?$")
func parse(line string) []string {
	defer func() {
		// Clear out newItems for next call
		newItems = nil
	}()
	// There could be multiple values on a single line, so take that into consideration.
	items := strings.Fields(line)
	for i := 0; i < len(items); i++ {
		item := items[i]
		
		// Try to handle items with regular quotes i.e A="B"
		if handleQuotes(item) {
			continue
		}
		
		// Try to handle items with "Key=Value"
		if singleValueContextRegex.MatchString(item) {
			newItems = append(newItems, item)
			continue
		}

		// Try to predict the next spaced key/val pair to see if it has quotes.
		if handleSpacedQuotes(item, items, i) {
			i++
			continue
		}

		// Try and pull next 2 items and append together to create a statement (i.e X=1).
		newItems = append(newItems, item+items[i+1]+items[i+2])
		
		// Skip the next two items, because we've already inserted them into our item list.
		i+=2
	}
	
	return newItems
}

var newItems []string
var tempSpacedItem string
var quotedItems []string

var parsingQuotes = false
var parsingSpacedQuotes = false
func handleQuotes(item string) bool {
	if strings.Contains(item, "\"") && !strings.HasSuffix(item, "\"") {
		parsingQuotes = true
	}

	if parsingQuotes {
		quotedItems = append(quotedItems, item)
		if strings.HasSuffix(item, "\"") {
			if parsingSpacedQuotes {
				tempSpacedItem += strings.Join(quotedItems, " ")
				newItems = append(newItems, tempSpacedItem)
			} else {
				newItems = append(newItems, strings.Join(quotedItems, " "))
			}
			resetItems()
		}
		return true
	}
	return  false
}

func handleSpacedQuotes(item string, items []string, iterator int) bool {
	if strings.Contains(items[iterator+2], "\"") && !strings.HasSuffix(items[iterator+2], "\"") {
		parsingQuotes = true
		parsingSpacedQuotes = true
		tempSpacedItem = item + items[iterator+1]
		return true
	}
	return false
}

func resetItems() {
	parsingSpacedQuotes = false
	parsingQuotes = false
	tempSpacedItem = ""
	quotedItems = nil
}

func transform(slice []string) []Pair {
	var splitOnEquals []string
	var pairs []Pair

	for _, item := range slice {
		splitOnEquals = strings.SplitN(item, "=", 2)
		pairs = append(pairs, Pair{
			Name:  splitOnEquals[0],
			Value: trimQuotes(splitOnEquals[1]),
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

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}