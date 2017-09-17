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

## Library
[agouti](https://github.com/sclevine/agouti)

>
>The MIT License(MIT)<br>
>Copyright (c) 2014 Stephen Levine

- License: The MIT License(MIT)
- Full Text: [agouti/LICENSE](https://github.com/sclevine/agouti/blob/master/LICENSE)

[cutter](https://github.com/oliamb/cutter)

>
>The MIT License(MIT)<br>
>Copyright (c) 2014 Olivier Amblet

- License: The MIT License(MIT)
- Full Text: [cutter/LICENSE](https://github.com/oliamb/cutter/blob/master/LICENSE)
