package main

import "strings"

/*
word filtering rules:
- remove suffix ',' and '.'
- all letters to lower case
- must start with a letter
*/

func getWords(sentance string) []string {
	// split the sentance on spaces to give words
	unfilteredWords := strings.Split(sentance, " ")
	filteredWords := []string{}

	// for each word, sanitise and add to filteredWords if starts with a letter
	for _, word := range unfilteredWords {
		word = strings.ToLower(word)
		if startWithLetter(word) {
			word = removeSuffixPunctuation(word)
			filteredWords = append(filteredWords, word)
		}
	}

	return filteredWords
}

func removeSuffixPunctuation(word string) string {
	word = strings.TrimSuffix(word, ".")
	word = strings.TrimSuffix(word, ",")
	return word
}

func startWithLetter(word string) bool {
	if len(word) == 0 {
		return false
	}

	r := word[0]
	return (97 <= int(r)) && (int(r) < 122)
}
