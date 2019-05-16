package fibonacci

import "testing"

var fibTests = []struct {
	a        int
	expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{20, 10946},
	{42, 433494437},
}

func TestSimple(t *testing.T) {
	for _, ft := range fibTests {
		if v := FibSimple(ft.a); v != ft.expected {
			t.Errorf("FibSimple(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}
