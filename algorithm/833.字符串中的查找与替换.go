package main

import "fmt"

func main() {
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	var ans []byte
	flag := 0
	for i := 0; i < len(s); {
		flag = getFlag(i, indices)
		if flag != -1 && i == indices[flag] {
			if i+len(sources[flag]) <= len(s) && s[i:i+len(sources[flag])] == sources[flag] {
				ans = append(ans, []byte(targets[flag])...)
				i += len(sources[flag])
				flag++
			} else {
				ans = append(ans, s[i])
				i++
			}
		} else {
			ans = append(ans, s[i])
			i++
		}
	}
	return string(ans)
}

func getFlag(i int, indices []int) int {
	for j := 0; j < len(indices); j++ {
		if i == indices[j] {
			return j
		}
	}
	return -1
}

// cjb
//func findReplaceString(s string, indices []int, sources []string, targets []string) string {
//	ch := []byte(s)
//	res := ""
//	for i := 0; i < len(ch); {
//		idx := in(i, indices)
//		if idx < len(indices) && idx != -1 {
//			if string(ch[indices[idx]:indices[idx]+len(sources[idx])]) == sources[idx] {
//				res += targets[idx]
//				i += len(sources[idx])
//			} else {
//				if len(sources[idx]) > len(ch)-indices[idx] {
//					res += string(ch[indices[idx]:])
//				} else {
//					res += string(ch[indices[idx] : indices[idx]+len(sources[idx])])
//				}
//				i += len(sources[idx])
//			}
//			idx++
//		} else {
//			res += string(ch[i])
//			i++
//		}
//	}
//	c := []byte(res)
//
//	for i, b := range c {
//		fmt.Println(i, " ", b)
//	}
//
//	return res
//}
//
//func in(idx int, arr []int) int {
//	for i := 0; i < len(arr); i++ {
//		if idx == arr[i] {
//			return i
//		}
//	}
//	return -1
//}
