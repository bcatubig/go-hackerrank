package rot

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		name      string
		rotations int
		in        []int
		want      []int
	}{
		{
			name:      "case 0",
			rotations: 4,
			in:        []int{1, 2, 3, 4, 5},
			want:      []int{5, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(st *testing.T) {
			got := RotateLeft(tc.in, tc.rotations)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				st.Error(diff)
			}
		})
	}
}
