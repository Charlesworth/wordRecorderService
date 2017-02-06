package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", 0)
}

func putSentence(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	logger.Println(r.RemoteAddr, "Put Request")

	if err != nil {
		w.WriteHeader(500)
		logger.Println(r.RemoteAddr, "Put Request Error 500:", err)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(400)
		logger.Println(r.RemoteAddr, "Put Request Error 400: no request body supplied")
		return
	}

	words := getWords(string(body))
	for _, word := range words {
		letterCounter.AddWord(word)
		wordCounter.AddWord(word)
	}

	w.WriteHeader(202)
	return
}

type Stats struct {
	Count         int
	Top_5_words   []string
	Top_5_letters []string
}

func getStats(w http.ResponseWriter, r *http.Request) {
	logger.Println(r.RemoteAddr, "Get /stats Request")

	count := wordCounter.GetCount()
	topWords := wordCounter.GetMostFrequentFive()
	topLetters := letterCounter.GetMostFrequentFive()

	stats := Stats{
		Count:         count,
		Top_5_words:   topWords,
		Top_5_letters: topLetters,
	}

	jsonStats, err := json.Marshal(stats)
	if err != nil {
		w.WriteHeader(500)
		logger.Println(r.RemoteAddr, "Put Request Error 500: unable to marshal stats to JSON")
		return
	}

	w.Write(jsonStats)
	w.WriteHeader(200)
	return
}
