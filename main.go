package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := ScrShot()
	if err != nil {
		log.Fatal(err)
	}
	err = Cutter()
	if err != nil {
		log.Fatal(err)
	}
	s, err := ColorCount()
	if err != nil {
		log.Fatal(err)
	}
	if s > SCORE {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
