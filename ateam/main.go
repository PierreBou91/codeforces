// https://codeforces.com/contest/231/problem/A
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 0; i < n; i++ {
		var p, v, t int
		fmt.Scan(&p, &v, &t)
		if p+v+t >= 2 {
			count++
		}
	}
	fmt.Println(count)
}
