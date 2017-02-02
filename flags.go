package main

import (
	"errors"
	"flag"
	"log"
)

var inputPort = flag.Int("inputPort", 5555, "port used for word input over http")
var statsPort = flag.Int("statsPort", 8080, "port used to call /stats")

func init() {
	flag.Parse()
	err := checkFlags()
	if err != nil {
		log.Fatal(err)
	}
}

func checkFlags() error {
	if (*inputPort < 1024) || (*statsPort < 1024) {
		return errors.New("Error: Cannot use port numbers < 1024")
	}

	if *inputPort == *statsPort {
		return errors.New("Error: Input port and Stats port cannot use the same port number")
	}

	return nil
}
