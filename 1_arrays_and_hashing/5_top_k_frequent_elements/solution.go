package topkfrequentelements

import (
	"container/heap"
	"sort"
)

// Time complexity: O(n * log n)
// Space complexity: O(n)
func TopKFrequentSort(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	type pair struct {
		num  int
		freq int
	}

	pairs := make([]pair, 0, len(freqMap))
	for num, freq := range freqMap {
		pairs = append(pairs, pair{num, freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}
	return result
}

type freqPair struct {
	num  int
	freq int
}

type minHeap []freqPair

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(freqPair)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// Time complexity: O(n * log k), where n is the number of elements in nums and k is the number of top frequent elements to return.
// Space complexity: O(n)
func TopKFrequentHeap(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	h := &minHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, freqPair{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(freqPair).num)
	}
	return result
}

// Time complexity: O(n)
// Space complexity: O(n)
func TopKFrequentBucket(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// max possible frequency is len(nums)
	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}
	return result[:k]
}
