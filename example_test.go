package geohashi_test

import (
	"fmt"

	"github.com/bsm/geohashi"
)

func ExampleEncode() {
	lat, lon := 51.52463, -0.08411

	hash1 := geohashi.EncodeWithPrecision(lat, lon, 10)
	fmt.Printf("%d (%d)\n", hash1, hash1.Precision())

	hash2 := geohashi.EncodeWithPrecision(lat, lon, 20)
	fmt.Printf("%d (%d)\n", hash2, hash2.Precision())

	// Output:
	// 45035996274208702 (10)
	// 90072520759854475 (20)
}

func ExampleHash() {
	hash := geohashi.Hash(90072520759854475)
	area := hash.Decode()
	lat, lon := area.Center()

	fmt.Printf("%.5f,%.5f\n", lat, lon)
	fmt.Println(area.Contains(51.52463, -0.08411))

	// Output:
	// 51.52460,-0.08394
	// true
}
