package location

type Location struct {
	LON  string
	LAT  string
	zoom string
}

var (
	Tokyo     = Location{"139.68875885009768", "35.679051851352995", "12"}
	Hokkaido  = Location{"", "", ""}
	Osaka     = Location{"", "", ""}
	Okayama   = Location{"", "", ""}
	Takamatsu = Location{"134.04178619384766", "34.32217300445413", "12"}
)
