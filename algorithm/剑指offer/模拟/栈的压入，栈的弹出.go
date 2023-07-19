package main

func validateStackSequences(pushed, popped []int) bool {
	var stack []int
	i := 0
	for _, e := range pushed {
		stack = append(stack, e)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return len(stack) == 0
}
