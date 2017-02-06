package wordCounting

import "testing"

func TestGetCount(t *testing.T) {
	t.Log("Test: getCount on empty wordCounter should return 0")
	wordCounter := NewWordCounter()
	count := wordCounter.GetCount()
	if count != 0 {
		t.Error("Error: did not return 0, instead returned:", count)
	}

	t.Log("Test: getCount on single word wordCounter should return 1")
	wordCounter = NewWordCounter()
	wordCounter.AddWord("hi")
	count = wordCounter.GetCount()
	if count != 1 {
		t.Error("Error: did not return 1, instead returned:", count)
	}

	t.Log("Test: getCount on multiple word wordCounter should return correct count")
	wordCounter = NewWordCounter()
	wordCounter.AddWord("hi")
	wordCounter.AddWord("hello")
	wordCounter.AddWord("test")
	wordCounter.AddWord("bye")
	count = wordCounter.GetCount()
	if count != 4 {
		t.Error("Error: did not return 4, instead returned:", count)
	}
}

func TestAddWord(t *testing.T) {
	t.Log("Test: addWord on single word should add it and update count and mostFrequent")
	wordCounter := NewWordCounter()
	wordCounter.AddWord("hi")
	if wordCounter.GetCount() != 1 {
		t.Error("Error: count returned from single add word did not equal 1:", wordCounter.GetCount())
	}
	if wordCounter.mostFrequent[0] != "hi" {
		t.Error("Error: after addWord, incorrect word returned from mostFrequent")
	}

	t.Log("Test: addWord on multiple words should add them and update count and mostFrequent")
	wordCounter = NewWordCounter()
	wordCounter.AddWord("hi")
	wordCounter.AddWord("hi")
	wordCounter.AddWord("bye")
	if wordCounter.GetCount() != 3 {
		t.Error("Error: count returned from multiple addWord was incorrect, expected 3, recieved", wordCounter.GetCount())
	}
	if (wordCounter.mostFrequent[0] != "hi") && (wordCounter.mostFrequent[1] != "bye") {
		t.Error("Error: after addWord, incorrect words returned from mostFrequent")
	}
}

func TestGetMostFrequentFive(t *testing.T) {
	t.Log("Test: GetMostFrequentFive on empty wordCounter should return empty string slice")
	wordCounter := NewWordCounter()
	if len(wordCounter.mostFrequent) != 0 {
		t.Error("Error: returned a non empty string slice")
	}

	t.Log("Test: GetMostFrequentFive on wordCounter with less than 5 words should return those words")
	wordCounter = NewWordCounter()
	testWords := []string{"hi", "bye", "test"}
	wordCounter.AddWord(testWords[0])
	wordCounter.AddWord(testWords[1])
	wordCounter.AddWord(testWords[0])
	wordCounter.AddWord(testWords[2])
	if len(wordCounter.GetMostFrequentFive()) != 3 {
		t.Error("Error: returned inncorrect length string slice")
	}
	testFail := false
	for _, testWord := range testWords {
		if !contains(wordCounter.GetMostFrequentFive(), testWord) {
			testFail = true
		}
	}
	if testFail {
		t.Error("Error: GetMostFrequentFive return slice did not contain correct words")
	}

	t.Log("Test: GetMostFrequentFive on wordCounter with more than 5 words")
	wordCounter = NewWordCounter()
	testMostFrequentWords := []string{"hi", "bye", "test", "test2", "test3"}
	wordCounter.AddWord(testMostFrequentWords[0])
	wordCounter.AddWord(testMostFrequentWords[0])
	wordCounter.AddWord(testMostFrequentWords[1])
	wordCounter.AddWord(testMostFrequentWords[1])
	wordCounter.AddWord(testMostFrequentWords[2])
	wordCounter.AddWord(testMostFrequentWords[2])
	wordCounter.AddWord(testMostFrequentWords[3])
	wordCounter.AddWord(testMostFrequentWords[3])
	wordCounter.AddWord(testMostFrequentWords[4])
	wordCounter.AddWord(testMostFrequentWords[4])
	wordCounter.AddWord("dontReturnMe")
	if len(wordCounter.GetMostFrequentFive()) != 5 {
		t.Error("Error: returned inncorrect length string slice")
	}
	testFail = false
	t.Log(wordCounter.GetMostFrequentFive())
	for _, testWord := range testMostFrequentWords {
		if !contains(wordCounter.GetMostFrequentFive(), testWord) {
			testFail = true
		}
	}
	if testFail {
		t.Error("Error: GetMostFrequentFive return slice did not contain correct words")
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
