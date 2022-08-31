package main

import (
	"fmt"
)

func main() {
	var n int;
	fmt.Print("Masukan nilai N: ");
	fmt.Scan(&n);
	var i int = 0;
	var result int = 0;
	for i != n && i < n {
		var p int = 0;
		for p < 4 {
			j := i;
			if p < 2 {
				i += 1;
			} else {
				i *= 2;
			}
			if i == n {
				result = j;
				break;
			}
			p++;
		}
	}
	fmt.Println(result);
}

