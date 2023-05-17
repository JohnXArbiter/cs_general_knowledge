package main

import "fmt"

type Array struct {
	data   []any
	length int
}

func main() {
	array := []any{312, 23, 32, 11, 5345, 3213, 13}
	arrayList := initArray(array)
	agetAll(arrayList)
	aget(arrayList, 2)
	//aremove(arrayList, 3)
	agetAll(arrayList)
	adestroy(arrayList)
}

func initArray(data []any) (array *Array) {
	array = new(Array)
	var l Array
	l.data = data
	l.length = 7
	array = &l
	return array
}

func agetAll(list *Array) {
	fmt.Println("所有数据：")
	for i := 0; i < list.length; i++ {
		fmt.Printf("%d ", list.data[i])
	}
	fmt.Println()
}

func aget(list *Array, index int) {
	for i := 0; i < list.length; i++ {
		if i == index-1 {
			fmt.Printf("%d ", list.data[i])
			break
		}
	}
}

//func aremove(list *Array, index int) {
//	var array []any
//	copy(array, list.data)
//	var j int
//	for i := 0; i < list.length; i++ {
//		if i == index-1 {
//			fmt.Printf("删除的数据：%d ", list.data[i])
//			j = i
//			break
//		}
//	}
//	list.data[j:] = array[j+1:]
//}

func adestroy(list *Array) {

}
