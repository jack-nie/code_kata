package code_kata

import "testing"

func TestDivision(t *testing.T) {
	for _, c := range []struct {
		want  int
		value int
	}{
		{2, 1},
		{3, 2},
		{5, 3},
	} {
		got := division(c.value)
		if got != c.want {
			t.Errorf("test failed, expected %d, got %d", c.want, got)
		}
	}
}
