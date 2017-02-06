package main

import (
	"log"
	"net/http"
	"wordRecorderService/letterCounting"
	"wordRecorderService/wordCounting"
)

var letterCounter = letterCounting.NewLetterCounter()
var wordCounter = wordCounting.NewWordCounter()

func main() {
	statsMux := http.NewServeMux()
	statsMux.HandleFunc("/stats", getStats)
	go http.ListenAndServe(":8081", statsMux)

	log.Println("server starting")

	inputMux := http.NewServeMux()
	inputMux.HandleFunc("/put", putSentence)
	log.Fatalln(http.ListenAndServe(":8080", inputMux))
}
