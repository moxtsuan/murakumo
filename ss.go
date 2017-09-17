package main

import (
	"github.com/moxtsuan/murakumo/location"
	"github.com/sclevine/agouti"
)

func getImage(loc string) error {
	var l location.Location
	switch loc {
	case "tokyo":
		l = location.Tokyo
	case "osaka":
		l = location.Osaka
	case "takamatsu":
		l = location.Takamatsu
	default:
		l = location.Tokyo
	}
	url := BASE_URL + "lon=" + l.LON + "&lat=" + l.LAT + "&" + OPT
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

func ScrShot(loc string) error {
	err := getImage(loc)
	return err
}
