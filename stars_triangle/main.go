package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukan jumlah bintang: ")
	fmt.Scan(&n)
	var result string = ""
	var m int = n
	for i := 1; i <= n; i++ {
		for j := m; j > 1; j-- {
			result += " "
		}
		m--
		for s := 1; s <= i; s++ {
			result += "*"
		}
		result += "\n"
	}
	fmt.Println(result)
}
