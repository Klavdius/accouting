package main

import (
	"fmt"
	"strconv"
	"strings"
)

var message = []string{"Введите новый баланс", "Введите новую цель", "Введите новое поступление", "Введите расходы", "Введите текущий день в формате день.месяц (числами) ", "До какого дня считаем в формате день.месяц (числами)"}
var teg = []string{"balance", "target", "salary", "expenses", "currentDay", "beforeDay"}
var fullTeg = []string{"balance", "target", "salary", "expenses", "currentDay", "beforeDay", "daysLeft", "moneyInDay"}
var dayInMonth = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}

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
			curDay = i //number line current day in slice
		}
		if res2 == true {
			befDay = i
		}
	}
	cDay := newLines[curDay]
	bDay := newLines[befDay]
	if bDay != "beforeDay -> 0" {
		int_cDay := FullDayFromData(cDay)
		int_bDay := FullDayFromData(bDay)
		daysLeft := int_bDay - int_cDay
		ActionAdd(newLines, "daysLeft", strconv.Itoa(daysLeft))

	}

	return newLines
}

func FullDayFromData(data string) int {
	stringData := strings.Split(data, ".")
	firNum := strings.Split(stringData[0], "-> ")
	intDay, err := strconv.Atoi(stringData[1])
	if err != nil {
		fmt.Println("err in conver str -> num")
	}
	res := dayInMonth[intDay-1]
	firDay, _ := strconv.Atoi(firNum[1])

	return res + firDay
}
