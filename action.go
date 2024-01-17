package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectionUserCommand(num int) {
	if num < 4 {
		fmt.Println(message[num])
		WriteNewDataInLines(num)
	}

	iter, ok := serviceMap[num]
	if ok != false {
		iter()
	}
}

// add user data in slice
func WriteNewDataInLines(num int) {
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
		for i, v := range lines {
			res = strings.Contains(v, teg[num])
			if res == true {
				targetLine = i
			}
		}
		lines[targetLine] = teg[num] + " -> " + str
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

func SetNewTargetDay() {
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
}

func IncrementExpenses() {
	var input string
	fmt.Println(message[4])
	_, err := fmt.Fscan(os.Stdin, &input)
	if err != nil {
		return
	}

	if _, err := strconv.Atoi(input); err == nil {
		ActionAdd(lines, "expenses", input)
	} else {
		fmt.Println("!!Введено не число!!\n")
	}
}

func IncrementBase() {
	var input string
	fmt.Println(message[3])
	_, err := fmt.Fscan(os.Stdin, &input)
	if err != nil {
		return
	}
	if newPlus, err := strconv.Atoi(input); err == nil {
		base := GetDataLine("base")
		base = base + newPlus
		strBase := strconv.Itoa(base)
		ActionAdd(lines, "base", strBase)
	} else {
		fmt.Println("!!Введено не число!!\n")
	}
}

func DisplayMainInfo() {
	myBal := GetStringDataLineFromSlice("balance")
	myTar := GetStringDataLineFromSlice("target")
	myMoney := GetStringDataLineFromSlice("saveMoney")
	myMID := GetStringDataLineFromSlice("moneyInDay")
	myDay := GetStringDataLineFromSlice("daysLeft")
	fmt.Println("Текущий баланс: " + myBal + "  Цель достич: " + myTar + "\n" + "Доступно: " + myMoney + " Доступно в день: " + myMID + " Осталось дней: " + myDay)
}

func ShowListNameAllFunction() {
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

func ShowInfoNextDay() {
	day := GetDataLine("daysLeft")
	save := GetDataLine("saveMoney")
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
