package merge_test

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"

	merge "github.com/Nutto555/merge-sorted"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name       string
		c1, c2, c3 []int
		want       []int
	}{
		{
			name: "three collections of different lengths",
			c1:   []int{9, 7, 3},
			c2:   []int{1, 4, 8, 10},
			c3:   []int{2, 5},
			want: []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
		},
		{
			name: "all empty",
			c1:   []int{},
			c2:   []int{},
			c3:   []int{},
			want: []int{},
		},
		{
			name: "all nil",
			want: []int{},
		},
		{
			name: "only the descending collection has values",
			c1:   []int{5, 3, 1},
			want: []int{1, 3, 5},
		},
		{
			name: "only the second collection has values",
			c2:   []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "only the third collection has values",
			c3:   []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "descending collection empty",
			c2:   []int{2, 4},
			c3:   []int{1, 3},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "one ascending collection empty",
			c1:   []int{6, 2},
			c3:   []int{1, 7},
			want: []int{1, 2, 6, 7},
		},
		{
			name: "single element each",
			c1:   []int{2},
			c2:   []int{3},
			c3:   []int{1},
			want: []int{1, 2, 3},
		},
		{
			name: "duplicates within a collection",
			c1:   []int{4, 4, 1},
			c2:   []int{2, 2},
			c3:   []int{3, 3, 3},
			want: []int{1, 2, 2, 3, 3, 3, 4, 4},
		},
		{
			name: "the same values in every collection",
			c1:   []int{2, 1},
			c2:   []int{1, 2},
			c3:   []int{1, 2},
			want: []int{1, 1, 1, 2, 2, 2},
		},
		{
			name: "every value identical",
			c1:   []int{7, 7},
			c2:   []int{7},
			c3:   []int{7, 7, 7},
			want: []int{7, 7, 7, 7, 7, 7},
		},
		{
			name: "collections do not overlap, descending one is lowest",
			c1:   []int{3, 2, 1},
			c2:   []int{4, 5},
			c3:   []int{6, 7},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "collections do not overlap, descending one is highest",
			c1:   []int{9, 8},
			c2:   []int{4, 5},
			c3:   []int{1, 2},
			want: []int{1, 2, 4, 5, 8, 9},
		},
		{
			name: "negative numbers and zero",
			c1:   []int{5, 0, -5},
			c2:   []int{-10, -1, 0},
			c3:   []int{-3, 2},
			want: []int{-10, -5, -3, -1, 0, 0, 2, 5},
		},
		{
			name: "one long and two short collections",
			c1:   []int{100},
			c2:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			c3:   []int{0},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge.Merge(tt.c1, tt.c2, tt.c3)
			if got == nil {
				t.Fatal("Merge returned nil, want a non-nil slice")
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("Merge(%v, %v, %v) = %v, want %v", tt.c1, tt.c2, tt.c3, got, tt.want)
			}
		})
	}
}

func TestMergeLeavesInputsUntouched(t *testing.T) {
	c1, c2, c3 := []int{9, 5, 1}, []int{2, 6}, []int{3, 4, 10}
	before1, before2, before3 := slices.Clone(c1), slices.Clone(c2), slices.Clone(c3)

	merge.Merge(c1, c2, c3)

	if !slices.Equal(c1, before1) || !slices.Equal(c2, before2) || !slices.Equal(c3, before3) {
		t.Errorf("inputs changed: got %v %v %v, want %v %v %v", c1, c2, c3, before1, before2, before3)
	}
}

// TestMergeRandomInputs checks two properties on generated inputs: the result is in ascending
// order, and it holds exactly the values that went in. Neither check needs a sort.
func TestMergeRandomInputs(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	for round := 0; round < 500; round++ {
		c1 := descending(rng, rng.Intn(40))
		c2 := ascending(rng, rng.Intn(40))
		c3 := ascending(rng, rng.Intn(40))

		got := merge.Merge(c1, c2, c3)

		for i := 1; i < len(got); i++ {
			if got[i-1] > got[i] {
				t.Fatalf("round %d: result not ascending at index %d: %v", round, i, got)
			}
		}
		want := map[int]int{}
		for _, c := range [][]int{c1, c2, c3} {
			for _, v := range c {
				want[v]++
			}
		}
		for _, v := range got {
			want[v]--
		}
		for v, n := range want {
			if n != 0 {
				t.Fatalf("round %d: value %d: count in the result differs from the inputs by %d", round, v, -n)
			}
		}
	}
}

// ascending builds a min-to-max slice by adding non-negative steps to a random start.
func ascending(rng *rand.Rand, n int) []int {
	out := make([]int, n)
	v := rng.Intn(41) - 20
	for i := range out {
		v += rng.Intn(4)
		out[i] = v
	}
	return out
}

// descending builds a max-to-min slice by filling an ascending sequence from the back.
func descending(rng *rand.Rand, n int) []int {
	out := make([]int, n)
	v := rng.Intn(41) - 20
	for i := n - 1; i >= 0; i-- {
		v += rng.Intn(4)
		out[i] = v
	}
	return out
}

func ExampleMerge() {
	fmt.Println(merge.Merge([]int{9, 7, 3}, []int{1, 4, 8}, []int{2, 5}))
	// Output: [1 2 3 4 5 7 8 9]
}
