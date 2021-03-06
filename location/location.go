package location

type Location struct {
	LON  string
	LAT  string
	ZOOM string
}

var (
	Tokyo     = Location{"139.68875885009768", "35.679051851352995", "12"}
	Sapporo   = Location{"141.33499145507812", "43.05082689000008", "12"}
	Sendai    = Location{"140.8598327636719", "38.25193153082154", "12"}
	Nigata    = Location{"139.01344299316406", "37.88271206292802", "12"}
	Nagoya    = Location{"136.90303802490234", "35.14770508596119", "12"}
	Osaka     = Location{"135.50159454345703", "34.671617436363604", "12"}
	Okayama   = Location{"133.93054962158203", "34.63518530664358", "12"}
	Takamatsu = Location{"134.04178619384766", "34.32217300445413", "12"}
	Hakata    = Location{"130.4128646850586", "33.569147246109395", "12"}
	Naha      = Location{"127.68310546875001", "26.192412147582772", "12"}
)
