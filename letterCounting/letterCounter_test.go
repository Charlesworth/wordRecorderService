package letterCounting

import "testing"

func TestNewLetterCounter(t *testing.T) {
	letterCounter := NewLetterCounter()
	if letterCounter == nil {
		t.Error("NewLetterCounter() returns an uninitialised letterCounter")
	}
}

func TestAddWord(t *testing.T) {
	t.Log("Test: AddWord with empty string adds no letters to letterCounter")
	testLetterCounter := NewLetterCounter()
	testLetterCounter.AddWord("")
	if len(testLetterCounter.letterMap) != 0 {
		t.Error("Error: AddWord added to letterCounter when an empty string was passed")
	}

	t.Log("Test: AddWord with single character string adds single letter to letterCounter")
	testLetterCounter = NewLetterCounter()
	testLetterCounter.AddWord("a")
	if len(testLetterCounter.letterMap) != 1 {
		t.Error("Error: AddWord added ", len(testLetterCounter.letterMap), " letters to letterCounter when one was supplied")
	}
	if testLetterCounter.letterMap[rune(0x61)] != 1 {
		t.Error("Error: AddWord(a) did not add the rune 'a' to the letterMap")
	}

	t.Log("Test: AddWord with word with repeated characters adds correctly to letterCounter")
	testLetterCounter = NewLetterCounter()
	testLetterCounter.AddWord("aaab")
	if len(testLetterCounter.letterMap) != 2 {
		t.Error("Error: AddWord added ", len(testLetterCounter.letterMap), " letters to letterCounter when 2 were supplied")
	}
	if testLetterCounter.letterMap[rune(0x61)] != 3 {
		t.Error("Error: AddWord(aaab) did not add the rune 'a' to the letterMap 3 times")
	}
}

func TestGetMostFrequentFive(t *testing.T) {
	t.Log("Test: GetMostFrequentFive on empty letterCounter returns no letters")
	testLetterCounter := NewLetterCounter()
	mostFrequentFive := testLetterCounter.GetMostFrequentFive()
	if len(mostFrequentFive) != 0 {
		t.Error("Error: returned []string was not empty")
	}

	t.Log("Test: GetMostFrequentFive with <5 letters entered returns only those letters")
	testLetterCounter = NewLetterCounter()
	testLetterCounter.AddWord("aaabbc")
	mostFrequentFive = testLetterCounter.GetMostFrequentFive()
	if len(mostFrequentFive) != 3 {
		t.Error("Error: returned []string did not contain 3 letters")
	}

	t.Log("Test: GetMostFrequentFive with equal number of letters returns correctly")
	testLetterCounter = NewLetterCounter()
	testLetterCounter.AddWord("abcdefghij")
	mostFrequentFive = testLetterCounter.GetMostFrequentFive()
	if len(mostFrequentFive) != 5 {
		t.Error("Error: returned []string did not contain 5 letters")
	}

	t.Log("Test: GetMostFrequentFive with mixed numbers of letters returns correct five")
	testLetterCounter = NewLetterCounter()
	testLetterCounter.AddWord("aaaabbbcccddeefghij")
	testMostFrequentFive := testLetterCounter.GetMostFrequentFive()
	// check that 5 were returned
	if len(testMostFrequentFive) != 5 {
		t.Error("Error: returned []string did not contain 5 letters")
	}

	// check that its the correct 5
	actualMostFrequentFive := []string{"a", "b", "c", "d", "e"}
	testFailed := false
	for _, letter := range actualMostFrequentFive {
		if !contains(testMostFrequentFive, letter) {
			testFailed = true
		}
	}
	if testFailed {
		t.Error("Error: returned letter are not the most frequent 5")
	}

	t.Log("Test: GetMostFrequentFive with a cached result should return cached result")
	testLetterCounter = NewLetterCounter()
	testLetterCounter.AddWord("zzzzz")
	testLetterCounter.mostFrequentCache = []string{"a"}
	mostFrequent := testLetterCounter.GetMostFrequentFive()
	if mostFrequent[0] != "a" {
		t.Error("Error: letterCounter did not return chached result")
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
