package code_kata

import "testing"

func TestChop(t *testing.T) {
	for _, c := range []struct {
		want   int
		arr    []int
		target int
	}{
		{-1, []int{}, 3},
		{-1, []int{1}, 3},
		{0, []int{1}, 1},
		{0, []int{1, 3, 5}, 1},
		{1, []int{1, 3, 5}, 3},
		{2, []int{1, 3, 5}, 5},
		{-1, []int{1, 3, 5}, 0},
		{-1, []int{1, 3, 5}, 2},
		{-1, []int{1, 3, 5}, 4},
		{-1, []int{1, 3, 5}, 6},
		{0, []int{1, 3, 5, 7}, 1},
		{1, []int{1, 3, 5, 7}, 3},
		{2, []int{1, 3, 5, 7}, 5},
		{3, []int{1, 3, 5, 7}, 7},
		{-1, []int{1, 3, 5, 7}, 0},
		{-1, []int{1, 3, 5, 7}, 2},
		{-1, []int{1, 3, 5, 7}, 4},
		{-1, []int{1, 3, 5, 7}, 6},
		{-1, []int{1, 3, 5, 7}, 8},
	} {
		got := chop(c.target, c.arr)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}
