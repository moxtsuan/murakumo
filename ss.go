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
	case "sapporo":
		l = location.Sapporo
	case "sendai":
		l = location.Sendai
	case "nigata":
		l = location.Nigata
	case "nagoya":
		l = location.Nagoya
	case "osaka":
		l = location.Osaka
	case "okayama":
		l = location.Okayama
	case "takamatsu":
		l = location.Takamatsu
	case "fukuoka":
		l = location.Hakata
	case "naha":
		l = location.Naha
	default:
		l = LOC
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
