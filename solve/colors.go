package solve

const (
	IdxBlack = iota
	IdxBlue
	IdxOrange
	IdxRed
)

const MaskCount byte = 0b00000011
const BitsPerColor = 2

// TileStack represents a collection of tiles sharing a common number but potentially different colors
// The collection may contain zero, one or two of each color.
//
// The count of each color is represented with 2-bits
type TileStack byte

func CountOfColor(tilestack TileStack, colorIndex int) int {
	return int(tilestack>>(BitsPerColor*colorIndex)) & int(MaskCount)
}

const EmptyStack TileStack = 0

const (
	OneBlack TileStack = 0b00000001 << (BitsPerColor * iota)
	OneBlue
	OneOrange
	OneRed
)

var Colors = []TileStack{OneBlack, OneBlue, OneOrange, OneRed}

// returns a + b but with a ceiling value of 2 in each color channel
func SumWithCeiling(a, b TileStack) TileStack {
	m := byte(MaskCount)

	const ceiling = 2 // maximum allowed value
	const nChannels = 4
	var fullResult byte

	var channel byte
	for channel = 0; channel < nChannels; channel++ {
		x := (byte(a)) >> (BitsPerColor * channel)
		y := (byte(b)) >> (BitsPerColor * channel)
		channelSum := (x & m) + (y & m)
		if channelSum > ceiling {
			channelSum = ceiling
		}
		fullResult += byte(channelSum << (BitsPerColor * channel))
	}

	return TileStack(fullResult)
}

// AllColorsGreaterThanOrEqualTo returns true if and only if
//    [a]_i >= [b]_i  for each color i
// Otherwise it returns false
func AllColorsGreaterThanOrEqualTo(a TileStack, b TileStack) bool {
	x := byte(a)
	y := byte(b)

	m := byte(MaskCount)

	ok := (x & m) >= (y & m)

	x = x >> BitsPerColor
	y = y >> BitsPerColor
	ok = ok && ((x & m) >= (y & m))

	x = x >> BitsPerColor
	y = y >> BitsPerColor
	ok = ok && ((x & m) >= (y & m))

	x = x >> BitsPerColor
	y = y >> BitsPerColor
	ok = ok && ((x & m) >= (y & m))

	return ok
}
