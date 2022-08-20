package problems_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aabishkaryal/DSA/problems"
)

func TestRotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		want []int
		args args
	}{
		{name: "test1", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test2", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 1}, want: []int{7, 1, 2, 3, 4, 5, 6}},
		{name: "test3", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 7}, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "test4", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 0}, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "test5", args: args{nums: []int{1}, k: 8}, want: []int{1}},
		{name: "test6", args: args{nums: []int{1}, k: 0}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			problems.Rotate(tt.args.nums, tt.args.k)
			require.Equal(t, tt.want, tt.args.nums)
		})
	}
}
