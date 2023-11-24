package main

import (
	"fmt"
	"os"
	"strings"
)

var message = []string{"Введите новый баланс", "Введите новую цель", "Введите новое поступление", "Введите расходы", "Введите текущий день", "До какого дня считаем"}
var teg = []string{"balance", "salary", "expenses", "currentDay", "target", "beforeDay"}
var fullTeg = []string{"balance", "target", "salary", "expenses", "currentDay", "beforeDay", "daysLeft", "moneyInDay"}

func Action(newLines []string, num int) {
	if num < 6 {
		ActionNew(newLines, num)
	}

	iter, ok := serviceMap[num]
	if ok != false {
		iter()
	}

}

func ActionNew(newLines []string, num int) {
	fmt.Println(message[num])
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

func ShowList() {
	for key, _ := range allMap {
		fmt.Println(key)
	}
}

func CheckDays(newLines []string) []string {
	var curDay, befDay int
	var res1, res2 bool
	for i, v := range newLines {
		res1 = strings.Contains(v, "currentDay")
		res2 = strings.Contains(v, "beforeDay")
		if res1 == true {
			curDay = i
		}
		if res2 == true {
			befDay = i
		}
	}
	fmt.Println(string(curDay) + " " + string(befDay))
	//curDay = newLines[curDay]
	//befDay = newLines[befDay]
	return newLines
}

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

func CreatEmptyList(slice []string) []string {
	for _, teg := range fullTeg {
		slice = append(slice, teg+" -> 0")
	}
	ActionAdd(slice, "currentDay", "1")

	return slice
}
