package main

import (
	"fmt"
	"generics/xys"
)

func mapFunc[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

// gotip run -gcflags=-G=3 main.go
func main() {
	// a := []int{1, 2, 3}
	// xys.Find(1, a)
	xys.Abc()
	vi := []int{1, 2, 3, 4, 5, 6}
	vs := mapFunc(vi, func(v int) string {
		return "<" + fmt.Sprint(v) + ">"
	})
	fmt.Println(vs)
}
