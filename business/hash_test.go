package business

import "testing"

func TestEncodeDecodeRoundTrip(t *testing.T) {
	ids := []int64{0, 1, 61, 62, 700, 999999, 9223372036854775807}

	for _, id := range ids {
		encoded := Encode(id)
		decoded := Decode(encoded)
		if decoded != id {
			t.Errorf("round trip failed: id=%d, encoded=%q, decoded=%d", id, encoded, decoded)
		}
	}
}
