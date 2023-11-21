package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("./finance.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_SYNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var command, res string
	var lines = []string{}

	fmt.Println(res)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	//check empty list finance.txt
	if len(lines) == 0 {
		lines = CreatEmptyList(lines)
	}

	for {
		fmt.Println("Введите команду или \"list\"")
		fmt.Fscan(os.Stdin, &command)
		if command == "exit" { //EXIT FROM FOR
			fmt.Println("exit from program")
			Write(*file, lines)
			break
		}
		switch command {
		case "list":
			fmt.Println("List")
		case "newbalance":
			NewBalance(lines)
		case "current":
			res := "current -- \n"
			file.WriteString(res)
		} //END SWITCH

	}
}

func NewBalance(newLines []string) {
	fmt.Println("Введите новый баланс")
	str := ""
	targetLine := 0
	fmt.Fscan(os.Stdin, &str)
	for i, v := range newLines {
		var res bool
		res = strings.Contains(v, "balance")
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = "balance -> " + str
}

func CreatEmptyList(slice []string) []string {
	slice = append(slice, "balance -> 0 ")
	slice = append(slice, "target -> 0 ")
	slice = append(slice, "salary -> 0 ")
	slice = append(slice, "expenses -> 0 ")
	slice = append(slice, "currentDay -> 0 ")
	slice = append(slice, "beforeDay -> 0 ")
	slice = append(slice, "moneyInDay -> 0 ")
	return slice
}
