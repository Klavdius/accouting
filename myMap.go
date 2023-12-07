package main

var allMap = map[string]int{
	"Balance": 0,
	"Target":  1,
	"Plus":    2,
	"Minus":   3,
	"Day":     4,
	"List":    5,
	"NextDay": 6,
}

var listHelp = map[int]string{
	0: "Введите ваш текущий баланс",
	1: "Введите какую сумму вы хотите",
	2: "Введите какая сумма должна поступить в период",
	3: "Введите текущие расходы",
	4: "Введите до какого дня вы копите",
	5: "Помощь с программой",
	6: "Доступные расходы на следующий день",
}
var serviceMap = map[int]func(slice []string){
	5: ShowList,
	6: NextDay,
}
