package encoding

const StandardBase int = 256

const TargetBase int = 62

var alphabet []byte = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

var lookup = createLookupTable()

func Encode(message []byte) []byte {
	indices := convert(message, StandardBase, TargetBase)

	return translate(indices, alphabet)
}

func Decode(encoded []byte) []byte {
	if !isBase62Encoding(encoded) {
		panic("invalid encoded byte array")
	}

	prepared := translate(encoded, lookup)

	return convert(prepared, TargetBase, StandardBase)
}

func convert(message []byte, sourceBase int, targetBase int) []byte {
	var out []byte

	source := message
	for len(source) > 0 {
		var quotient []byte
		remainder := 0

		for _, s := range source {
			accumulator := int(s&0xFF) + remainder*sourceBase
			digit := (accumulator - (accumulator % targetBase)) / targetBase

			remainder = accumulator % targetBase
			if len(quotient) > 0 || digit > 0 {
				quotient = append(quotient, byte(digit))
			}
		}

		out = append(out, byte(remainder))

		source = quotient
	}

	reverse(out)

	return out
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func translate(indices []byte, dictionary []byte) []byte {
	translation := make([]byte, len(indices))

	for i, index := range indices {
		translation[i] = dictionary[index]
	}

	return translation
}

func createLookupTable() []byte {
	lookup := make([]byte, 256)

	for i, letter := range alphabet {
		lookup[letter] = byte(i & 0xFF)
	}

	return lookup
}

func isBase62Encoding(bytes []byte) bool {
	if len(bytes) == 0 {
		return false
	}

	for _, e := range bytes {
		if '0' > e || '9' < e {
			if 'a' > e || 'z' < e {
				if 'A' > e || 'Z' < e {
					return false
				}
			}
		}
	}

	return true
}