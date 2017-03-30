package geohashi_test

import (
	"encoding/base64"
	"fmt"

	"github.com/bsm/geohashi"
)

func ExampleEncode() {
	lat, lon := 51.52463, -0.08411

	hash1, _ := geohashi.EncodeWithPrecision(lat, lon, 10)
	bin1, _ := hash1.MarshalBinary()
	fmt.Printf("%d/%d %q\n", hash1.Value(), hash1.Precision(), base64.RawURLEncoding.EncodeToString(bin1))

	hash2, _ := geohashi.EncodeWithPrecision(lat, lon, 20)
	bin2, _ := hash2.MarshalBinary()
	fmt.Printf("%d/%d %q\n", hash2.Value(), hash2.Precision(), base64.RawURLEncoding.EncodeToString(bin2))

	// Output:
	// 503742/10 "Cvy-PQ"
	// 528212444555/20 "FJaW0r7fHg"
}

func ExampleHash() {
	hash, _ := geohashi.NewWithPrecision(528212444555, 20)
	area := hash.Decode()
	lat, lon := area.Center()

	fmt.Printf("%.5f,%.5f\n", lat, lon)
	fmt.Println(area.Contains(51.52463, -0.08411))

	// Output:
	// 51.52460,-0.08394
	// true
}
