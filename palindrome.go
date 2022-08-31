package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringinput string;
	fmt.Print("Masukan karakter: "); 
	fmt.Scan(&stringinput);
	stringinput = strings.ToLower(stringinput);
	result := strings.Split(stringinput, "")
	i := 0;
	j := len(result) - 1;
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++;
		j--;
	}
	
	if stringinput == strings.Join(result, "") {
		fmt.Println("Yes");
	} else {
		fmt.Println("No");
	}
}
