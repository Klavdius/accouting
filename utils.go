package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func MakeAccountant(nameInput string, yearInput int) Accountant {
	var a Accountant
	a.name = nameInput
	a.year = yearInput
	return a
}

func DisplayFile(infos []string) {
	for _, v := range infos {
		fmt.Print(v + " ")
	}
	fmt.Println()
}

func CheckingLengInfos(infos []string) {
	if len(infos) == 0 {
		fmt.Println("В директории нет файлов учёта 'Acc_' ")
	}
}

func CreatFile(a Accountant) (string, error) {
	file, err := os.OpenFile("Acc_"+a.name+"_"+strconv.Itoa(a.year)+".txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return "", err
	}

	file.WriteString(CreatAccountantToString(a))
	return "Acc_" + a.name + "_" + strconv.Itoa(a.year), err
}

func CreatAccountantToString(a Accountant) string {
	var line string
	line = a.name + "\n" +
		ConvectIntToStr(a.year) +
		ConvectIntToStr(a.startBase) +
		ConvectIntToStr(a.currentBase) +
		ConvectIntToStr(a.expenses) +
		ConvectIntToStr(a.salary) +
		ConvectIntToStr(a.receipts) +
		ConvectIntToStr(a.target) +
		a.beforeDay
	return line
}

func New() string {
	var a Accountant
	t := time.Now()
	a.name = t.Month().String()
	a.year = t.Year()
	a.beforeDay = "2024-Jan-01"
	name, err := CreatFile(a)
	if err != nil {
		fmt.Println("файл не создан!")
	}
	return name
}

func ConvectIntToStr(i int) string {
	s := strconv.Itoa(i) + "\n"
	return s
}

func ConvectStrToInt(data string) int {
	intData, _ := strconv.Atoi(data)
	return intData
}
