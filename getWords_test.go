package main

import "testing"

type getWordsTableTest struct {
	testDescrip    string
	input          string
	expectedOutput []string
}

var getWordsTests = []getWordsTableTest{
	getWordsTableTest{
		testDescrip:    "Test: getwords with empty string should return empty string array",
		input:          "",
		expectedOutput: []string{},
	},
	getWordsTableTest{
		testDescrip:    "Test: getwords string with single word should return string slice with one word",
		input:          "hi",
		expectedOutput: []string{"hi"},
	},
	getWordsTableTest{
		testDescrip:    "Test: getwords string with multiple words should return all words in string slice",
		input:          "hi bye",
		expectedOutput: []string{"hi", "bye"},
	},
	getWordsTableTest{
		testDescrip:    "Test: getwords should decapitalize all letters",
		input:          "HI",
		expectedOutput: []string{"hi"},
	},
	getWordsTableTest{
		testDescrip:    "Test: getwords should remove suffix punctuation",
		input:          "hi, bye",
		expectedOutput: []string{"hi", "bye"},
	},
	getWordsTableTest{
		testDescrip:    "Test: getwords should not count words that do not start with a letter",
		input:          "$hi ,bye",
		expectedOutput: []string{},
	},
}

func TestGetWords(t *testing.T) {
	for _, test := range getWordsTests {
		t.Log(test.testDescrip)
		results := getWords(test.input)
		for i, result := range results {
			if test.expectedOutput[i] != result {
				t.Error("Error: expected:", test.expectedOutput, " but result was:", results)
			}
		}
	}
}
