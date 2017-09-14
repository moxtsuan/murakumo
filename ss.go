package main

import (
	"github.com/sclevine/agouti"
)

func getImage() error {
	url := BASE_URL + LON + "&" + LAT + "&" + OPT
	driver := agouti.PhantomJS()
	err := driver.Start()
	if err != nil {
		return err
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		return err
	}

	page.Navigate(url)
	if err != nil {
		return err
	}

	page.Screenshot(SS)

	return nil
}

func ScrShot() error {
	err := getImage()
	return err
}
