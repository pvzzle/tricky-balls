package main

import "testing"

func TestCanSort(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int64
		want   bool
	}{
		{"1x1 positive", [][]int64{{2}}, true},
		{"1x1 zero", [][]int64{{0}}, true},
		{"2x2 sortable", [][]int64{{1, 1}, {1, 1}}, true},
		{"2x2 unsortable", [][]int64{{2, 2}, {1, 1}}, false},
		{"3x3 sortable", [][]int64{{0, 2, 1}, {1, 1, 1}, {2, 0, 0}}, true},
		{"3x3 unsortable", [][]int64{{1, 3, 1}, {2, 1, 2}, {3, 3, 3}}, false},
		{"4x4 sortable", [][]int64{{3, 1, 2, 0}, {2, 2, 1, 1}, {1, 3, 2, 0}, {0, 0, 1, 3}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := canSort(tt.matrix)
			if res != tt.want {
				t.Errorf("\nresult: %v\nwant: %v", res, tt.want)
			}
		})
	}
}
