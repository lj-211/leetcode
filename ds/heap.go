package ds

func left_idx(i int) int {
	return 2*i + 1
}

func right_idx(i int) int {
	return 2*i + 2
}

type CmpFunc func(interface{}, interface{}) bool

func adjust_heap(in []interface{}, idx int, cmp CmpFunc) {
	in_size := len(in)
	lidx := left_idx(idx)
	ridx := right_idx(idx)

	min_idx := idx
	//if lidx < in_size && in[lidx] < in[idx] {
	if lidx < in_size && cmp(in[lidx], in[idx]) {
		min_idx = lidx
	}
	//if ridx < in_size && in[ridx] < in[min_idx] {
	if ridx < in_size && cmp(in[ridx], in[min_idx]) { //in[ridx] < in[min_idx] {
		min_idx = ridx
	}

	if idx != min_idx {
		tmp := in[idx]
		in[idx] = in[min_idx]
		in[min_idx] = tmp
		adjust_heap(in, min_idx, cmp)
	}
}

func BuildHeap(in []interface{}, cmp CmpFunc) []interface{} {
	hsize := len(in) / 2
	for i := hsize; i >= 0; i-- {
		adjust_heap(in, i, cmp)
	}

	return in
}

func PopTop(in []interface{}, cmp CmpFunc) (heap []interface{}, min interface{}) {
	in_size := len(in)
	if in_size == 0 {
		return []interface{}{}, -1
	}
	tmp := in[0]
	in[0] = in[in_size-1]

	heap = in[0 : in_size-1]

	adjust_heap(heap, 0, cmp)

	return heap, tmp
}

func getParentIdx(idx int) int {
	if idx == 0 {
		return -1
	}
	if idx%2 == 0 {
		return (idx - 2) / 2
	}

	return (idx - 1) / 2
}

func HeapAdd(heap []interface{}, ele interface{}, cm CmpFunc) []interface{} {
	size := len(heap)
	size++
	pidx := getParentIdx(size)
	new_heap := append(heap, ele)

	for pidx != -1 {
		adjust_heap(new_heap, pidx, cm)
		pidx = getParentIdx(pidx)
	}

	return new_heap
}
