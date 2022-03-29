package main

import "fmt"

func main() {
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
	if num := 17; num%2 == 1 {
		fmt.Println("Number is odd.")
	} else if num%2 == 0 {
		fmt.Println("Number is even.")
	}

	if num := 17; num < 10 {
		fmt.Println("Number is single digit.")
	} else {
		fmt.Println("Number is multi-digit.")
	}
}
