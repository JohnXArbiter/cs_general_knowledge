package main

import "math"

func change(a, b []int, ai, bi int) int {
	if ai == len(a) && bi == len(b) {
		return 0
	}
	if ai == len(a) && bi < len(b) {
		return change(a, b, ai, bi+1)
	}
	if ai < len(a) && bi == len(b) {
		return change(a, b, ai+1, bi)
	}
	// 选择1：删掉 a[ai]
	p1 := a[ai] + change(a, b, ai+1, bi)
	// 选择2：删掉 b[bi]
	p2 := b[bi] + change(a, b, ai, bi+1)
	// 选择3：同时删掉，不仅开销最高，也会被前两种顶调
	// p3 := a[ai] + b[bi] + change(a, b, ai+1, bi+1)
	// 选择4：变，a[ai] -> b[bi] or b[bi] -> a[ai]
	p4 := int(math.Abs(float64(a[ai]-b[bi]))) + change(a, b, ai+1, bi+1)
	res := math.Min(float64(p1), math.Min(float64(p2), float64(p4)))
	return int(res)
}
