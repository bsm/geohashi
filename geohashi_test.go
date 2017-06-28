package geohashi

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hash", func() {
	const lat, lon = 51.524632318, -0.0841140747

	It("should not encode bad precisions", func() {
		Expect(EncodeWithPrecision(lat, lon, 0)).To(Equal(Hash(0)))
		Expect(EncodeWithPrecision(lat, lon, 27)).To(Equal(Hash(0)))
	})

	DescribeTable("should encode",
		func(lat, lon float64, prec int, exp Hash) {
			hash := EncodeWithPrecision(lat, lon, uint8(prec))
			Expect(hash).To(Equal(exp))
			Expect(hash.Precision()).To(Equal(uint8(prec)))
		},

		Entry("precision: 01", lat, lon, 1, Hash(4503599627370497)),  // 0x0010000000000001
		Entry("precision: 02", lat, lon, 2, Hash(9007199254740999)),  // 0x0020000000000007
		Entry("precision: 03", lat, lon, 3, Hash(13510798882111518)), // 0x003000000000001e
		Entry("precision: 04", lat, lon, 4, Hash(18014398509482106)),
		Entry("precision: 05", lat, lon, 5, Hash(22517998136852971)),
		Entry("precision: 06", lat, lon, 6, Hash(27021597764224943)),
		Entry("precision: 07", lat, lon, 7, Hash(31525197391601342)),
		Entry("precision: 08", lat, lon, 8, Hash(36028797018995451)),
		Entry("precision: 09", lat, lon, 9, Hash(40532396646460399)),
		Entry("precision: 10", lat, lon, 10, Hash(45035996274208702)),
		Entry("precision: 11", lat, lon, 11, Hash(49539595903090426)),
		Entry("precision: 12", lat, lon, 12, Hash(54043195536505834)),
		Entry("precision: 13", lat, lon, 13, Hash(58546795188055977)),
		Entry("precision: 14", lat, lon, 14, Hash(63050394912145060)),
		Entry("precision: 15", lat, lon, 15, Hash(67553994926389905)),
		Entry("precision: 16", lat, lon, 16, Hash(72057596101257797)),
		Entry("precision: 17", lat, lon, 17, Hash(76561201918617878)),
		Entry("precision: 18", lat, lon, 18, Hash(81064826305946712)),
		Entry("precision: 19", lat, lon, 19, Hash(85568524973150562)),
		Entry("precision: 20", lat, lon, 20, Hash(90072520759854475)),
		Entry("precision: 21", lat, lon, 21, Hash(94577705024558637)),
		Entry("precision: 22", lat, lon, 22, Hash(99087643201263796)),
		Entry("precision: 23", lat, lon, 23, Hash(103616597025972945)),
		Entry("precision: 24", lat, lon, 24, Hash(108221613442698053)),
		Entry("precision: 25", lat, lon, 25, Hash(113130880227486997)),
		Entry("precision: 26", lat, lon, 26, Hash(119257148484531284)),

		Entry("SW1", -75.0, -120.0, 1, Hash(0x0010000000000000)),
		Entry("SW2", -75.0, -20.0, 1, Hash(0x0010000000000000)),
		Entry("SW3", -5.0, -20.0, 1, Hash(0x0010000000000000)),
		Entry("SW4", -5.0, -120.0, 1, Hash(0x0010000000000000)),

		Entry("NW1", 75.0, -120.0, 1, Hash(0x0010000000000001)),
		Entry("NW2", 75.0, -20.0, 1, Hash(0x0010000000000001)),
		Entry("NW3", 5.0, -20.0, 1, Hash(0x0010000000000001)),
		Entry("NW4", 5.0, -120.0, 1, Hash(0x0010000000000001)),

		Entry("SE1", -75.0, 120.0, 1, Hash(0x0010000000000002)),
		Entry("SE2", -75.0, 20.0, 1, Hash(0x0010000000000002)),
		Entry("SE3", -5.0, 20.0, 1, Hash(0x0010000000000002)),
		Entry("SE4", -5.0, 120.0, 1, Hash(0x0010000000000002)),

		Entry("NE1", 75.0, 120.0, 1, Hash(0x0010000000000003)),
		Entry("NE2", 75.0, 20.0, 1, Hash(0x0010000000000003)),
		Entry("NE3", 5.0, 20.0, 1, Hash(0x0010000000000003)),
		Entry("NE4", 5.0, 120.0, 1, Hash(0x0010000000000003)),

		Entry("generated test-case #001", -84.058614, -66.348116, 2, Hash(9007199254740994)),
		Entry("generated test-case #002", -31.506116, -91.139509, 2, Hash(9007199254740993)),
		Entry("generated test-case #003", -78.651825, 161.012861, 2, Hash(9007199254741002)),
		Entry("generated test-case #004", -12.911255, 96.956395, 2, Hash(9007199254741003)),
		Entry("generated test-case #005", 14.043176, -36.898120, 2, Hash(9007199254740998)),
		Entry("generated test-case #006", 8.633066, 178.335932, 3, Hash(13510798882111546)),
		Entry("generated test-case #007", 76.628258, -80.659264, 3, Hash(13510798882111517)),
		Entry("generated test-case #008", 67.720181, 147.289301, 3, Hash(13510798882111551)),
		Entry("generated test-case #009", 28.034480, -51.236922, 3, Hash(13510798882111513)),
		Entry("generated test-case #010", -13.729166, 45.342552, 3, Hash(13510798882111527)),
		Entry("generated test-case #011", 46.673677, 117.053852, 4, Hash(18014398509482226)),
		Entry("generated test-case #012", -62.686902, 30.188373, 4, Hash(18014398509482118)),
		Entry("generated test-case #013", 53.771926, -48.025206, 4, Hash(18014398509482099)),
		Entry("generated test-case #014", -46.222334, 118.960375, 4, Hash(18014398509482151)),
		Entry("generated test-case #015", -62.602523, 107.772026, 4, Hash(18014398509482148)),
		Entry("generated test-case #016", 12.761788, -66.661137, 5, Hash(22517998136852876)),
		Entry("generated test-case #017", 68.940262, 61.255672, 5, Hash(22517998136853362)),
		Entry("generated test-case #018", -40.284294, -119.711563, 5, Hash(22517998136852578)),
		Entry("generated test-case #019", 59.784886, 26.804451, 5, Hash(22517998136853325)),
		Entry("generated test-case #020", 4.761185, -36.171938, 5, Hash(22517998136852896)),
		Entry("generated test-case #021", -2.577317, -164.089397, 6, Hash(27021597764223325)),
		Entry("generated test-case #022", -60.319654, -35.949059, 6, Hash(27021597764223683)),
		Entry("generated test-case #023", 42.767691, -62.599473, 6, Hash(27021597764224800)),
		Entry("generated test-case #024", 78.856594, -129.197829, 6, Hash(27021597764224467)),
		Entry("generated test-case #025", 43.679957, -174.872572, 6, Hash(27021597764224256)),
		Entry("generated test-case #026", -68.855276, -169.646366, 7, Hash(31525197391593562)),
		Entry("generated test-case #027", 59.697633, 90.974174, 7, Hash(31525197391608912)),
		Entry("generated test-case #028", 65.542007, 3.469559, 7, Hash(31525197391607043)),
		Entry("generated test-case #029", -29.534680, -158.163975, 7, Hash(31525197391594603)),
		Entry("generated test-case #030", -60.053714, -122.137203, 7, Hash(31525197391594276)),
		Entry("generated test-case #031", -62.875596, 52.708074, 8, Hash(36028797018999843)),
		Entry("generated test-case #032", 16.201661, -151.418982, 8, Hash(36028797018981216)),
		Entry("generated test-case #033", 81.480099, -141.566281, 8, Hash(36028797018986446)),
		Entry("generated test-case #034", 36.385808, -166.218663, 8, Hash(36028797018981782)),
		Entry("generated test-case #035", 70.238171, -76.336655, 8, Hash(36028797018993859)),
		Entry("generated test-case #036", -84.874249, 80.177206, 9, Hash(40532396646476296)),
		Entry("generated test-case #037", -75.426052, 144.806097, 9, Hash(40532396646506994)),
		Entry("generated test-case #038", 11.744500, 67.341639, 9, Hash(40532396646540975)),
		Entry("generated test-case #039", 9.722976, 0.601094, 9, Hash(40532396646531409)),
		Entry("generated test-case #040", -46.856189, 173.402933, 9, Hash(40532396646514476)),
		Entry("generated test-case #041", -36.755125, 46.804887, 10, Hash(45035996274328614)),
		Entry("generated test-case #042", -48.727081, -27.057492, 10, Hash(45035996273892174)),
		Entry("generated test-case #043", -37.220939, 69.331278, 10, Hash(45035996274336119)),
		Entry("generated test-case #044", -56.071581, 13.244381, 10, Hash(45035996274248822)),
		Entry("generated test-case #045", 76.167304, -56.125295, 10, Hash(45035996274194500)),
		Entry("generated test-case #046", -73.892821, 65.808060, 11, Hash(49539595903330876)),
		Entry("generated test-case #047", -15.900868, -105.174808, 11, Hash(49539595901573250)),
		Entry("generated test-case #048", 59.389078, -44.420769, 11, Hash(49539595903062095)),
		Entry("generated test-case #049", -41.741089, -164.225114, 11, Hash(49539595901346499)),
		Entry("generated test-case #050", 76.876920, -168.986566, 11, Hash(49539595902471161)),
		Entry("generated test-case #051", -55.561769, 152.002671, 12, Hash(54043195539779606)),
		Entry("generated test-case #052", 28.706993, 71.528531, 12, Hash(54043195541966247)),
		Entry("generated test-case #053", -68.191232, 160.451493, 12, Hash(54043195539671315)),
		Entry("generated test-case #054", -16.129829, -148.468256, 12, Hash(54043195529903469)),
		Entry("generated test-case #055", -61.431061, -90.469666, 12, Hash(54043195529408456)),
		Entry("generated test-case #056", -27.548971, -168.186186, 13, Hash(58546795160424865)),
		Entry("generated test-case #057", -76.148427, -126.276052, 13, Hash(58546795158037624)),
		Entry("generated test-case #058", 9.756322, -60.853014, 13, Hash(58546795181626171)),
		Entry("generated test-case #059", -81.944774, -13.798907, 13, Hash(58546795166884147)),
		Entry("generated test-case #060", -32.608633, 46.738882, 13, Hash(58546795195750779)),
		Entry("generated test-case #061", 47.592151, 36.014898, 14, Hash(63050395004009578)),
		Entry("generated test-case #062", 47.546942, 131.691283, 14, Hash(63050395037695111)),
		Entry("generated test-case #063", -61.643558, -153.063979, 14, Hash(63050394789540054)),
		Entry("generated test-case #064", 74.786866, -110.868294, 14, Hash(63050394882811021)),
		Entry("generated test-case #065", 10.227259, 81.663786, 14, Hash(63050394995904865)),
		Entry("generated test-case #066", 27.661492, 7.181843, 15, Hash(67553995234267554)),
		Entry("generated test-case #067", 58.666457, -163.631515, 15, Hash(67553994753615379)),
		Entry("generated test-case #068", 33.710374, 2.991577, 15, Hash(67553995237036881)),
		Entry("generated test-case #069", 11.819978, 50.096897, 15, Hash(67553995253806266)),
		Entry("generated test-case #070", -27.960327, 2.905270, 15, Hash(67553995019146641)),
		Entry("generated test-case #071", -5.372769, 84.490702, 16, Hash(72057596717389414)),
		Entry("generated test-case #072", -0.280916, 120.995156, 16, Hash(72057597116367245)),
		Entry("generated test-case #073", -46.785350, -168.008241, 16, Hash(72057594134511934)),
		Entry("generated test-case #074", 15.142142, 93.786228, 16, Hash(72057597814692419)),
		Entry("generated test-case #075", 3.192852, -154.950138, 16, Hash(72057595146465876)),
		Entry("generated test-case #076", 75.525111, -59.214997, 17, Hash(76561201660670341)),
		Entry("generated test-case #077", -83.935240, 73.072053, 17, Hash(76561202929449445)),
		Entry("generated test-case #078", 59.992638, 16.937555, 17, Hash(76561207750837037)),
		Entry("generated test-case #079", 37.524426, 17.295856, 17, Hash(76561206944520070)),
		Entry("generated test-case #080", 23.585613, 167.751686, 17, Hash(76561209649583985)),
		Entry("generated test-case #081", -76.040358, -168.752895, 18, Hash(81064793422606326)),
		Entry("generated test-case #082", -20.066471, 165.142451, 18, Hash(81064844332601255)),
		Entry("generated test-case #083", 22.615591, 73.643903, 18, Hash(81064848628264594)),
		Entry("generated test-case #084", 13.134028, -143.048079, 18, Hash(81064811426153354)),
		Entry("generated test-case #085", 71.257153, -132.309947, 18, Hash(81064818063891772)),
		Entry("generated test-case #086", 24.857191, -71.599264, 19, Hash(85568501045369020)),
		Entry("generated test-case #087", 77.180755, -112.907372, 19, Hash(85568493560499705)),
		Entry("generated test-case #088", 24.395893, 90.395405, 19, Hash(85568637801943018)),
		Entry("generated test-case #089", 74.084387, -70.461965, 19, Hash(85568518513188156)),
		Entry("generated test-case #090", -58.302014, 89.793591, 19, Hash(85568546375770715)),
		Entry("generated test-case #091", 62.104249, 157.223097, 20, Hash(90073066217256838)),
		Entry("generated test-case #092", -1.905053, -139.314638, 20, Hash(90072095386283136)),
		Entry("generated test-case #093", 55.295150, -51.912415, 20, Hash(90072488868263896)),
		Entry("generated test-case #094", 42.622917, -17.180423, 20, Hash(90072516711977320)),
		Entry("generated test-case #095", 0.383442, -23.791470, 20, Hash(90072442080536045)),
		Entry("generated test-case #096", -51.912385, -103.140564, 21, Hash(94575852681297800)),
		Entry("generated test-case #097", -75.028013, 158.021144, 21, Hash(94578518466160538)),
		Entry("generated test-case #098", 9.518240, 27.587578, 21, Hash(94578931416347698)),
		Entry("generated test-case #099", -80.818601, 75.168620, 21, Hash(94577966636192017)),
		Entry("generated test-case #100", 22.055852, -147.045096, 21, Hash(94576797663276417)),
		Entry("generated test-case #101", 47.858674, -63.086957, 22, Hash(99087045695235707)),
		Entry("generated test-case #102", -73.644359, 75.039429, 22, Hash(99088753245441108)),
		Entry("generated test-case #103", -58.477977, 10.751534, 22, Hash(99088279916722000)),
		Entry("generated test-case #104", -37.048813, 95.606410, 22, Hash(99091306478782511)),
		Entry("generated test-case #105", -32.194079, -165.350827, 22, Hash(99080350747954011)),
		Entry("generated test-case #106", -37.083646, -80.089829, 23, Hash(103596099409285514)),
		Entry("generated test-case #107", 6.872270, 61.011301, 23, Hash(103637988441291084)),
		Entry("generated test-case #108", -14.404467, 54.683772, 23, Hash(103625789123276356)),
		Entry("generated test-case #109", -3.384762, -149.265198, 23, Hash(103589224177456608)),
		Entry("generated test-case #110", 25.019701, 179.515032, 23, Hash(103648414017775562)),
		Entry("generated test-case #111", -35.540316, 33.107316, 24, Hash(108247395431548240)),
		Entry("generated test-case #112", 23.890821, -160.568329, 24, Hash(108161879252297895)),
		Entry("generated test-case #113", 48.644812, -16.590939, 24, Hash(108220948340788028)),
		Entry("generated test-case #114", -15.464398, -79.657436, 24, Hash(108144022669358525)),
		Entry("generated test-case #115", 33.577681, -43.408546, 24, Hash(108206264806460282)),
		Entry("generated test-case #116", 60.813820, -178.563645, 25, Hash(112947456530841447)),
		Entry("generated test-case #117", 31.020849, 107.169461, 25, Hash(113596942442760823)),
		Entry("generated test-case #118", 13.992137, -77.457127, 25, Hash(113019103609474001)),
		Entry("generated test-case #119", -84.562653, -93.056676, 25, Hash(112636769814058642)),
		Entry("generated test-case #120", -34.334365, -84.240379, 25, Hash(112803022372291870)),
		Entry("generated test-case #121", 29.108849, 48.581989, 26, Hash(120687743755182671)),
		Entry("generated test-case #122", -46.813856, -74.018881, 26, Hash(117758477798413200)),
		Entry("generated test-case #123", 37.109642, -178.701971, 26, Hash(118308961334282828)),
		Entry("generated test-case #124", -31.818757, -36.275119, 26, Hash(118099096766653771)),
		Entry("generated test-case #125", 7.032831, -124.282703, 26, Hash(118367839390536337)),
	)

	DescribeTable("should decode areas",
		func(lat, lon float64, _ int, hash Hash) {
			area := hash.Decode()
			Expect(area.MinLat).To(BeNumerically("<=", lat))
			Expect(area.MinLon).To(BeNumerically("<=", lon))
			Expect(area.MaxLat).To(BeNumerically(">=", lat))
			Expect(area.MaxLon).To(BeNumerically(">=", lon))
		},

		Entry("precision: 01", lat, lon, 1, Hash(4503599627370497)),
		Entry("precision: 02", lat, lon, 2, Hash(9007199254740999)),
		Entry("precision: 03", lat, lon, 3, Hash(13510798882111518)),
		Entry("precision: 04", lat, lon, 4, Hash(18014398509482106)),
		Entry("precision: 05", lat, lon, 5, Hash(22517998136852971)),
		Entry("precision: 06", lat, lon, 6, Hash(27021597764224943)),
		Entry("precision: 07", lat, lon, 7, Hash(31525197391601342)),
		Entry("precision: 08", lat, lon, 8, Hash(36028797018995451)),
		Entry("precision: 09", lat, lon, 9, Hash(40532396646460399)),
		Entry("precision: 10", lat, lon, 10, Hash(45035996274208702)),
		Entry("precision: 11", lat, lon, 11, Hash(49539595903090426)),
		Entry("precision: 12", lat, lon, 12, Hash(54043195536505834)),
		Entry("precision: 13", lat, lon, 13, Hash(58546795188055977)),
		Entry("precision: 14", lat, lon, 14, Hash(63050394912145060)),
		Entry("precision: 15", lat, lon, 15, Hash(67553994926389905)),
		Entry("precision: 16", lat, lon, 16, Hash(72057596101257797)),
		Entry("precision: 17", lat, lon, 17, Hash(76561201918617878)),
		Entry("precision: 18", lat, lon, 18, Hash(81064826305946712)),
		Entry("precision: 19", lat, lon, 19, Hash(85568524973150562)),
		Entry("precision: 20", lat, lon, 20, Hash(90072520759854475)),
		Entry("precision: 21", lat, lon, 21, Hash(94577705024558637)),
		Entry("precision: 22", lat, lon, 22, Hash(99087643201263796)),
		Entry("precision: 23", lat, lon, 23, Hash(103616597025972945)),
		Entry("precision: 24", lat, lon, 24, Hash(108221613442698053)),
		Entry("precision: 25", lat, lon, 25, Hash(113130880227486997)),
		Entry("precision: 26", lat, lon, 26, Hash(119257148484531284)),

		Entry("SW1", -75.0, -120.0, 1, Hash(0x0010000000000000)),
		Entry("SW2", -75.0, -20.0, 1, Hash(0x0010000000000000)),
		Entry("SW3", -5.0, -20.0, 1, Hash(0x0010000000000000)),
		Entry("SW4", -5.0, -120.0, 1, Hash(0x0010000000000000)),

		Entry("NW1", 75.0, -120.0, 1, Hash(0x0010000000000001)),
		Entry("NW2", 75.0, -20.0, 1, Hash(0x0010000000000001)),
		Entry("NW3", 5.0, -20.0, 1, Hash(0x0010000000000001)),
		Entry("NW4", 5.0, -120.0, 1, Hash(0x0010000000000001)),

		Entry("SE1", -75.0, 120.0, 1, Hash(0x0010000000000002)),
		Entry("SE2", -75.0, 20.0, 1, Hash(0x0010000000000002)),
		Entry("SE3", -5.0, 20.0, 1, Hash(0x0010000000000002)),
		Entry("SE4", -5.0, 120.0, 1, Hash(0x0010000000000002)),

		Entry("NE1", 75.0, 120.0, 1, Hash(0x0010000000000003)),
		Entry("NE2", 75.0, 20.0, 1, Hash(0x0010000000000003)),
		Entry("NE3", 5.0, 20.0, 1, Hash(0x0010000000000003)),
		Entry("NE4", 5.0, 120.0, 1, Hash(0x0010000000000003)),

		Entry("generated test-case #001", -84.058614, -66.348116, 2, Hash(9007199254740994)),
		Entry("generated test-case #002", -31.506116, -91.139509, 2, Hash(9007199254740993)),
		Entry("generated test-case #003", -78.651825, 161.012861, 2, Hash(9007199254741002)),
		Entry("generated test-case #004", -12.911255, 96.956395, 2, Hash(9007199254741003)),
		Entry("generated test-case #005", 14.043176, -36.898120, 2, Hash(9007199254740998)),
		Entry("generated test-case #006", 8.633066, 178.335932, 3, Hash(13510798882111546)),
		Entry("generated test-case #007", 76.628258, -80.659264, 3, Hash(13510798882111517)),
		Entry("generated test-case #008", 67.720181, 147.289301, 3, Hash(13510798882111551)),
		Entry("generated test-case #009", 28.034480, -51.236922, 3, Hash(13510798882111513)),
		Entry("generated test-case #010", -13.729166, 45.342552, 3, Hash(13510798882111527)),
		Entry("generated test-case #011", 46.673677, 117.053852, 4, Hash(18014398509482226)),
		Entry("generated test-case #012", -62.686902, 30.188373, 4, Hash(18014398509482118)),
		Entry("generated test-case #013", 53.771926, -48.025206, 4, Hash(18014398509482099)),
		Entry("generated test-case #014", -46.222334, 118.960375, 4, Hash(18014398509482151)),
		Entry("generated test-case #015", -62.602523, 107.772026, 4, Hash(18014398509482148)),
		Entry("generated test-case #016", 12.761788, -66.661137, 5, Hash(22517998136852876)),
		Entry("generated test-case #017", 68.940262, 61.255672, 5, Hash(22517998136853362)),
		Entry("generated test-case #018", -40.284294, -119.711563, 5, Hash(22517998136852578)),
		Entry("generated test-case #019", 59.784886, 26.804451, 5, Hash(22517998136853325)),
		Entry("generated test-case #020", 4.761185, -36.171938, 5, Hash(22517998136852896)),
		Entry("generated test-case #021", -2.577317, -164.089397, 6, Hash(27021597764223325)),
		Entry("generated test-case #022", -60.319654, -35.949059, 6, Hash(27021597764223683)),
		Entry("generated test-case #023", 42.767691, -62.599473, 6, Hash(27021597764224800)),
		Entry("generated test-case #024", 78.856594, -129.197829, 6, Hash(27021597764224467)),
		Entry("generated test-case #025", 43.679957, -174.872572, 6, Hash(27021597764224256)),
		Entry("generated test-case #026", -68.855276, -169.646366, 7, Hash(31525197391593562)),
		Entry("generated test-case #027", 59.697633, 90.974174, 7, Hash(31525197391608912)),
		Entry("generated test-case #028", 65.542007, 3.469559, 7, Hash(31525197391607043)),
		Entry("generated test-case #029", -29.534680, -158.163975, 7, Hash(31525197391594603)),
		Entry("generated test-case #030", -60.053714, -122.137203, 7, Hash(31525197391594276)),
		Entry("generated test-case #031", -62.875596, 52.708074, 8, Hash(36028797018999843)),
		Entry("generated test-case #032", 16.201661, -151.418982, 8, Hash(36028797018981216)),
		Entry("generated test-case #033", 81.480099, -141.566281, 8, Hash(36028797018986446)),
		Entry("generated test-case #034", 36.385808, -166.218663, 8, Hash(36028797018981782)),
		Entry("generated test-case #035", 70.238171, -76.336655, 8, Hash(36028797018993859)),
		Entry("generated test-case #036", -84.874249, 80.177206, 9, Hash(40532396646476296)),
		Entry("generated test-case #037", -75.426052, 144.806097, 9, Hash(40532396646506994)),
		Entry("generated test-case #038", 11.744500, 67.341639, 9, Hash(40532396646540975)),
		Entry("generated test-case #039", 9.722976, 0.601094, 9, Hash(40532396646531409)),
		Entry("generated test-case #040", -46.856189, 173.402933, 9, Hash(40532396646514476)),
		Entry("generated test-case #041", -36.755125, 46.804887, 10, Hash(45035996274328614)),
		Entry("generated test-case #042", -48.727081, -27.057492, 10, Hash(45035996273892174)),
		Entry("generated test-case #043", -37.220939, 69.331278, 10, Hash(45035996274336119)),
		Entry("generated test-case #044", -56.071581, 13.244381, 10, Hash(45035996274248822)),
		Entry("generated test-case #045", 76.167304, -56.125295, 10, Hash(45035996274194500)),
		Entry("generated test-case #046", -73.892821, 65.808060, 11, Hash(49539595903330876)),
		Entry("generated test-case #047", -15.900868, -105.174808, 11, Hash(49539595901573250)),
		Entry("generated test-case #048", 59.389078, -44.420769, 11, Hash(49539595903062095)),
		Entry("generated test-case #049", -41.741089, -164.225114, 11, Hash(49539595901346499)),
		Entry("generated test-case #050", 76.876920, -168.986566, 11, Hash(49539595902471161)),
		Entry("generated test-case #051", -55.561769, 152.002671, 12, Hash(54043195539779606)),
		Entry("generated test-case #052", 28.706993, 71.528531, 12, Hash(54043195541966247)),
		Entry("generated test-case #053", -68.191232, 160.451493, 12, Hash(54043195539671315)),
		Entry("generated test-case #054", -16.129829, -148.468256, 12, Hash(54043195529903469)),
		Entry("generated test-case #055", -61.431061, -90.469666, 12, Hash(54043195529408456)),
		Entry("generated test-case #056", -27.548971, -168.186186, 13, Hash(58546795160424865)),
		Entry("generated test-case #057", -76.148427, -126.276052, 13, Hash(58546795158037624)),
		Entry("generated test-case #058", 9.756322, -60.853014, 13, Hash(58546795181626171)),
		Entry("generated test-case #059", -81.944774, -13.798907, 13, Hash(58546795166884147)),
		Entry("generated test-case #060", -32.608633, 46.738882, 13, Hash(58546795195750779)),
		Entry("generated test-case #061", 47.592151, 36.014898, 14, Hash(63050395004009578)),
		Entry("generated test-case #062", 47.546942, 131.691283, 14, Hash(63050395037695111)),
		Entry("generated test-case #063", -61.643558, -153.063979, 14, Hash(63050394789540054)),
		Entry("generated test-case #064", 74.786866, -110.868294, 14, Hash(63050394882811021)),
		Entry("generated test-case #065", 10.227259, 81.663786, 14, Hash(63050394995904865)),
		Entry("generated test-case #066", 27.661492, 7.181843, 15, Hash(67553995234267554)),
		Entry("generated test-case #067", 58.666457, -163.631515, 15, Hash(67553994753615379)),
		Entry("generated test-case #068", 33.710374, 2.991577, 15, Hash(67553995237036881)),
		Entry("generated test-case #069", 11.819978, 50.096897, 15, Hash(67553995253806266)),
		Entry("generated test-case #070", -27.960327, 2.905270, 15, Hash(67553995019146641)),
		Entry("generated test-case #071", -5.372769, 84.490702, 16, Hash(72057596717389414)),
		Entry("generated test-case #072", -0.280916, 120.995156, 16, Hash(72057597116367245)),
		Entry("generated test-case #073", -46.785350, -168.008241, 16, Hash(72057594134511934)),
		Entry("generated test-case #074", 15.142142, 93.786228, 16, Hash(72057597814692419)),
		Entry("generated test-case #075", 3.192852, -154.950138, 16, Hash(72057595146465876)),
		Entry("generated test-case #076", 75.525111, -59.214997, 17, Hash(76561201660670341)),
		Entry("generated test-case #077", -83.935240, 73.072053, 17, Hash(76561202929449445)),
		Entry("generated test-case #078", 59.992638, 16.937555, 17, Hash(76561207750837037)),
		Entry("generated test-case #079", 37.524426, 17.295856, 17, Hash(76561206944520070)),
		Entry("generated test-case #080", 23.585613, 167.751686, 17, Hash(76561209649583985)),
		Entry("generated test-case #081", -76.040358, -168.752895, 18, Hash(81064793422606326)),
		Entry("generated test-case #082", -20.066471, 165.142451, 18, Hash(81064844332601255)),
		Entry("generated test-case #083", 22.615591, 73.643903, 18, Hash(81064848628264594)),
		Entry("generated test-case #084", 13.134028, -143.048079, 18, Hash(81064811426153354)),
		Entry("generated test-case #085", 71.257153, -132.309947, 18, Hash(81064818063891772)),
		Entry("generated test-case #086", 24.857191, -71.599264, 19, Hash(85568501045369020)),
		Entry("generated test-case #087", 77.180755, -112.907372, 19, Hash(85568493560499705)),
		Entry("generated test-case #088", 24.395893, 90.395405, 19, Hash(85568637801943018)),
		Entry("generated test-case #089", 74.084387, -70.461965, 19, Hash(85568518513188156)),
		Entry("generated test-case #090", -58.302014, 89.793591, 19, Hash(85568546375770715)),
		Entry("generated test-case #091", 62.104249, 157.223097, 20, Hash(90073066217256838)),
		Entry("generated test-case #092", -1.905053, -139.314638, 20, Hash(90072095386283136)),
		Entry("generated test-case #093", 55.295150, -51.912415, 20, Hash(90072488868263896)),
		Entry("generated test-case #094", 42.622917, -17.180423, 20, Hash(90072516711977320)),
		Entry("generated test-case #095", 0.383442, -23.791470, 20, Hash(90072442080536045)),
		Entry("generated test-case #096", -51.912385, -103.140564, 21, Hash(94575852681297800)),
		Entry("generated test-case #097", -75.028013, 158.021144, 21, Hash(94578518466160538)),
		Entry("generated test-case #098", 9.518240, 27.587578, 21, Hash(94578931416347698)),
		Entry("generated test-case #099", -80.818601, 75.168620, 21, Hash(94577966636192017)),
		Entry("generated test-case #100", 22.055852, -147.045096, 21, Hash(94576797663276417)),
		Entry("generated test-case #101", 47.858674, -63.086957, 22, Hash(99087045695235707)),
		Entry("generated test-case #102", -73.644359, 75.039429, 22, Hash(99088753245441108)),
		Entry("generated test-case #103", -58.477977, 10.751534, 22, Hash(99088279916722000)),
		Entry("generated test-case #104", -37.048813, 95.606410, 22, Hash(99091306478782511)),
		Entry("generated test-case #105", -32.194079, -165.350827, 22, Hash(99080350747954011)),
		Entry("generated test-case #106", -37.083646, -80.089829, 23, Hash(103596099409285514)),
		Entry("generated test-case #107", 6.872270, 61.011301, 23, Hash(103637988441291084)),
		Entry("generated test-case #108", -14.404467, 54.683772, 23, Hash(103625789123276356)),
		Entry("generated test-case #109", -3.384762, -149.265198, 23, Hash(103589224177456608)),
		Entry("generated test-case #110", 25.019701, 179.515032, 23, Hash(103648414017775562)),
		Entry("generated test-case #111", -35.540316, 33.107316, 24, Hash(108247395431548240)),
		Entry("generated test-case #112", 23.890821, -160.568329, 24, Hash(108161879252297895)),
		Entry("generated test-case #113", 48.644812, -16.590939, 24, Hash(108220948340788028)),
		Entry("generated test-case #114", -15.464398, -79.657436, 24, Hash(108144022669358525)),
		Entry("generated test-case #115", 33.577681, -43.408546, 24, Hash(108206264806460282)),
		Entry("generated test-case #116", 60.813820, -178.563645, 25, Hash(112947456530841447)),
		Entry("generated test-case #117", 31.020849, 107.169461, 25, Hash(113596942442760823)),
		Entry("generated test-case #118", 13.992137, -77.457127, 25, Hash(113019103609474001)),
		Entry("generated test-case #119", -84.562653, -93.056676, 25, Hash(112636769814058642)),
		Entry("generated test-case #120", -34.334365, -84.240379, 25, Hash(112803022372291870)),
		Entry("generated test-case #121", 29.108849, 48.581989, 26, Hash(120687743755182671)),
		Entry("generated test-case #122", -46.813856, -74.018881, 26, Hash(117758477798413200)),
		Entry("generated test-case #123", 37.109642, -178.701971, 26, Hash(118308961334282828)),
		Entry("generated test-case #124", -31.818757, -36.275119, 26, Hash(118099096766653771)),
		Entry("generated test-case #125", 7.032831, -124.282703, 26, Hash(118367839390536337)),
	)

	It("should zoom out", func() {
		p26 := Encode(lat, lon)
		p25 := p26.Parent()
		Expect(p25).To(Equal(Hash(113130880227486997)))
		Expect(p25.Precision()).To(Equal(uint8(25)))

		p24 := p25.Parent()
		Expect(p24).To(Equal(Hash(108221613442698053)))
		Expect(p24.Precision()).To(Equal(uint8(24)))

		p1 := EncodeWithPrecision(lat, lon, 1)
		Expect(p1).To(Equal(Hash(4503599627370497)))
		Expect(p1.Parent()).To(Equal(Hash(0)))
		Expect(p1.Parent().Parent()).To(Equal(Hash(0)))
	})

	It("should zoom in", func() {
		Expect(Hash(0).Children()).To(Equal([]Hash{
			4503599627370496,
			4503599627370497,
			4503599627370498,
			4503599627370499,
		}))

		Expect(Hash(108221613442698053).Children()).To(Equal([]Hash{
			113130880227486996,
			113130880227486997,
			113130880227486998,
			113130880227486999,
		}))
	})

	It("should move X", func() {
		hash := Hash(108221613442698053)
		east := hash.MoveX(1)
		west := hash.MoveX(-1)

		Expect(east.Precision()).To(Equal(uint8(24)))
		Expect(west.Precision()).To(Equal(uint8(24)))

		Expect(west).To(Equal(Hash(108221613442697711)))

		Expect(east.MoveX(1)).To(Equal(hash.MoveX(2)))
		Expect(east.MoveX(-1)).To(Equal(hash))

		Expect(west.MoveX(1)).To(Equal(hash))
		Expect(west.MoveX(-1)).To(Equal(hash.MoveX(-2)))
	})

	It("should move Y", func() {
		hash := Hash(108221613442698053)
		north := hash.MoveY(1)
		south := hash.MoveY(-1)

		Expect(north.Precision()).To(Equal(uint8(24)))
		Expect(south.Precision()).To(Equal(uint8(24)))

		Expect(north).To(Equal(Hash(108221613442698064)))
		Expect(south).To(Equal(Hash(108221613442698052)))

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
	hash := Hash(119257148484531284)
	for i := 0; i < b.N; i++ {
		hash.Decode()
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(51.524632318, -0.0841140747)
	}
}
