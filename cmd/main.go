package main

import (
	"github.com/setlog/systemuri"
	"log"
)

func main() {
	if err := systemuri.RegisterURLHandler("OSCA DC TU", "josca-osca-dc-test", "test", "%s"); err != nil {
		log.Fatal(err)
	}
	if err := systemuri.UnregisterURLHandler("josca-osca-dc-test"); err != nil {
		log.Fatal(err)
	}
}
