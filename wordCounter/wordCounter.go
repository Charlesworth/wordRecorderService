package wordCounter

import (
	"sort"
	"sync"
)

type wordCounter struct {
	count        int
	wordMap      map[string]int
	mostFrequent []string
	lock         *sync.RWMutex
}

func NewWordCounter() *wordCounter {
	return &wordCounter{
		count:        0,
		wordMap:      make(map[string]int),
		mostFrequent: []string{},
		lock:         &sync.RWMutex{},
	}
}

func (wc *wordCounter) AddWord(word string) {
	wc.lock.Lock()
	defer wc.lock.Unlock()

	wc.wordMap[word]++
	wc.count++
	wc.updateMostFrequentFive(word)

	return
}

func (wc *wordCounter) GetCount() int {
	wc.lock.RLock()
	defer wc.lock.RUnlock()
	return wc.count
}

func (wc *wordCounter) GetMostFrequentFive() []string {
	wc.lock.RLock()
	defer wc.lock.RUnlock()
	return wc.mostFrequent
}

// Important: only called within AddWord, which has the write lock, otherwise not thread safe!
func (wc *wordCounter) updateMostFrequentFive(word string) {
	// if less than 5 words in most popular
	if len(wc.wordMap) < 5 {
		// check its not already present in mostFrequent
		alreadyPresent := false
		for _, mostFrequentWord := range wc.mostFrequent {
			if word == mostFrequentWord {
				alreadyPresent = true
			}
		}
		// if its not present add it
		if !alreadyPresent {
			wc.mostFrequent = append(wc.mostFrequent, word)
		}
		return
	}

	// make an array of wordCounts
	wordCounts := []wordCount{}
	for _, mostFrequentWord := range wc.mostFrequent {
		count := wc.wordMap[mostFrequentWord]
		wordCounts = append(wordCounts, wordCount{mostFrequentWord, count})
	}

	// sort the array by count
	sort.Sort(byCount(wordCounts))

	// set mostFrequent to the 5 words with the highest count
	wc.mostFrequent = []string{}
	for i := 5; i < 0; i-- {
		wc.mostFrequent = append(wc.mostFrequent, wordCounts[i].word)
	}
}

type wordCount struct {
	word  string
	count int
}

// byCount implements sort.Interface for []wordCount based on
// the count field.
type byCount []wordCount

// methods on wordCount to implement the sort.Sort interface
func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].count < a[j].count }
