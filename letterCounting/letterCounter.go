package letterCounting

import (
	"sort"
	"sync"
)

type letterCounter struct {
	letterMap         map[rune]int
	lock              *sync.RWMutex
	mostFrequentCache []string
}

// NewLetterCounter returns a pointer to a new letterCounter
func NewLetterCounter() *letterCounter {
	return &letterCounter{
		letterMap: make(map[rune]int),
		lock:      &sync.RWMutex{},
	}
}

// AddWord adds a string's component letters to the letterCounter
func (lc *letterCounter) AddWord(word string) {
	lc.lock.Lock()
	defer lc.lock.Unlock()
	for _, r := range word {
		if isLowerCaseLetter(r) {
			lc.letterMap[r]++
		}
	}
	// invalidate the mostFrequentCache
	lc.mostFrequentCache = nil
}

// GetMostFrequentFive returns the five most frequently used letters
func (lc *letterCounter) GetMostFrequentFive() []string {
	lc.lock.RLock()
	defer lc.lock.RUnlock()

	// if cache result has not been invalidated, return it
	if lc.mostFrequentCache != nil {
		return lc.mostFrequentCache
	}

	// if letterCounter is empty return empty []string
	if len(lc.letterMap) == 0 {
		return []string{}
	}

	// if 5 or less letters, return those without sorting
	if len(lc.letterMap) < 6 {
		var mostFrequent []string
		for letterRune := range lc.letterMap {
			mostFrequent = append(mostFrequent, string(letterRune))
		}
		return mostFrequent
	}

	// make an array of letterCounts
	letterCounts := []letterCount{}
	for letterRune, count := range lc.letterMap {
		letterCounts = append(letterCounts, letterCount{string(letterRune), count})
	}

	// sort the array by count
	sort.Sort(byCount(letterCounts))

	// return the 5 letters with the biggest count
	mostFrequentFive := []string{}
	for i := len(letterCounts) - 5; i < len(letterCounts); i++ {
		mostFrequentFive = append(mostFrequentFive, letterCounts[i].letter)
	}
	lc.mostFrequentCache = mostFrequentFive
	return mostFrequentFive
}

// letterCount is used to sort the frequency of the letters
// for the GetMostFrequentFive method
type letterCount struct {
	letter string
	count  int
}

// byCount implements sort.Interface for []letterCount based on
// the count field.
type byCount []letterCount

// methods on letterCount to implement the sort.Sort interface
func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].count < a[j].count }

func isLowerCaseLetter(r rune) bool {
	return (97 <= int(r)) && (int(r) < 122)
}
