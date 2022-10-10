package problems_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aabishkaryal/DSA/problems"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test1", args: args{s: []byte("hello")}, want: []byte("olleh")},
		{name: "test2", args: args{s: []byte("Hannah")}, want: []byte("hannaH")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]byte, len(tt.args.s))
			copy(original, tt.args.s)
			if problems.ReverseString(tt.args.s); !assert.Equal(t, tt.want, tt.args.s) {
				t.Errorf("Test: %s, ReverseString() = %v, want %v", tt.name, tt.args.s, tt.want)
			}
		})
	}
}
