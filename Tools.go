package main

import (
	"fmt"
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

func CheckTodayDays() {
	data := GetStringDataLine("beforeDay")
	if data != "0" {
		tm, _ := time.Parse(layout, data)
		dayLeft := tm.Unix()
		day := time.Now().Unix()
		remains := dayLeft - day
		needDay := float64(remains / 86400)
		needDay = math.Round(needDay)
		leftDays := int(needDay)
		str_dayLeft := strconv.Itoa(leftDays)
		ActionAdd(lines, "daysLeft", str_dayLeft)
	}

}

// find money in day
func FindManeyInOneDay() {
	days := GetDataLine("daysLeft")
	meSave := GetDataLine("saveMoney")
	ManeyInOneDay := meSave / days
	strMID := strconv.Itoa(ManeyInOneDay)
	ActionAdd(lines, "moneyInDay", strMID)
}

func FindSaveMoney() {
	meSalary := GetDataLine("salary")
	meExp := GetDataLine("expenses")
	meTarget := GetDataLine("target")
	meBase := GetDataLine("base")
	meSave := (meBase + meSalary) - meTarget - meExp
	strSave := strconv.Itoa(meSave)
	ActionAdd(lines, "saveMoney", strSave)
}

// get data fron target slice
func GetDataLine(nameFieldInSlice string) int {
	var (
		res        bool
		targetLine int
	)
	for i, v := range lines {
		res = strings.Contains(v, nameFieldInSlice)
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
func GetStringDataLine(nameLineInSlice string) string {
	var (
		res        bool
		targetLine int
	)
	for i, v := range lines {
		res = strings.Contains(v, nameLineInSlice)
		if res == true {
			targetLine = i
		}
	}
	dateslice := strings.Split(lines[targetLine], "-> ")
	myStr := dateslice[1]
	return myStr
}

// Find current Balance with current date in slise
func CheckBalance() {
	var (
		base, exp, balance int
	)
	base = GetDataLine("base")
	exp = GetDataLine("expenses")
	balance = base - exp
	strBalance := strconv.Itoa(balance)
	ActionAdd(lines, "balance", strBalance)
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

func BuildSliceLinesFromFile(fileName string) {
	fileData, err := os.ReadFile(fileName)
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
}
