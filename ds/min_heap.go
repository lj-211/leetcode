package ds

func left_idx(i int) int {
	return 2*i + 1
}

func right_idx(i int) int {
	return 2*i + 2
}

type CmpFunc func(interface{}, interface{}) bool

// min heap
func adjust_min_heap(in []interface{}, idx int, cmp CmpFunc) {
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
		adjust_min_heap(in, min_idx, cmp)
	}
}

func BuildMinHeap(in []int, cmp CmpFunc) []int {
	hsize := len(in) / 2
	for i := hsize; i >= 0; i-- {
		adjust_min_heap(in, i, cmp)
	}

	return in
}

func PopMinTop(in []int) (heap []int, min int) {
	in_size := len(heap)
	if in_size == 0 {
		return []int{}, -1
	}
	tmp := in[0]
	in[0] = in[in_size-1]
	return in[0 : in_size-1], tmp
}

// max heap
func adjust_max_heap(in []int, idx int) {
	in_size := len(in)
	lidx := left_idx(idx)
	ridx := right_idx(idx)

	max_idx := idx
	if lidx < in_size && in[lidx] > in[idx] {
		max_idx = lidx
	}
	if ridx < in_size && in[ridx] > in[max_idx] {
		max_idx = ridx
	}

	if idx != max_idx {
		tmp := in[idx]
		in[idx] = in[max_idx]
		in[max_idx] = tmp
		adjust_max_heap(in, max_idx)
	}
}

func BuildMaxHeap(in []int) []int {
	hsize := len(in) / 2
	for i := hsize; i >= 0; i-- {
		adjust_max_heap(in, i)
	}

	return in
}

func PopMaxTop(in []int) (heap []int, max int) {
	in_size := len(in)
	if in_size == 0 {
		return []int{}, -1
	}
	tmp := in[0]
	in[0] = in[in_size-1]
	in = in[0 : in_size-1]
	adjust_max_heap(in, 0)

	return in[0 : in_size-1], tmp
}
