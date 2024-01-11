package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MainAction(newLines []string, num int) {
	if num < 4 {
		ActionNew(newLines, num)
	}

	iter, ok := serviceMap[num]
	if ok != false {
		iter(newLines)
	}
}

// add user data in slice
func ActionNew(newLines []string, num int) {
	fmt.Println(message[num])

	var (
		str        string
		targetLine int
		res        bool
	)
	_, err := fmt.Fscan(os.Stdin, &str)
	if err != nil {
		return
	}
	if _, err := strconv.Atoi(str); err == nil {
		for i, v := range newLines {
			res = strings.Contains(v, teg[num])
			if res == true {
				targetLine = i
			}
		}
		newLines[targetLine] = teg[num] + " -> " + str
		newLines = CheckDays(newLines)

		balance := FindIntBalance(newLines)
		strBalance := strconv.Itoa(balance)
		ActionAdd(newLines, "balance", strBalance)

	} else {
		fmt.Println("!!Введено не число!!\n")
	}

}

// add data in target slice
func ActionAdd(newLines []string, line string, data string) {
	var (
		res        bool
		targetLine int
	)
	for i, v := range newLines {
		res = strings.Contains(v, line)
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = line + " -> " + data
}

func SetDay(lines []string) {
	fmt.Println(message[5])
	for _, v := range month {
		fmt.Print(v + "  ")
	}
	fmt.Println()
	var (
		str        string
		targetLine int
		res        bool
	)
	_, err := fmt.Fscan(os.Stdin, &str)
	if err != nil {
		return
	}
	for i, v := range lines {
		res = strings.Contains(v, teg[5])
		if res == true {
			targetLine = i
		}
	}
	lines[targetLine] = teg[5] + " -> " + str
	lines = CheckDays(lines)
}

func Minus(lines []string) {
	var input, strBalance string
	fmt.Println(message[4])
	_, err := fmt.Fscan(os.Stdin, &input)
	if err != nil {
		return
	}

	if _, err := strconv.Atoi(input); err == nil {
		ActionAdd(lines, "expenses", input)
		balance := FindIntBalance(lines)
		strBalance = strconv.Itoa(balance)
		ActionAdd(lines, "balance", strBalance)
		CheckDays(lines)
	} else {
		fmt.Println("!!Введено не число!!\n")
	}
}

func Plus(lines []string) {
	var input string
	fmt.Println(message[3])
	_, err := fmt.Fscan(os.Stdin, &input)
	if err != nil {
		return
	}
	if newPlus, err := strconv.Atoi(input); err == nil {
		base := GetDataLine(lines, "base")
		base = base + newPlus
		strBase := strconv.Itoa(base)
		ActionAdd(lines, "base", strBase)
		CheckDays(lines)
	} else {
		fmt.Println("!!Введено не число!!\n")
	}
}

// show base info
func Display(lines []string) {
	myBal := GetStringDataLine(lines, "balance")
	myTar := GetStringDataLine(lines, "target")
	myMoney := GetStringDataLine(lines, "saveMoney")
	myMID := GetStringDataLine(lines, "moneyInDay")
	myDay := GetStringDataLine(lines, "daysLeft")
	fmt.Println("Текущий баланс: " + myBal + "  Цель достич: " + myTar + "\n" + "Доступно: " + myMoney + " Доступно в день: " + myMID + " Осталось дней: " + myDay)
}

// show all function
func ShowList(lines []string) {
	for key, value := range allMap {
		fmt.Print(key)
		leng := len(key)
		for i := leng; i < 15; i++ {
			fmt.Print(".")
		}
		fmt.Println(listHelp[value])
	}
	fmt.Println()
}

func NextDay(lines []string) {
	day := GetDataLine(lines, "daysLeft")
	save := GetDataLine(lines, "saveMoney")
	nexMID := save / (day - 1)
	strNMID := strconv.Itoa(nexMID)
	fmt.Println("На следующий день доступно: " + strNMID + "\n")
}
func CreatEmptyList(slice []string) []string {
	for _, teg := range fullTeg {
		slice = append(slice, teg+" -> 0")
	}
	return slice
}

func Exit(lines []string) {

}
