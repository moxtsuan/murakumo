murakumo
====

Rain Cloud Detect in Japan

## Requirement

- \*nix
- PhantomJS

## Installation

```
$ go get github.com/moxtsuan/murakumo
```

## Usage

Set location in var.go you want to survey.

### example
```
// Suita City
LON = "lon=135.5266571044922"
LAT = "lat=34.77274096606007"
```

'd'option shows a rain cloud score.

'l'option sets location.

```
tokyo
sapporo
sendai
nigata
nagoya
osaka
okayama
takamatsu
hakata
naha
```

```
$ murakumo [-d] [-l] location
```

## License

MIT License
