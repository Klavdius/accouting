package main

import (
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var message = []string{"Введите начальную сумму", "Введите новый баланс", "Введите новую цель", "Введите новое поступление", "Введите расходы", "До какого дня считаем в формате год-месяц-день (2023-Dec-25)"}
var teg = []string{"base", "balance", "target", "salary", "expenses", "beforeDay"}
var fullTeg = []string{"base", "balance", "target", "salary", "expenses", "beforeDay", "daysLeft", "moneyInDay", "saveMoney"}
var month = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

const layout = "2006-Jan-02"

func CheckDays(newLine []string) []string {
	data := GetStringDataLine(newLine, "beforeDay")
	if data != "0" {
		tm, _ := time.Parse(layout, data)
		dayLeft := tm.Unix()
		day := time.Now().Unix()
		remains := dayLeft - day
		needDay := float64(remains / 86400)
		needDay = math.Round(needDay)
		leftDays := int(needDay)
		FindMID(newLine, leftDays)
		str_dayLeft := strconv.Itoa(leftDays)
		ActionAdd(newLine, "daysLeft", str_dayLeft)
	}
	return newLine
}

// find money in day
func FindMID(lines []string, days int) {
	meSalary := GetDataLine(lines, "salary")
	meExp := GetDataLine(lines, "expenses")
	meTarget := GetDataLine(lines, "target")
	meBase := GetDataLine(lines, "base")
	meSave := (meBase + meSalary) - meTarget - meExp
	MID := meSave / days
	strMID := strconv.Itoa(MID)
	strSave := strconv.Itoa(meSave)
	ActionAdd(lines, "saveMoney", strSave)
	ActionAdd(lines, "moneyInDay", strMID)
}

func GetTimeStamp() string {
	tm := time.Now().Unix()
	int_tm := int(tm)
	str_data := strconv.Itoa(int_tm)
	return str_data
}

// get data fron target slice
func GetDataLine(lines []string, line string) int {
	var (
		res        bool
		targetLine int
	)
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

// get data from target slice in string format, for show
func GetStringDataLine(lines []string, line string) string {
	var (
		res        bool
		targetLine int
	)
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

// Find current Balance with current date in slise
func FindIntBalance(lines []string) int {
	var (
		base, exp, balance int
	)
	base = GetDataLine(lines, "base")
	exp = GetDataLine(lines, "expenses")
	balance = base - exp
	return balance
}

// custom write date in txt file
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
