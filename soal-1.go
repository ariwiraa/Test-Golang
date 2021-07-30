package main

import "fmt"

func main() {
	var a int
	fmt.Println("Masukan angka : ")
	fmt.Scan(&a)

	if a%3 == 0 && a%5 == 0 {
		fmt.Println("Hello World")
	} else if a%5 == 0 {
		fmt.Println("World")
	} else {
		fmt.Println("Hello")
	}
}
