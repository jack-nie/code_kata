package code_kata

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, c := range []struct {
		want []int
		arr  []int
		low  int
		high int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{2, 3, 7, 6, 4, 5, 1}, 0, 6},
	} {
		quickSort(c.arr, c.low, c.high)
		if !reflect.DeepEqual(c.arr, c.want) {
			t.Errorf("test failed, expected %s, got %s", c.want, c.arr)
		}
	}
}
