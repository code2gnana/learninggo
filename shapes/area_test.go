package shapes

import "testing"

func TestArea(t *testing.T) {
	s := square{sideLength: 20}
	t := triangle{base: 10, height: 20}

	got := getArea(s)
	want := 400

	t.Errorf("got %q, wanted %q", got, want)

}
