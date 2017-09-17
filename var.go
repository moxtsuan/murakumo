package main

import "github.com/moxtsuan/murakumo/location"

var (
	BASE_URL = "http://www.river.go.jp/x/krd0107010.php?"
	//LON       = "lon=134.04178619384766"
	//LAT       = "lat=34.32217300445413"
	LOC       = location.Location{"134.04178919384766", "34.32217300445413", "12"}
	OPT       = "opa=0&zoom=12&leg=0&ext=0"
	SS        = "/tmp/amagumo_murakumo.png"
	TRIM      = "/tmp/trim_murakumo.png"
	SCORE     = 20000
	APP_RANGE = 3
)
