package squareAdder

import "testing"

func Test_squareAdder(t *testing.T) {
	want := 7
	result, err := squareAdder(7, 7, 0)
	if result != want {
		t.Errorf("got: %d, want: %d, %v", result, want, err.Error())
	}
}
