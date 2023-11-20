package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("./finance.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var command, res string
	var lines = []string{}

	myFin := Finance{}

	fmt.Println(res)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	res = lines[0]
	myFin.balance = 1
	for {
		fmt.Println("Введите команду или LIST")
		fmt.Fscan(os.Stdin, &command)
		if command == "EXIT" { //EXIT FROM FOR
			fmt.Println("exit from program")
			break
		}
		switch command {
		case "LIST":
			fmt.Println("List")
		case "NEWBALANCE":
			NewBalance(lines)
		case "CURRENT":
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
		res = strings.Contains(v, "BALANCE")
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = "BALANCE -> " + str

}
