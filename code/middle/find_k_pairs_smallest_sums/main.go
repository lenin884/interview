package find_k_pairs_smallest_sums

/*
Нужно разобраться с кучами
*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var heap [][]int

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if len(heap) == k && heap[0][0]+heap[0][1] <= nums1[i]+nums2[j] {
				break
			}

			heap = append(heap, []int{nums1[i], nums2[j]})
			heapUp(heap, len(heap)-1)

			if len(heap) > k {
				heap[0] = heap[len(heap)-1]
				heap = heap[:len(heap)-1]
				heapDown(heap, 0, len(heap)-1)
			}
		}
	}

	return heap
}

func heapDown(heap [][]int, p, limit int) {
	l, r := 2*p+1, 2*p+2
	larger := p

	if l <= limit && heap[l][0]+heap[l][1] > heap[larger][0]+heap[larger][1] {
		larger = l
	}

	if r <= limit && heap[r][0]+heap[r][1] > heap[larger][0]+heap[larger][1] {
		larger = r
	}

	if larger != p {
		heap[p], heap[larger] = heap[larger], heap[p]
		heapDown(heap, larger, limit)
	}
}

func heapUp(heap [][]int, p int) {
	parent := (p - 1) / 2

	if parent >= 0 && heap[p][0]+heap[p][1] > heap[parent][0]+heap[parent][1] {
		heap[parent], heap[p] = heap[p], heap[parent]
		heapUp(heap, parent)
	}
}
