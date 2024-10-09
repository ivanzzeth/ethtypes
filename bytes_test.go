package ethtypes

import "testing"

func TestBytes(t *testing.T) {
	b1 := Bytes([]byte{1, 2, 3, 4})
	js, err := b1.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("js: %v", string(js))

	decodedB1 := Bytes{}
	err = decodedB1.UnmarshalJSON(js)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("decoded: %v", decodedB1)
}
