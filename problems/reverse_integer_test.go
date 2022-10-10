package problems_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aabishkaryal/DSA/problems"
)

func TestReverseInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{x: 123}, want: 321},
		{name: "test2", args: args{x: -123}, want: -321},
		{name: "test3", args: args{x: 120}, want: 21},
		{name: "test4", args: args{x: 0}, want: 0},
		{name: "test5", args: args{x: 1534236469}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.ReverseInt(tt.args.x); !assert.Equal(t, tt.want, got) {
				t.Errorf("Test %s: ReverseInt() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
