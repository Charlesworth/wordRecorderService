package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"wordRecorderService/letterCounting"
	"wordRecorderService/wordCounting"
)

func TestPutSentence(t *testing.T) {
	t.Log("Test: request with no body should return 400")
	letterCounter = letterCounting.NewLetterCounter()
	wordCounter = wordCounting.NewWordCounter()

	testServer := httptest.NewServer(http.HandlerFunc(putSentence))

	req, _ := http.NewRequest("PUT", testServer.URL, bytes.NewBuffer([]byte("")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Test Setup Failure: unable to make request to the test server:", err)
	}

	if resp.StatusCode != 400 {
		t.Error("Error: request with no body should return 400, returned", resp.StatusCode)
	}

	mostFrequentLetters := letterCounter.GetMostFrequentFive()
	if len(mostFrequentLetters) != 0 {
		t.Error("Error: request with no body should not add to the letter counter")
	}
	mostFrequentWords := wordCounter.GetMostFrequentFive()
	if len(mostFrequentWords) != 0 {

	}
	testServer.Close()

	t.Log("Test: request with valid sentence body should return 202")
	letterCounter = letterCounting.NewLetterCounter()
	wordCounter = wordCounting.NewWordCounter()

	testServer = httptest.NewServer(http.HandlerFunc(putSentence))

	req, _ = http.NewRequest("PUT", testServer.URL, bytes.NewBuffer([]byte("aaa bbb")))
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Error("Test Setup Failure: unable to make request to the test server:", err)
	}

	mostFrequentLetters = letterCounter.GetMostFrequentFive()
	for _, letter := range mostFrequentLetters {
		if (letter != "a") && (letter != "b") {
			t.Error("Error: letterCounter contains unsent letters")
		}
	}

	mostFrequentWords = wordCounter.GetMostFrequentFive()
	if len(mostFrequentWords) != 2 {
		t.Error("Error: letterCounter contains unsent letters")
	}

	if resp.StatusCode != 202 {
		t.Error("Error: HTTP status code was not 202:", resp.StatusCode)
	}
	testServer.Close()
}

func TestGetStats(t *testing.T) {
	t.Log("Test: request for stats should return stats JSON and 200")
	letterCounter = letterCounting.NewLetterCounter()
	wordCounter = wordCounting.NewWordCounter()
	wordCounter.AddWord("abc")
	letterCounter.AddWord("abc")

	testServer := httptest.NewServer(http.HandlerFunc(getStats))

	req, _ := http.NewRequest("GET", testServer.URL, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Test Setup Failure: unable to make request to the test server:", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Error: stats request should return 200, returned", resp.StatusCode)
	}

	var wordStats Stats
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyByte, &wordStats)
	if err != nil {
		t.Error("Test Setup Failure: unable to unmarshal GET /stats body")
	}

	if !sliceEquality(wordStats.Top_5_letters, []string{"a", "b", "c"}) {
		t.Error()
	}

	if !sliceEquality(wordStats.Top_5_words, []string{"abc"}) {
		t.Error()
	}

	resp.Body.Close()
	testServer.Close()
}

func sliceEquality(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	for _, i := range a {
		test := false
		for _, j := range b {
			if i == j {
				test = true
			}
		}
		if !test {
			return false
		}
	}

	return true
}
