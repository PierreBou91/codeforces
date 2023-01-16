// https://codeforces.com/problemset/problem/236/A
package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)
	mmap := make(map[rune]int)
	for i, v := range name {
		mmap[v] = i
	}
	if len(mmap)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
		return
	}
	fmt.Println("IGNORE HIM!")
}
