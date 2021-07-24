package xys

import "fmt"

func Abc() {
	fmt.Println("123")
}

// Find test generics
// func Find[T comparable](what T, where []T) int {
// 	for i, v := range where {
// 		if v == what {
// 			return i
// 		}
// 	}
// 	return -1
// }
