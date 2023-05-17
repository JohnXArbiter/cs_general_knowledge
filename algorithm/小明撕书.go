package main

import "fmt"

func main() {
	var start int
	var end int
	fmt.Printf("请输入起始页和末页：\n")
	fmt.Scanf("%d%d", &start, &end)

	var page int
	if start%2 == 0 {
		page = (end - start + 1) / 2
	} else {
		page = (end-start+1)/2 + 1
	}
	fmt.Printf("需要撕掉：%d", page)
}
