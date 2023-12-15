package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("./finance.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var command string
	var lines = []string{}

	//custom system read from file
	fileData, err := os.ReadFile("./finance.txt")
	if err != nil {
		fmt.Println(err)
	}
	stringFileData := string(fileData)
	stringData := strings.Split(stringFileData, "\n")
	if len(stringData) < 4 {
		stringData = nil
	}
	for _, word := range stringData {
		lines = append(lines, word)
	}

	//***************************************************************

	//check empty list finance.txt
	if len(lines) <= 1 {
		lines = CreatEmptyList(lines)
	}

	//update all days
	CheckDays(lines)
	/*********************************/

	for {
		Display(lines)
		fmt.Println("\n" + "Введите команду или \"List\" для справки")
		fmt.Fscan(os.Stdin, &command)
		if command == "Exit" || command == "exit" { //EXIT FROM FOR
			fmt.Println("exit from program")
			Write(*file, lines)
			WriteLog(lines)
			break
		}

		iter, ok := allMap[command]
		if ok != false {
			MainAction(lines, iter)
		} else {
			fmt.Println("Нет такой команды(( наберите list")
		}
	}
}
