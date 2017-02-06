package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/charlesworth/wordRecorderService/letterCounting"
	"github.com/charlesworth/wordRecorderService/wordCounting"
)

type behavioralAPITableTest struct {
	testDescription string
	putSentences    []string
	expectedStats   Stats
}

var behavioralAPITests = []behavioralAPITableTest{
	behavioralAPITableTest{
		testDescription: "putting a single word sentence should return only that word in stats",
		putSentences: []string{
			"aaaaa",
		},
		expectedStats: Stats{
			Count:         1,
			Top_5_words:   []string{"aaaaa"},
			Top_5_letters: []string{"a"},
		},
	},
	behavioralAPITableTest{
		testDescription: "putting a single multiple word sentence should return correct stats",
		putSentences: []string{
			"aa bb cc",
		},
		expectedStats: Stats{
			Count:         3,
			Top_5_words:   []string{"aa", "bb", "cc"},
			Top_5_letters: []string{"a", "b", "c"},
		},
	},
	behavioralAPITableTest{
		testDescription: "putting words with suffix punctuation, the punctuation will be removed",
		putSentences: []string{
			"aa, bb.",
		},
		expectedStats: Stats{
			Count:         2,
			Top_5_words:   []string{"aa", "bb"},
			Top_5_letters: []string{"a", "b"},
		},
	},
	behavioralAPITableTest{
		testDescription: "putting multiple sentences should return correct stats",
		putSentences: []string{
			"aa",
			"bb",
		},
		expectedStats: Stats{
			Count:         2,
			Top_5_words:   []string{"aa", "bb"},
			Top_5_letters: []string{"a", "b"},
		},
	},
}

// Behavioral tests for the http API
func TestAPITest(t *testing.T) {
	testServerPut := httptest.NewServer(http.HandlerFunc(putSentence))
	defer testServerPut.Close()
	testServerGet := httptest.NewServer(http.HandlerFunc(getStats))
	defer testServerGet.Close()
	client := &http.Client{}

TestLoop:
	for _, test := range behavioralAPITests {
		t.Log("Test:", test.testDescription)
		letterCounter = letterCounting.NewLetterCounter()
		wordCounter = wordCounting.NewWordCounter()

		for _, sentence := range test.putSentences {
			req, _ := http.NewRequest("PUT", testServerPut.URL, bytes.NewBuffer([]byte(sentence)))
			resp, err := client.Do(req)
			if err != nil {
				t.Error("Test Setup Failure: unable to make PUT request to the test server:", err)
				continue TestLoop
			}
			if resp.StatusCode != 202 {
				t.Error("Error: PUT request with should return 202, returned", resp.StatusCode)
				continue TestLoop
			}
		}

		req, _ := http.NewRequest("GET", testServerGet.URL+"/stats", nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Error("Test Setup Failure: unable to make GET request to the test server:", err)
			continue TestLoop
		}
		if resp.StatusCode != 200 {
			t.Error("Error: stats request should return 200, returned", resp.StatusCode)
			continue TestLoop
		}

		var wordStats Stats
		bodyByte, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(bodyByte, &wordStats)
		if err != nil {
			t.Error("Test Setup Failure: unable to unmarshal GET /stats body")
		}

		if !sliceEqual(test.expectedStats.Top_5_letters, wordStats.Top_5_letters) {
			t.Error("Error: stats Top_5_letters [", wordStats.Top_5_letters, "] is not equal to expected output [", test.expectedStats.Top_5_letters, "]")
		}
		if !sliceEqual(test.expectedStats.Top_5_words, wordStats.Top_5_words) {
			t.Error("Error: stats Top_5_words [", wordStats.Top_5_words, "] is not equal to expected output [", test.expectedStats.Top_5_words, "]")
		}
		if test.expectedStats.Count != wordStats.Count {
			t.Error("Error: stats count [", wordStats.Count, "] is not equal to expected output [", test.expectedStats.Count, "]")
		}
	}
}

func sliceEqual(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	for _, i := range a {
		contains := false
		for _, j := range b {
			if i == j {
				contains = true
			}
		}
		if !contains {
			return false
		}
	}

	return true
}
