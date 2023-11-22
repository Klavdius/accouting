package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./finance.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var command, res string
	var lines = []string{}

	date := make([]byte, 0)
	fmt.Println(res)
	_, err = file.Read(date)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(date)
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

		iter, ok := allMap[command]
		if ok != false {
			Action(lines, iter)
		} else {
			fmt.Println("Нет такой команды(( наберите list")
		}
	}
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
