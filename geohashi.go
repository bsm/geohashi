package geohashi

import (
	"errors"
	"math"
)

/* Limits from EPSG:900913 / EPSG:3785 / OSGEO:41001 */
const (
	LatMin = -85.05112878
	LatMax = -LatMin
	LonMin = -180.0
	LonMax = -LonMin

	PrecisionMin = 1
	PrecisionMax = 26

	latScale = LatMax - LatMin
	lonScale = LonMax - LonMin
)

var errInvalidPrecision = errors.New("geohashi: invalid precision")

// --------------------------------------------------------------------

// Area is a square area defined through a min and max lat/lon
type Area struct{ MinLat, MaxLat, MinLon, MaxLon float64 }

// Center returns the area's centeroid coordinates
func (a Area) Center() (lat, lon float64) {
	lat = (a.MinLat + a.MaxLat) / 2.0
	lon = (a.MinLon + a.MaxLon) / 2.0
	return
}

// Contains returns true if coordinates are contained within the area.
func (a Area) Contains(lat, lon float64) bool {
	return (a.MinLat <= lat && lat <= a.MaxLat &&
		a.MinLon <= lon && lon <= a.MaxLon)
}

func maxDecimalPower(r float64) float64 {
	m := int(math.Floor(math.Log10(r)))
	return math.Pow10(m)
}

// --------------------------------------------------------------------

// A Hash a numeric geohash value
type Hash uint64

func newHash(base uint64, prec uint8) Hash {
	return Hash(base) | Hash(prec)<<52
}

// Encode converts a lat/lon to an geohash with maximum precision
func Encode(lat, lon float64) Hash {
	return EncodeWithPrecision(lat, lon, PrecisionMax)
}

// EncodeWithPrecision converts a lat/lon to an numeric geohash
func EncodeWithPrecision(lat, lon float64, prec uint8) Hash {
	if prec < PrecisionMin || prec > PrecisionMax {
		return 0
	}

	dx := (lat - LatMin) / latScale
	dy := (lon - LonMin) / lonScale
	gn := float64(uint64(1) << prec)

	base := interleave64(uint64(dx*gn), uint64(dy*gn))
	return newHash(base, prec)
}

// Precision returns the prec level
func (h Hash) Precision() uint8 { return uint8(h >> 52) }

func (h Hash) base() uint64 { return uint64(h & s8) }

// Decode decodes a hash into an area
func (h Hash) Decode() (area Area) {
	x, y := deinterleave64(h.base())
	fx, fy := float64(x), float64(y)

	gn := float64(uint(1) << h.Precision())
	gx, gy := gn/latScale, gn/lonScale

	area.MinLat = LatMin + fx/gx
	area.MinLon = LonMin + fy/gy
	area.MaxLat = LatMin + (fx+1)/gx
	area.MaxLon = LonMin + (fy+1)/gy
	return
}

// Parent zooms out, returning the parent hash, lowering the precision. This function may
// return Hash(0) if unable to zoom out further
func (h Hash) Parent() Hash {
	prec := h.Precision()
	if prec <= PrecisionMin {
		return 0
	}
	return newHash(h.base()>>2, prec-1)
}

// Children zooms in, returning four child hashes, in the following order SW, SE, NW, NE.
// This function may return nil if unable to zoom in further
func (h Hash) Children() []Hash {
	prec := h.Precision()
	if prec >= PrecisionMax {
		return nil
	}

	child := newHash(h.base()<<2, prec+1)
	return []Hash{child, child | 1, child | 2, child | 3}
}

// MoveX moves n steps east (positive number) or west (negative number) and
// returns the resulting hash
func (h Hash) MoveX(n int) Hash {
	if n == 0 {
		return h
	}

	prec := h.Precision()
	shift := (64 - prec*2)
	east := n > 0
	if !east {
		n = -n
	}

	base := h.base()
	for i := 0; i < n; i++ {
		x := base & s7
		y := base & s1
		zz := uint64(s1 >> shift)

		if east {
			x += zz + 1
		} else {
			x = (x | zz) - zz - 1
		}
		x &= s7 >> shift
		base = x | y
	}
	return newHash(base, prec)
}

// MoveY moves n steps north (positive number) or south (negative number) and
// returns the resulting hash
func (h Hash) MoveY(n int) Hash {
	if n == 0 {
		return h
	}

	prec := h.Precision()
	shift := (64 - prec*2)
	east := n > 0
	if !east {
		n = -n
	}

	base := h.base()
	for i := 0; i < n; i++ {
		x := base & s7
		y := base & s1
		zz := uint64(s7 >> shift)

		if east {
			y += zz + 1
		} else {
			y = (y | zz) - zz - 1
		}
		y &= s1 >> shift
		base = x | y
	}
	return newHash(base, prec)
}
