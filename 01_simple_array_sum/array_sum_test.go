package arraysum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArraySum(t *testing.T) {
	tests := []struct {
		name          string
		arraySize     int
		arrayElements []int
		want          int
	}{
		{
			name:          "test 1",
			arraySize:     6,
			arrayElements: []int{1, 2, 3, 4, 10, 11},
			want:          31,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(st *testing.T) {
			got := ArraySum(tc.arraySize, tc.arrayElements)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				st.Fatal(diff)
			}
		})
	}
}
