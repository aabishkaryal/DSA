package problems_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aabishkaryal/DSA/problems"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{digits: []int{1, 2, 3}}, want: []int{1, 2, 4}},
		{name: "test2", args: args{digits: []int{4, 3, 2, 1}}, want: []int{4, 3, 2, 2}},
		{name: "test3", args: args{digits: []int{0}}, want: []int{1}},
		{name: "test4", args: args{digits: []int{9}}, want: []int{1, 0}},
		{name: "test5", args: args{digits: []int{9, 9}}, want: []int{1, 0, 0}},
		{name: "test6", args: args{digits: []int{6, 9, 2}}, want: []int{6, 9, 3}},
		{name: "test7", args: args{digits: []int{9, 9, 9}}, want: []int{1, 0, 0, 0}},
		{name: "test8", args: args{digits: []int{9, 9, 1}}, want: []int{9, 9, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.PlusOne(tt.args.digits); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("test: %s plusOne() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
