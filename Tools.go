package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var message = []string{"Введите новый баланс", "Введите новую цель", "Введите новое поступление", "Введите расходы", "До какого дня считаем в формате год-месяц-день (2023-Dec-25)"}
var teg = []string{"balance", "target", "salary", "expenses", "beforeDay", "currentDay"}
var fullTeg = []string{"balance", "target", "salary", "expenses", "beforeDay", "currentDay", "daysLeft", "moneyInDay", "saveMoney"}
var dayInMonth = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
var tm = time.Unix(0, 0)
var epoch = tm.Unix()

const layout = "2006-Jan-02"

func ShowList() {
	for key, _ := range allMap {
		fmt.Println(key)
	}
}

func CheckDays(newLine []string) []string {
	data := GetStringDataLine(newLine, "beforeDay")
	tm, _ := time.Parse(layout, data)
	dayLeft := tm.Day()
	day := time.Now().Day()
	dayLeft = dayLeft - day

	FindMID(newLine, dayLeft)
	str_dayLeft := strconv.Itoa(dayLeft)
	ActionAdd(newLine, "daysLeft", str_dayLeft)
	return newLine
}

func FindMID(lines []string, days int) {
	meSalary := GetDataLine(lines, "salary")
	meExp := GetDataLine(lines, "expenses")
	meTarget := GetDataLine(lines, "target")
	meBal := GetDataLine(lines, "balance")
	meSave := meTarget - meBal
	MID := (meSalary - meExp - meSave) / days
	strMID := strconv.Itoa(MID)
	ActionAdd(lines, "moneyInDay", strMID)
}
func GetTimeStamp() string {
	tm := time.Now().Unix()
	int_tm := int(tm)
	str_data := strconv.Itoa(int_tm)
	return str_data
}

func GetDataLine(lines []string, line string) int {
	var res bool
	var targetLine int
	for i, v := range lines {
		res = strings.Contains(v, line)
		if res == true {
			targetLine = i
		}
	}
	dateslice := strings.Split(lines[targetLine], "-> ")
	myStr := dateslice[1]
	myInt, _ := strconv.Atoi(myStr)
	return myInt
}

func GetStringDataLine(lines []string, line string) string {
	var res bool
	var targetLine int
	for i, v := range lines {
		res = strings.Contains(v, line)
		if res == true {
			targetLine = i
		}
	}
	dateslice := strings.Split(lines[targetLine], "-> ")
	myStr := dateslice[1]
	return myStr
}

func Write(file os.File, list []string) {
	leng := len(list)
	for i, v := range list {
		if i < (leng - 1) {
			file.WriteString(v + "\n")
		} else {
			file.WriteString(v)
		}
	}
}
