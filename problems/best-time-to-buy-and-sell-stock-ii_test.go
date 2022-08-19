package problems_test

import (
	"testing"

	"github.com/aabishkaryal/DSA/problems"
)

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 1", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 7},
		{name: "Test 2", args: args{prices: []int{1, 2, 3, 4, 5}}, want: 4},
		{name: "Test 3", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
		{name: "Test 4", args: args{prices: []int{7, 1, 5, 3, 6, 4, 2}}, want: 7},
		{name: "Test 5", args: args{prices: []int{}}, want: 0},
		{name: "Test 6", args: args{prices: []int{1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.MaxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
