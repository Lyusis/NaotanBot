package main

import (
	"fmt"
	_ "fmt"
)

var (
	basicWord  []string
	dictionary []string
)

const pwdLength = 3

func main() {

	basicWord = append(basicWord, "a", "b", "c")
	test := addPassword("")
	fmt.Println(test)

}

func addPassword(pwd string) string {
	for _, word := range basicWord {
		pwd += word
		if len([]rune(pwd)) != pwdLength {
			addPassword(pwd)
		} else {
			dictionary = append(dictionary, pwd)
			fmt.Println(dictionary)
		}
	}
	return pwd
}
