package main

import "fmt"

func reverseString(kalimat string) string {
	strings := []rune(kalimat)
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}

	return string(strings)
}

func main() {
	var str = "Aplikasi"
	fmt.Println(reverseString(str))
}
