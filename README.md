# GeoHashi

[![Build Status](https://travis-ci.org/bsm/geohashi.png?branch=master)](https://travis-ci.org/bsm/geohashi)
[![GoDoc](https://godoc.org/github.com/bsm/geohashi?status.png)](http://godoc.org/github.com/bsm/geohashi)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsm/geohashi)](https://goreportcard.com/report/github.com/bsm/geohashi)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Geohash library for Go, optimised for performance.

### Examples

Encode coordinates:

```go
func Encode() {
	lat, lon := 51.52463, -0.08411

	hash1, _ := geohashi.EncodeWithPrecision(lat, lon, 10)
	bin1, _ := hash1.MarshalBinary()
	fmt.Printf("%d/%d %q\n", hash1.Value(), hash1.Precision(), base64.RawURLEncoding.EncodeToString(bin1))

	hash2, _ := geohashi.EncodeWithPrecision(lat, lon, 20)
	bin2, _ := hash2.MarshalBinary()
	fmt.Printf("%d/%d %q\n", hash2.Value(), hash2.Precision(), base64.RawURLEncoding.EncodeToString(bin2))

}
```

Decode hash:

```go
func Decode() {
	hash, _ := geohashi.NewWithPrecision(528212444555, 20)
	area := hash.Decode()
	lat, lon := area.Center()

	fmt.Printf("%.5f,%.5f\n", lat, lon)
	fmt.Println(area.Contains(51.52463, -0.08411))

}
```
