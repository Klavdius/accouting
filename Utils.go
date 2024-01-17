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

func CheckDayUntil() {
	data := GetStringDataFromFieldSlice("beforeDay")
	if data != "0" {
		tm, _ := time.Parse(layout, data)
		dayLeft := tm.Unix()
		day := time.Now().Unix()
		remains := dayLeft - day
		needDay := float64(remains / 86400)
		needDay = math.Round(needDay)
		leftDays := int(needDay)
		str_dayLeft := strconv.Itoa(leftDays)
		AddDataInLineSlice("daysLeft", str_dayLeft)
	}

}

func FindManeyInOneDay() {
	days := GetIntDataFromFieldSlice("daysLeft")
	meSave := GetIntDataFromFieldSlice("saveMoney")
	if days != 0 {
		ManeyInOneDay := meSave / days
		strMID := strconv.Itoa(ManeyInOneDay)
		AddDataInLineSlice("moneyInDay", strMID)
	} else {
		AddDataInLineSlice("moneyInDay", "0")
	}

}

func FindSaveMoney() {
	meSalary := GetIntDataFromFieldSlice("salary")
	meExp := GetIntDataFromFieldSlice("expenses")
	meTarget := GetIntDataFromFieldSlice("target")
	meBase := GetIntDataFromFieldSlice("base")
	meSave := (meBase + meSalary) - meTarget - meExp
	strSave := strconv.Itoa(meSave)
	AddDataInLineSlice("saveMoney", strSave)
}

// get data fron target slice
func GetIntDataFromFieldSlice(nameFieldInSlice string) int {
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

func GetStringDataFromFieldSlice(nameFieldInSlice string) string {
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
	return myStr
}

func ChengeBalance() {
	var (
		base, exp, balance int
	)
	base = GetIntDataFromFieldSlice("base")
	exp = GetIntDataFromFieldSlice("expenses")
	balance = base - exp
	strBalance := strconv.Itoa(balance)
	AddDataInLineSlice("balance", strBalance)
}

// custom write date in txt file
func WriteDateInFile(file os.File, list []string) {
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
	CheckEmptyFile := 2
	if len(stringData) < CheckEmptyFile {
		stringData = nil
	}
	for _, word := range stringData {
		lines = append(lines, word)
	}
}
