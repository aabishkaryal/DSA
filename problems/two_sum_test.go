package problems_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aabishkaryal/DSA/problems"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		want []int
		args args
	}{
		{name: "test1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "test2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "test3", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
		{name: "test4", args: args{nums: []int{3, 3}, target: 7}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.TwoSum(tt.args.nums, tt.args.target); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("test: %s, TwoSum() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
