// Package merge combines three already-sorted integer collections into one ascending slice.
package merge

// Merge returns every value of the three collections in ascending order.
//
// collection1 must already be sorted from max to min, collection2 and collection3 from min to
// max. Merge reads collection1 from its last element backwards, so all three are consumed
// smallest value first and each step only has to compare the next unread value of each one.
// It runs in O(n) time for n values in total and never calls a sort function.
//
// The inputs are not modified. The result is never nil. If an input is not in the order
// described above the result is not sorted.
func Merge(collection1, collection2, collection3 []int) []int {
	total := len(collection1) + len(collection2) + len(collection3)
	merged := make([]int, 0, total)

	// next unread position in each collection: collection1 is walked from its end
	i, j, k := len(collection1)-1, 0, 0

	for len(merged) < total {
		source, smallest := 0, 0
		if i >= 0 {
			source, smallest = 1, collection1[i]
		}
		if j < len(collection2) && (source == 0 || collection2[j] < smallest) {
			source, smallest = 2, collection2[j]
		}
		if k < len(collection3) && (source == 0 || collection3[k] < smallest) {
			source, smallest = 3, collection3[k]
		}

		merged = append(merged, smallest)
		switch source {
		case 1:
			i--
		case 2:
			j++
		default:
			k++
		}
	}
	return merged
}
