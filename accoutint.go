package main

import (
	"fmt"
	"os"
)

var lines = []string{}

func main() {
	file, err := os.OpenFile("./finance.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var userCommand string

	BuildSliceLinesFromFile("./finance.txt")
	//check empty list finance.txt
	if len(lines) <= 1 {
		lines = CreatEmptyList(lines)
	}
	//update all days

	/*********************************/

	for {
		CheckTodayDays()
		CheckBalance()
		FindSaveMoney()
		FindManeyInOneDay()
		DisplayMainInfo()

		fmt.Println("\n" + "Введите команду или \"List\" для справки")
		fmt.Fscan(os.Stdin, &userCommand)
		if userCommand == "Exit" || userCommand == "exit" { //EXIT FROM FOR
			Write(*file, lines)
			WriteLog(lines)
			break
		}

		iter, ok := allMap[userCommand]
		if ok != false {
			SelectionUserCommand(iter)
		} else {
			fmt.Println("Нет такой команды(( наберите list")
		}
	}
}
