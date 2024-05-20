package reflectx

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	vals := []any{
		true,
		int(0),
		int32(1),
		int64(2),
		float32(3),
		float64(4),
		complex64(5),
		complex128(6),
		new(int),
		new(int),
		[2]int{7, 8},
		struct{ a, b int }{9, 10},
		[]int{11, 12},
		map[int]string{13: "foo"},
		any(14),
	}

	var seen []uint64

	r := require.New(t)

	for _, v := range vals {
		h := Hash(v)
		t.Logf("%v (%[1]T) => %d", v, h)
		r.NotEqual(0, h)

		r.False(slices.Contains(seen, h))

		seen = append(seen, h)
	}
}
