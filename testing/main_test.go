package main

import "testing"

func TestSum(t *testing.T) {
	//total := sum(5, 5)
	//
	//if total != 11 {
	//	t.Errorf("obtuvimos %d y se esperaba %d", total, 11)
	//}
<<<<<<< HEAD

	// slice de struct para testing
=======
>>>>>>> 27e26b9b4a298ffe403935a07558225b7c8708d6
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 3, 4},
		{2, 2, 4},
		{3, 3, 7},
	}
<<<<<<< HEAD
	// recorrido de slice
=======
>>>>>>> 27e26b9b4a298ffe403935a07558225b7c8708d6
	for _, item := range tables {
		total := sum(item.a, item.b)
		if total != item.n {
			t.Errorf("dio %d se esperaba %d", total, item.n)
		}
	}
}
