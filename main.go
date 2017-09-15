package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	ccdisp = flag.Bool("d", false, "Display Color Count")

	exitCode = 0
)

func scoreCheck(s int) {
	if s > SCORE {
		exitCode = 2
	}
	return
}

func main() {
	flag.Parse()

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
	if *ccdisp {
		fmt.Println(s)
	}
	scoreCheck(s)
	os.Exit(exitCode)
}
