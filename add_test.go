package main

import "testing"

var table = []struct{
	x int
	y int
	want int
}{
	{2,2,4},
	{7,4,11},
	{20,12,32},
	{476,98,574},
}

func test_add(t *testing.T) {
	for _, row := range table {
		got := add(row.x, row.y)
		if got != row.want {
			t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
		}
	}
}