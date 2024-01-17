package main

var allMap = map[string]int{
	"Base":    0,
	"Balance": 1,
	"Target":  2,
	"Salary":  3,
	"Plus":    4,
	"Minus":   5,
	"Day":     6,
	"List":    7,
	"NextDay": 8,
	"Exit":    9,
}

var listHelp = map[int]string{
	0: "Введите начальную сумму",
	1: "Введите ваш текущий баланс",
	2: "Введите какую сумму вы хотите",
	3: "Введите какая сумма должна поступить в период",
	4: "Введите пополнения",
	5: "Введите текущие расходы",
	6: "Введите до какого дня вы копите",
	7: "Помощь с программой",
	8: "Доступные расходы на следующий день",
	9: "Выйти из программы",
}
var serviceMap = map[int]func(){
	4: IncrementBase,
	5: IncrementExpenses,
	6: SetNewTargetDay,
	7: ShowListNameAllFunction,
	8: ShowInfoNextDay,
}
