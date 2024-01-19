package seven

import "testing"

func TestGetNumber(t *testing.T) {
	res := GetNumber()

	if res != number {
		t.Fail()
	}
}
