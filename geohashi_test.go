package geohashi

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hash", func() {
	const lat, lon = 51.524632318, -0.0841140747
	const value uint64 = 2163558172898388

	DescribeTable("should init",
		func(bits uint64, prec uint8) {
			hash, err := NewWithPrecision(bits, prec)
			Expect(err).NotTo(HaveOccurred())
			Expect(hash.Value()).To(Equal(bits))
			Expect(hash.Precision()).To(Equal(prec))
		},

		Entry("min-precision", 1, 1),
		Entry("max-precision", 2163558172898388, 26),
	)

	It("should encode with precision", func() {
		for _, test := range testCases {
			hash, err := EncodeWithPrecision(test.lat, test.lon, test.prec)
			Expect(err).NotTo(HaveOccurred())
			Expect(hash.Value()).To(Equal(test.bits), "for %+v", test)
			Expect(hash.Precision()).To(Equal(test.prec), "for %+v", test)
		}

		hash := Encode(lat, lon)
		Expect(hash.Value()).To(Equal(value))
		Expect(hash.Precision()).To(Equal(uint8(26)))

		_, err := EncodeWithPrecision(lat, lon, 0)
		Expect(err).To(Equal(ErrInvalidPrecision))

		_, err = EncodeWithPrecision(lat, lon, 27)
		Expect(err).To(Equal(ErrInvalidPrecision))
	})

	It("should decode", func() {
		for _, test := range testCases {
			hash, _ := NewWithPrecision(test.bits, test.prec)
			area := hash.Decode()
			Expect(area.MinLat).To(BeNumerically("<=", test.lat), "for %+v -> %s", test, area)
			Expect(area.MinLon).To(BeNumerically("<=", test.lon), "for %+v -> %s", test, area)
			Expect(area.MaxLat).To(BeNumerically(">=", test.lat), "for %+v -> %s", test, area)
			Expect(area.MaxLon).To(BeNumerically(">=", test.lon), "for %+v -> %s", test, area)
		}
	})

	It("should zoom out", func() {
		hash := Encode(51.5246323180, -0.0841140747)
		p1, err := hash.Parent()
		Expect(err).NotTo(HaveOccurred())
		Expect(p1.Precision()).To(Equal(uint8(25)))
		Expect(p1.Value()).To(Equal(uint64(540889543224597)))

		p2, err := p1.Parent()
		Expect(err).NotTo(HaveOccurred())
		Expect(p2.Precision()).To(Equal(uint8(24)))
		Expect(p2.Value()).To(Equal(uint64(135222385806149)))

		root, _ := NewWithPrecision(0, 1)
		_, err = root.Parent()
		Expect(err).To(Equal(ErrInvalidPrecision))
	})

	It("should zoom in", func() {
		hash, _ := NewWithPrecision(135222385806149, 24)
		subs := hash.Children()
		Expect(subs).To(HaveLen(4))

		Expect(subs[0].Precision()).To(Equal(uint8(25)))
		Expect(subs[1].Precision()).To(Equal(uint8(25)))
		Expect(subs[2].Precision()).To(Equal(uint8(25)))
		Expect(subs[3].Precision()).To(Equal(uint8(25)))

		Expect(subs[0].Value()).To(Equal(uint64(540889543224596)))
		Expect(subs[1].Value()).To(Equal(uint64(540889543224597)))
		Expect(subs[2].Value()).To(Equal(uint64(540889543224598)))
		Expect(subs[3].Value()).To(Equal(uint64(540889543224599)))
	})

	It("should move X", func() {
		hash, _ := NewWithPrecision(135222385806149, 24)
		east := hash.MoveX(1)
		west := hash.MoveX(-1)

		Expect(east.Precision()).To(Equal(uint8(24)))
		Expect(west.Precision()).To(Equal(uint8(24)))

		Expect(east.Value()).To(Equal(uint64(135222385806151)))
		Expect(west.Value()).To(Equal(uint64(135222385805807)))

		Expect(east.MoveX(1)).To(Equal(hash.MoveX(2)))
		Expect(east.MoveX(-1)).To(Equal(hash))

		Expect(west.MoveX(1)).To(Equal(hash))
		Expect(west.MoveX(-1)).To(Equal(hash.MoveX(-2)))
	})

	It("should move Y", func() {
		hash, _ := NewWithPrecision(135222385806149, 24)
		north := hash.MoveY(1)
		south := hash.MoveY(-1)

		Expect(north.Precision()).To(Equal(uint8(24)))
		Expect(south.Precision()).To(Equal(uint8(24)))

		Expect(north.Value()).To(Equal(uint64(135222385806160)))
		Expect(south.Value()).To(Equal(uint64(135222385806148)))

		Expect(north.MoveY(1)).To(Equal(hash.MoveY(2)))
		Expect(north.MoveY(-1)).To(Equal(hash))

		Expect(south.MoveY(1)).To(Equal(hash))
		Expect(south.MoveY(-1)).To(Equal(hash.MoveY(-2)))
	})

})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "github.com/bsm/geohashi")
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash := New(uint64(rand.Int63() >> 12))
		hash.Decode()
	}
}
