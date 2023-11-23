package main

import (
	"fmt"
	"os"
	"strings"
)

var message = []string{"Введите новый баланс", "Введите новую цель", "Введите новое поступление", "Введите расходы", "Введите текущий день", "До какого дня считаем"}
var teg = []string{"balance", "target", "salary", "expenses", "currentDay", "beforeDay"}
var fullTeg = []string{"balance", "target", "salary", "expenses", "currentDay", "beforeDay", "moneyInDay"}

func Action(newLines []string, num int) {
	fmt.Println(message[num])
	str := ""
	targetLine := 0
	fmt.Fscan(os.Stdin, &str)
	for i, v := range newLines {
		var res bool
		res = strings.Contains(v, teg[num])
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = teg[num] + " -> " + str
}
