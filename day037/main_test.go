package main

import "testing"

func TestSum(t *testing.T) {
	/* 	total := Sum(3, 4)

	   	if total != 7 {
	   		t.Errorf("Sum not correct, expected: %v, got: %v", total, 7)
	   	} */

	// test tables

	tables := []struct {
		x, y, n int
	}{
		{1, 1, 2},
		{5, 6, 11},
		{10, 10, 20},
		{15, 3, 18},
		{19, 1, 20},
		{12, 1, 13},
		{3, 4, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) is not correct, expected: %d, got: %d.", table.x, table.y, table.n, total)
		}
	}
}

// go test -cover -coverprofile=c.out
// go tool cover -html=c.out -o coverage.html
// generates and html page which shows the test coverage
