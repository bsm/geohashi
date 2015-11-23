package geohashi

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bitops", func() {

	tests := []struct {
		x, y uint64
		i    uint64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 2},
		{1, 1, 3},
		{2, 2, 12},
		{345, 456, 242113},
		{28116097, 17974564, 921773536331809},
	}

	It("should interleave", func() {
		for _, test := range tests {
			Expect(interleave64(test.x, test.y)).To(Equal(test.i), "for %v", test)
		}
	})

	It("should deinterleave", func() {
		for _, test := range tests {
			x, y := deinterleave64(test.i)
			Expect(x).To(Equal(test.x), "for %v", test)
			Expect(y).To(Equal(test.y), "for %v", test)
		}
	})

})
