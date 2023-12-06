package main

var allMap = map[string]int{
	"Balance": 0,
	"Target":  1,
	"Plus":    2,
	"Minus":   3,
	"Day":     4,
	"list":    5,
}

var listHelp = map[int]string{
	0: "Введите ваш текущий баланс",
	1: "Введите какую сумму вы хотите",
	2: "Введите какая сумма должна поступить в период",
	3: "Введите текущие расходы",
	4: "Введите до какого дня вы копите",
	5: "Помощь с программой",
}
var serviceMap = map[int]func(){
	5: ShowList,
}
