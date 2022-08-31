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
	var calc_pd int = 0;
	var calc_sd int = 0;
	for pd < d {
		calc_pd += result[pd][pd];
		calc_sd += result[pd][sd];
		pd += 1;
		sd -= 1;
	}
	hasil := calc_pd - calc_sd;
	fmt.Println(hasil);
}

