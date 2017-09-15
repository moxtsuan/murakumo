package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	ccdisp = flag.Bool("d", false, "Display Color Count")
)

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
	if s > SCORE {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
