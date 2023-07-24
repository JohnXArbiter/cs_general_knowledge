package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("s", "aAAAZZZZZ"))
}

func numJewelsInStones(jewels string, stones string) int {
	var ans int
	jmap := map[int32]struct{}{}
	for _, j := range jewels {
		jmap[j] = struct{}{}
	}
	for _, s := range stones {
		if _, ok := jmap[s]; ok {
			ans++
		}
	}
	return ans
}
