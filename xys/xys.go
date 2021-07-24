package xys

import "fmt"

func Abc() {
	fmt.Println("123")
}

// // func Find[T any, []T any](what T, where []T) (idx int) {
// func Find[T comparable](what T, where []T) int {
// 	for i, v := range where {
// 		if v == what {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func mapFunc[T any, M any](a []T, f func(T) M) []M {
// 	n := make([]M, len(a), cap(a))
// 	for i, e := range a {
// 		n[i] = f(e)
// 	}
// 	return n
// }
