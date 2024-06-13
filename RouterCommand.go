package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RouterCommand struct {
	fileList []string
	usingMap map[string]fn
	account  Accountant
	data     int
	arg1     string
	arg2     string
	res      int
}

func (r *RouterCommand) InputCommand(command string) bool {
	exitFlag := false
	command = strings.TrimSpace(command)
	var commandSplit = []string{"", "", ""}
	commandSplit = strings.Split(command, " ")
	if len(commandSplit) == 3 {
		r.arg1 = commandSplit[1]
		r.arg2 = commandSplit[2]
	}
	if len(commandSplit) == 2 {
		r.arg1 = commandSplit[1]
		r.arg2 = ""
	}
	if commandSplit[0] != "Exit" {
		if commandSplit[0] != "Help" {
			_, ok := r.usingMap[commandSplit[0]]
			if ok {
				r.usingMap[commandSplit[0]](r)
				r.CleanArgument()
			} else {
				fmt.Println("Такой команды нет( набирите Help")
			}
		} else {
			HelpList()
		}
	} else {
		exitFlag = true
		if r.account.name != "" {
			fmt.Println("Сохранить файл???  Y/N")
			word := ""
			fmt.Fscan(os.Stdin, &word)
			if word == "Y" || word == "Yes" {
				r.account.WriteDataInFile()
			}
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
	fmt.Print("->")
}

func Select(r *RouterCommand) {
	reg, _ := regexp.MatchString("\\.txt", r.arg1)
	if reg {
		r.arg1 = r.arg1[:len(r.arg1)-4]
	}
	r.account = MakeAccountant(r.arg1, time.Now().Year())
	errorReadFile := r.account.ReadInfoFromFile()
	if errorReadFile != nil {
		fmt.Println("Нет такого файла! Хотите его создать? Y/N")
		word := ""
		fmt.Fscan(os.Stdin, &word)
		if word == "Y" || word == "Yes" {
			nameNewFile := New(r.account.name)
			_ = r.account.ReadInfoFromFile()
			r.AppendNameFileToSlice(nameNewFile)
		}
	}
}

func NewFile(r *RouterCommand) {
	nameNewFile := New(r.arg1)
	r.AppendNameFileToSlice(nameNewFile)
	DisplayFile(r.fileList)
}

func Base(r *RouterCommand) {
	if r.arg1 == "" {
		fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.startBase))
		fmt.Println("Введите стартовый баланс")
		fmt.Fscan(os.Stdin, &r.arg1)
	}
	base, _ := strconv.Atoi(r.arg1)
	r.account.SetStartBase(base)
}

func Minus(r *RouterCommand) {
	if r.arg1 == "" {
		fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.expenses))
		fmt.Println("Введите текущие расходы")
		fmt.Fscan(os.Stdin, &r.arg1)
	}
	minus, _ := strconv.Atoi(r.arg1)
	r.account.SetExpenses(minus)
	r.account.FoundCurrentBase()
	r.account.FoundReceipts()
}

func Plus(r *RouterCommand) {
	if r.arg1 == "" {
		fmt.Println("Введите новые поступление")
		fmt.Fscan(os.Stdin, &r.arg1)
	}
	plus, _ := strconv.Atoi(r.arg1)
	r.account.SetNewPlus(plus)
	r.account.FoundCurrentBase()
	r.account.FoundReceipts()
}

func Salary(r *RouterCommand) {
	if r.arg1 == "" {
		fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.salary))
		fmt.Println("Введите ожидаемый доход")
		fmt.Fscan(os.Stdin, &r.arg1)
	}
	salary, _ := strconv.Atoi(r.arg1)
	r.account.SetSalary(salary)
	r.account.FoundCurrentBase()
	r.account.FoundReceipts()
}

func Target(r *RouterCommand) {
	if r.arg1 == "" {
		fmt.Println("Предыдущее значение: " + ConvectIntToStr(r.account.target))
		fmt.Println("Введите какую сумму хотим достигнуть")
		fmt.Fscan(os.Stdin, &r.arg1)
	}
	target, _ := strconv.Atoi(r.arg1)
	r.account.SetTarget(target)
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
	if r.arg1 == "" {
		fmt.Println("Предыдущие имя файла: " + r.account.name)
		fmt.Println("Введите новое имя для файла")
		fmt.Fscan(os.Stdin, &r.arg1)
	}
	r.account.SetNewName(r.arg1)
}

func Save(r *RouterCommand) {
	r.account.WriteDataInFile()
}

func Delete(r *RouterCommand) {
	if r.account.name != "" {
		fmt.Println("Сначала завершите работу с текущим файлом и выйдите из него")
	} else {
		if r.arg1 == "" {
			fmt.Println("Какой файл вы хотите удалить?")
			fmt.Fscan(os.Stdin, &r.arg1)
		}
		DeleteFile(r.arg1)
		r.RebuildFileList()
	}
}
func s(r *RouterCommand) {

}

func Cancel(r *RouterCommand) {
	r.account.name = ""
	fmt.Println("Текущий файл закрыт")
}

func (r *RouterCommand) CleanArgument() {
	r.arg1 = ""
	r.arg2 = ""
}

func (r *RouterCommand) RebuildFileList() {
	newList := []string{}
	for _, v := range r.fileList {
		if v != r.arg1+".txt" {
			newList = append(newList, v)
		}
	}
	r.fileList = newList
}
