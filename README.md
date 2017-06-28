# GeoHashi

[![Build Status](https://travis-ci.org/bsm/geohashi.png?branch=master)](https://travis-ci.org/bsm/geohashi)
[![GoDoc](https://godoc.org/github.com/bsm/geohashi?status.png)](http://godoc.org/github.com/bsm/geohashi)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsm/geohashi)](https://goreportcard.com/report/github.com/bsm/geohashi)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Geohash library for Go, optimised for performance. Inspired by [geohash-int](https://github.com/yinqiwen/geohash-int).

## Examples

Encode coordinates:

```go
func Encode() {
	lat, lon := 51.52463, -0.08411

	hash1 := geohashi.EncodeWithPrecision(lat, lon, 10)
	fmt.Printf("%d (%d)\n", hash1, hash1.Precision())

	hash2 := geohashi.EncodeWithPrecision(lat, lon, 20)
	fmt.Printf("%d (%d)\n", hash2, hash2.Precision())
}
```

Decode hash:

```go
func Decode() {
	hash := geohashi.Hash(90072520759854475)
	area := hash.Decode()
	lat, lon := area.Center()

	fmt.Printf("%.5f,%.5f\n", lat, lon)
	fmt.Println(area.Contains(51.52463, -0.08411))
}
```

## Precision

This library allows you to select the precision of the hash you want to create, up to a maximum of 26 bits.

The following table shows the maximum uncertainty (at the equator) for a given bit-precision:

|Bits| Uncertainty |
|----|-------------|
| 26 | ±0.35m      |
| 25 | ±0.65m      |
| 24 | ±1.3m       |
| 23 | ±2.6m       |
| 22 | ±5.3m       |
| 21 | ±11m        |
| 20 | ±21m        |
| 19 | ±42m        |
| 18 | ±84m        |
| 17 | ±170m       |
| 16 | ±340m       |
| 15 | ±680m       |
| 14 | ±1400m      |
| 13 | ±2700m      |
| 12 | ±5400m      |
| 11 | ±11000m     |
| 10 | ±22000m     |
| 9  | ±43000m     |
| 8  | ±86000m     |
| 7  | ±170000m    |
| 6  | ±350000m    |
| 5  | ±690000m    |
| 4  | ±1400000m   |
| 3  | ±2800000m   |
| 2  | ±5400000m   |
| 1  | ±10000000m  |
