package problems_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aabishkaryal/DSA/problems"
)

func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{2, 2, 1}}, want: 1},
		{name: "test2", args: args{nums: []int{4, 1, 2, 1, 2}}, want: 4},
		{name: "test3", args: args{nums: []int{1}}, want: 1},
		{name: "test4", args: args{nums: []int{2, 2, 4, 6, 8, 4, 8, 3, 5, 6, 3, 9, 9}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.SingleNumber(tt.args.nums); got != tt.want {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
