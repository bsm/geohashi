package geohashi

import (
	"encoding/binary"
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

// ErrInvalidPrecision is returned on bad precision inputs
var ErrInvalidPrecision = errors.New("geohashi: invalid precision")

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
type Hash struct {
	bits uint64
	prec uint8
}

// New creates a new Hash with default precision
func New(value uint64) Hash {
	hash, _ := NewWithPrecision(value, PrecisionMax)
	return hash
}

// NewWithPrecision creates a new Hash with custom precision
func NewWithPrecision(value uint64, prec uint8) (h Hash, err error) {
	if prec < PrecisionMin || prec > PrecisionMax {
		return h, ErrInvalidPrecision
	}
	h.bits = value
	h.prec = prec
	return
}

// Encode converts a lat/lon to an geohash with maximum precision
func Encode(lat, lon float64) Hash {
	hash, _ := EncodeWithPrecision(lat, lon, PrecisionMax)
	return hash
}

// EncodeWithPrecision converts a lat/lon to an numeric geohash
func EncodeWithPrecision(lat, lon float64, prec uint8) (h Hash, err error) {
	if prec < PrecisionMin || prec > PrecisionMax {
		return h, ErrInvalidPrecision
	}

	dx := (lat - LatMin) / latScale
	dy := (lon - LonMin) / lonScale
	gn := float64(uint64(1) << prec)

	h.bits = interleave64(uint64(dx*gn), uint64(dy*gn))
	h.prec = prec
	return
}

// Value returns the hash value
func (h Hash) Value() uint64 { return h.bits }

// Precision returns the prec level
func (h Hash) Precision() uint8 { return h.prec }

// Parent zooms out, returning the parent hash, lowering the precision. May return ErrInvalidPrecision
// if unable to zoom out any further
func (h Hash) Parent() (Hash, error) {
	parent := Hash{prec: h.prec - 1}
	if parent.prec < PrecisionMin {
		return parent, ErrInvalidPrecision
	}
	parent.bits = h.bits >> 2
	return parent, nil
}

// Children zooms in, returning four child hashes, in the followin order SW, SE, NW, NE. May return
// nil if unable to zoom in further
func (h Hash) Children() []Hash {
	base := Hash{prec: h.prec + 1, bits: (h.bits << 2)}
	if base.prec > PrecisionMax {
		return nil
	}

	return []Hash{
		base,
		{prec: base.prec, bits: base.bits | 1},
		{prec: base.prec, bits: base.bits | 2},
		{prec: base.prec, bits: base.bits | 3},
	}
}

// MoveX moves n steps east (positive number) or west (negative number) and
// returns the resulting hash
func (h Hash) MoveX(n int) Hash {
	if n == 0 {
		return h
	}

	shift := (64 - h.prec*2)
	east := n > 0
	if !east {
		n = -n
	}

	bits := h.bits
	for i := 0; i < n; i++ {
		x := bits & s7
		y := bits & s1
		zz := uint64(s1 >> shift)

		if east {
			x += zz + 1
		} else {
			x = (x | zz) - zz - 1
		}
		x &= s7 >> shift
		bits = x | y
	}
	return Hash{bits, h.prec}
}

// MoveY moves n steps north (positive number) or south (negative number) and
// returns the resulting hash
func (h Hash) MoveY(n int) Hash {
	if n == 0 {
		return h
	}

	shift := (64 - h.prec*2)
	east := n > 0
	if !east {
		n = -n
	}

	bits := h.bits
	for i := 0; i < n; i++ {
		x := bits & s7
		y := bits & s1
		zz := uint64(s7 >> shift)

		if east {
			y += zz + 1
		} else {
			y = (y | zz) - zz - 1
		}
		y &= s1 >> shift
		bits = x | y
	}
	return Hash{bits, h.prec}
}

// Decode decodes a hash into an area
func (h Hash) Decode() (area Area) {
	x, y := deinterleave64(h.bits)
	fx, fy := float64(x), float64(y)

	gn := float64(uint(1) << h.Precision())
	gx, gy := gn/latScale, gn/lonScale

	area.MinLat = LatMin + fx/gx
	area.MinLon = LonMin + fy/gy
	area.MaxLat = LatMin + (fx+1)/gx
	area.MaxLon = LonMin + (fy+1)/gy
	return
}

// MarshalBinary implements encoding.BinaryMarshaler
func (h Hash) MarshalBinary() ([]byte, error) {
	b := make([]byte, binary.MaxVarintLen64+1)
	b[0] = h.prec
	n := binary.PutVarint(b[1:], int64(h.bits))
	return b[:n+1], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (h *Hash) UnmarshalBinary(b []byte) error {
	if len(b) < 2 {
		return nil
	}

	n, _ := binary.Varint(b[1:])
	*h = Hash{
		prec: b[0],
		bits: uint64(n),
	}
	return nil
}
