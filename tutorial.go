package main

import "fmt"

func completeName(input *string) {
	if *input == "Reginald" {
		*input = "Reginald Bayoan"
	} else if *input == "Tyson" {
		*input = "Tyson Talanya"
	} else {
		*input = "Teacher not registered"
	}
}

func main2() {
	x := "Clifford"
	completeName(&x)

	fmt.Println(x)
}
