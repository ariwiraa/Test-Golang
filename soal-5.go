package main

import "fmt"

func palindrome(kalimat string) bool {
	for i := 0; i < len(kalimat)/2; i++ {
		if kalimat[i] != kalimat[len(kalimat)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	var str = "katak"
	var str2 = "ibu"
	result := palindrome(str)
	result2 := palindrome(str2)

	
	fmt.Println(result)
	fmt.Println(result2)
}
