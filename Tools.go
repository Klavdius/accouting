package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var message = []string{"Введите новый баланс", "Введите новую цель", "Введите новое поступление", "Введите расходы", "До какого дня считаем в формате год/месяц/день (числами)"}
var teg = []string{"balance", "target", "salary", "expenses", "beforeDay", "currentDay"}
var fullTeg = []string{"balance", "target", "salary", "expenses", "beforeDay", "currentDay", "daysLeft", "moneyInDay"}
var dayInMonth = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
var tm = time.Unix(0, 0)
var epoch = tm.Unix()

func ShowList() {
	for key, _ := range allMap {
		fmt.Println(key)
	}
}

//func CheckDays(newLines []string) []string {
//
//}

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

func GetTimeStamp() string {
	tm := time.Now().Unix()
	int_tm := int(tm)
	str_data := strconv.Itoa(int_tm)
	return str_data
}
