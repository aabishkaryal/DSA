package problems_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aabishkaryal/DSA/problems"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums: []int{0, 1, 0, 3, 12}}, want: []int{1, 3, 12, 0, 0}},
		{name: "test2", args: args{nums: []int{0, 0, 0, 0, 0}}, want: []int{0, 0, 0, 0, 0}},
		{name: "test3", args: args{nums: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 4, 5}},
		{name: "test4", args: args{nums: []int{0, 0, 0, 0, 1}}, want: []int{1, 0, 0, 0, 0}},
		{name: "test5", args: args{nums: []int{1, 0, 0, 0, 0}}, want: []int{1, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if problems.MoveZeroes(tt.args.nums); !assert.Equal(t, tt.args.nums, tt.want) {
				t.Errorf("test: %s MoveZeroes() = %v, want %v", tt.name, tt.args.nums, tt.want)
			}
		})
	}
}
