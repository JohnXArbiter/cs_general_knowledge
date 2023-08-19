package main

import "fmt"

func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func main() {
	asd := []int{1, 2, 3, 4, 5}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("阿斯达斯的")
		}
	}()
	go modify(asd)
	fmt.Println(asd)

}

func modify(asd []int) {

	asd[0] = 111
	asd[10] = 2
}
