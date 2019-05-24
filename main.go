// +build !test
// Our first program

package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("======================================")

	log.WithFields(log.Fields{
		"fields": "values",
	}).Info("Log Info With Fields")
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}
