package main

import "testing"

func TestSum(t *testing.T) {

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{5, 20, 20},
	}

	for _, item := range tables {
		maxim := GetMax(item.a, item.b)
		if maxim != item.n {
			t.Errorf("GetMax was incorrect, got %v, expected %v", maxim, item.n)
		}
	}
}
