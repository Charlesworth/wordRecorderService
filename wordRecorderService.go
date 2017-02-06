package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/charlesworth/wordRecorderService/letterCounting"
	"github.com/charlesworth/wordRecorderService/wordCounting"
)

var letterCounter = letterCounting.NewLetterCounter()
var wordCounter = wordCounting.NewWordCounter()

func main() {
	log.Println("Word Recorder Service")
	statsMux := http.NewServeMux()
	statsMux.HandleFunc("/stats", getStats)
	statsPortStr := ":" + strconv.Itoa(*statsPort)
	log.Println("Stats port", statsPortStr)
	go http.ListenAndServe(statsPortStr, statsMux)

	inputMux := http.NewServeMux()
	inputMux.HandleFunc("/", putSentence)
	inputPortStr := ":" + strconv.Itoa(*inputPort)
	log.Println("Input port", inputPortStr)
	log.Fatalln(http.ListenAndServe(inputPortStr, inputMux))
}
