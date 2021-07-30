package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex = regexp.MustCompile("[a-zA-Z0-9\\+\\.\\_\\%\\-\\+]{1,20}" +
		"\\@" +
		"[.]" +
		"[a-zA-Z0-9][a-zA-Z0-9\\-]{0,64}" +
		"(" +
		"\\." +
		"([co]{2}|[id]{2})" +
		")+")

	var str string
	fmt.Println("input email : ")
	fmt.Scan(&str)
	fmt.Println(regex.MatchString(str))

}
