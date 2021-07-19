package encoding

import "testing"

var tests = []struct {
	decoded string
	encoded string
}{
	{"000000000000000000000000000000000000000", "ypA44DH5DGrsZ9a4rJwHVNXkaf4DBKwMEkJzDDBsy6CWYr1HNyuu"},
	{"testing@example.com", "41NQfwZe8tpCaIcRnedjpMRGwv"},
	{"https://example.com/query?test=true", "S7es9fr0BB8VmDqkLswCGmVfEzXv5yGM16NihEY5j8bNMBB"},
	{"Lorem ipsum dolor sit amet", "E3hStfZPIJA0Sk63V4EZkumvKp9oDfUe6yK"},
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		if output := Encode([]byte(test.decoded)); string(output) != test.encoded {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.decoded, test.encoded, string(output))
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range tests {
		if output, _ := Decode([]byte(test.encoded)); string(output) != test.decoded {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.encoded, test.decoded, string(output))
		}
	}

	if _, err := Decode([]byte("")); err == nil {
		t.Error("Test Failed: decode should fail on empty input")
	}

	if _, err := Decode([]byte("%&")); err == nil {
		t.Error("Test Failed: decode should fail on invalid chars")
	}
}
