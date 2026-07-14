package business

import "testing"

func TestEncodeDecodeRoundTrip(t *testing.T) {
	ids := []uint64{0, 1, 61, 62, 700, 999999, 18446744073709551615}

	for _, id := range ids {
		encoded := Encode(id)
		decoded := Decode(encoded)
		if decoded != id {
			t.Errorf("round trip failed: id=%d, encoded=%q, decoded=%d", id, encoded, decoded)
		}
	}
}
