package problems_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aabishkaryal/DSA/problems"
)

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}}, want: false},
		{name: "test2", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 1}}, want: true},
		{name: "test3", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}}, want: false},
		{name: "test4", args: args{nums: []int{1}}, want: false},
		{name: "test5", args: args{nums: []int{1, 1}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.ContainsDuplicate(tt.args.nums); got != tt.want {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
