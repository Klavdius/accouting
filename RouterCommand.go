package main

import (
	"fmt"
	"os"
	"time"
)

type RouterCommand struct {
	fileList []string
	usingMap map[string]fn
	account  Accountant
	data     int
}

func (r *RouterCommand) InputCommand(command string) bool {
	exitFlag := false

	if command != "Exit" {
		if command != "Help" {
			_, ok := r.usingMap[command]
			if ok {
				r.usingMap[command](r)
			} else {
				fmt.Println("Такой команды нет( набирите Help")
			}
		} else {
			HelpList()
		}
	} else {
		exitFlag = true
		if r.account.name != "" {
			r.account.WriteDataInFile()
		}
	}
	return exitFlag
}

func (r *RouterCommand) AppendNameFileToSlice(name string) {
	double := false
	for _, v := range r.fileList {
		if v == name {
			double = true
		}
	}
	if !double {
		r.fileList = append(r.fileList, name)
	}
}

func (r *RouterCommand) DisplayInfoAboutAccount() {
	if r.account.name != "" {
		fmt.Println("Текущий баланс -- " + ConvectIntToStr(r.account.currentBase) +
			" Цель достигнуть -> " + ConvectIntToStr(r.account.target))
		MoD := r.account.MoneyOnDay()
		fmt.Println("Доступно в день -- " + ConvectIntToStr(MoD))
	}
}

func Select(r *RouterCommand) {
	var newCommand string
	fmt.Println("Введите название файла")
	fmt.Fscan(os.Stdin, &newCommand)
	r.account = MakeAccountant(newCommand, time.Now().Year())
	r.account.ReadInfoFromFile()
}

func NewFile(r *RouterCommand) {
	nameNewFile := New()
	r.AppendNameFileToSlice(nameNewFile)
	DisplayFile(r.fileList)
}

func Base(r *RouterCommand) {

	fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.startBase))
	fmt.Println("Введите стартовый баланс")
	fmt.Fscan(os.Stdin, &r.data)
	r.account.SetStartBase(r.data)
}

func Minus(r *RouterCommand) {
	fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.expenses))
	fmt.Println("Введите текущие расходы")
	fmt.Fscan(os.Stdin, &r.data)
	r.account.SetExpenses(r.data)
	r.account.FoundCurrentBase()
	r.account.FoundReceipts()
}

func Plus(r *RouterCommand) {
	fmt.Println("Введите новые поступление")
	fmt.Fscan(os.Stdin, &r.data)
	r.account.SetNewPlus(r.data)
	r.account.FoundCurrentBase()
	r.account.FoundReceipts()
}

func Salary(r *RouterCommand) {
	fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.salary))
	fmt.Println("Введите ожидаемый доход")
	fmt.Fscan(os.Stdin, &r.data)
	r.account.SetSalary(r.data)
	r.account.FoundCurrentBase()
	r.account.FoundReceipts()
}

func Target(r *RouterCommand) {
	fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.target))
	fmt.Println("Введите какую сумму хотим достигнуть")
	fmt.Fscan(os.Stdin, &r.data)
	r.account.SetTarget(r.data)
}

func BeforeDay(r *RouterCommand) {
	fmt.Println("Предыдущее значение: " + r.account.beforeDay)
	fmt.Println("Введите до какой даты считаем в формате 2024-Jan-01")
	PrintMonthHelp()
	var date string
	fmt.Fscan(os.Stdin, &date)
	r.account.SetBeforeDay(date)
}

func Rename(r *RouterCommand) {
	fmt.Println("Предыдущие имя файла: " + r.account.name)
	fmt.Println("Введите новое имя для файла")
	var data string
	fmt.Fscan(os.Stdin, &data)
	r.account.SetNewName(data)
}

func Save(r *RouterCommand) {
	r.account.WriteDataInFile()
}

func s(r *RouterCommand) {

}

func Cancel(r *RouterCommand) {

}
