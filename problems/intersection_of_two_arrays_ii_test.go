package problems_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aabishkaryal/DSA/problems"
)

func TestIntersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{name: "test2", args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
		{name: "test3", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2}}, want: []int{2}},
		{name: "test4", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{3}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.Intersect(tt.args.nums1, tt.args.nums2); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectSorted(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums1: []int{1, 1, 2, 2}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{name: "test2", args: args{nums1: []int{4, 5, 9}, nums2: []int{4, 4, 9, 8, 9}}, want: []int{4, 9}},
		{name: "test3", args: args{nums1: []int{1, 1, 2, 2}, nums2: []int{2}}, want: []int{2}},
		{name: "test4", args: args{nums1: []int{1, 1, 2, 2}, nums2: []int{3}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.IntersectSorted(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Testname: %s, IntersectSorted() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
