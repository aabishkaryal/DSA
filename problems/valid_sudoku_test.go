package problems_test

import (
	"testing"

	"github.com/aabishkaryal/DSA/problems"
)

func TestIsValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		}, want: true},
		{name: "test2", args: args{
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		}, want: false},
		{name: "test3", args: args{
			board: [][]byte{
				{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
				{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
				{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
				{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
				{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
				{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
				{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
			},
		}, want: false},
		{name: "test4", args: args{
			board: [][]byte{
				{'.', '.', '.', '.', '5', '.', '1', '1', '.'},
				{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '1'},
				{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
				{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
				{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
				{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
				{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
			},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.IsValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("Test: %s, IsValidSudoku() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
