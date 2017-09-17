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
About LOC, first field is longitude, second is latitude, third is zoom(now is dummy).

### example
```
// Suita City
LOC = location.Location{"135.5266571044922", "34.77274096606007", "12"}
```

'd'option shows a rain cloud score.

'l'option sets location. Default is LOC in var.go.

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
