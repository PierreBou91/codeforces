package main

import (
	"fmt"
)

func main() {
	var s string
	hello := "hello"
	fmt.Scan(&s)
	lastIndex := -1
	count := 0

outer:
	for _, v := range hello {
		for i, v2 := range s {
			if v == v2 && lastIndex < i {
				lastIndex = i
				count++
				continue outer
			}
		}
	}
	if count == 5 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
