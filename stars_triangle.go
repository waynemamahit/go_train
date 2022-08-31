package main

import (
	"fmt"
)

func main() {
	var n int;
	fmt.Print("Masukan jumlah bintang: ");
	fmt.Scan(&n);
	var result string = "";
	var i int = 1;
	var m int = n;
	for i <= n {
		j := m;
		for j > 1 {
			result += " ";
			j--;
		}
		m--;
		s := 1;
		for s <= i {
			result += "*";
			s += 1;
		}
		result += "\n";
		i++;
	}
	fmt.Println(result);
}

