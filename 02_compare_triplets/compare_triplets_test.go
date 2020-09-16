package triplets

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCompareTriplets(t *testing.T) {
	tests := []struct {
		name  string
		alice []int
		bob   []int
		want  []int
	}{
		{
			name:  "test 0",
			alice: []int{5, 6, 7},
			bob:   []int{3, 6, 10},
			want:  []int{1, 1},
		},
		{
			name:  "test 1",
			alice: []int{17, 28, 30},
			bob:   []int{99, 16, 8},
			want:  []int{2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(st *testing.T) {
			got := CompareTriplets(tc.alice, tc.bob)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				st.Error(diff)
			}
		})
	}
}
