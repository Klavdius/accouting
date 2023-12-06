package main

import (
	"fmt"
	"os"
	"strings"
)

func MainAction(newLines []string, num int) {
	if num < 5 {
		ActionNew(newLines, num)
	}

	iter, ok := serviceMap[num]
	if ok != false {
		iter()
	}

}

// add user data in slice
func ActionNew(newLines []string, num int) {
	fmt.Println(message[num])
	if num == 4 {
		for _, v := range month {
			fmt.Print(v + "  ")
		}
		fmt.Println()
	}
	var str string
	var targetLine int
	var res bool
	fmt.Fscan(os.Stdin, &str)
	for i, v := range newLines {
		res = strings.Contains(v, teg[num])
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = teg[num] + " -> " + str

	newLines = CheckDays(newLines)

}

// add data in target slice
func ActionAdd(newLines []string, line string, data string) {
	var res bool
	var targetLine int
	for i, v := range newLines {
		res = strings.Contains(v, line)
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = line + " -> " + data
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
func ShowList() {
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

func CreatEmptyList(slice []string) []string {
	for _, teg := range fullTeg {
		slice = append(slice, teg+" -> 0")
	}

	return slice
}
