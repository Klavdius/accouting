package main

type fn func(*RouterCommand)

var function = map[string]fn{
	"New":       NewFile,
	"Select":    Select,
	"s":         s,
	"Minus":     Minus,
	"Plus":      Plus,
	"Base":      Base,
	"Salary":    Salary,
	"Target":    Target,
	"BeforeDay": BeforeDay,
	"Save":      Save,
	"Rename":    Rename,
	"Cancel":    Cancel,
}

var helpMap = map[string]string{
	"New":       "Создание нового файла на текущий месяц",
	"Select":    "Выбрать лог нужного месяца",
	"s":         "s",
	"Minus":     "Указать какие расходы в месяце",
	"Plus":      "Прибавление поступлений",
	"Base":      "Начальная сумма на первое число месяца",
	"Salary":    "Какие поступления ожидаются в текущем месяце",
	"Target":    "Сумма, которую мы хотим получить в конце месяца",
	"BeforeDay": "До какого дня мы расчитываем бюджет",
	"Save":      "Сохранить расчеты в файл",
	"Rename":    "Переименовать файл",
	"Cancel":    "Выйти из текущего лога",
}
