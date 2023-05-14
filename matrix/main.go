package main

import (
	"fmt"
)

func main() {
	var d int;
	fmt.Print("Masukan jumlah dimensi matriks: ");
	fmt.Scan(&d);
	result := [][]int{};
	var pd int = 0;
	var sd int = d - 1;
	var calcPd int
	var calcSd int
	for pd < d {
		calcPd += result[pd][pd];
		calcSd += result[pd][sd];
		pd += 1;
		sd -= 1;
	}
	hasil := calcPd - calcSd;
	fmt.Println(hasil);
}

