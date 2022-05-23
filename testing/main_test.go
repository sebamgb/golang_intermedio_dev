package main

import "testing"

func TestSum(t *testing.T) {
	//total := sum(5, 5)
	//
	//if total != 11 {
	//	t.Errorf("obtuvimos %d y se esperaba %d", total, 11)
	//}

	// slice de struct para testing
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 3, 4},
		{2, 2, 4},
		{3, 3, 6},
	}

	// recorrido de slice
	for _, item := range tables {
		total := sum(item.a, item.b)
		if total != item.n {
			t.Errorf("dio %d se esperaba %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		r int
	}{
		{3, 2, 3},
		{4, 2, 4},
		{2, 4, 5},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.r {
			t.Error("got", max, "expected", item.r)
		}
	}
}
