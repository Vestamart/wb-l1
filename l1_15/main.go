package main

import "fmt"

var justString string

func createHugeString(size int) string {
	s := make([]byte, size)
	for i := range s {
		s[i] = 'a'
	}
	return string(s)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
	fmt.Println("justString (длина):", len(justString))
}
