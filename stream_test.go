package radiko

import "testing"

func TestGetStreamMultiURL(t *testing.T) {
	items, err := GetStreamMultiURL("LFR")
	if err != nil {
		t.Error(err)
	}

	const expected = 4
	if actual := len(items); expected != actual {
		t.Errorf("expected %d, but %d", expected, actual)
	}
}
