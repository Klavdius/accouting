package main

import (
	"fmt"
	"os"
)

func main() {
	res := 0
	res = CheckFile()
	if res == 1 {
		fmt.Println("ups")
	}
}

func CheckFile() int {
	file, err := os.OpenFile("./finance.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var int = 1
	return int
}
