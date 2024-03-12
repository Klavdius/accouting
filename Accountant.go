package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Accountant struct {
	name        string
	year        int
	startBase   int
	currentBase int
	expenses    int
	salary      int
	receipts    int
	target      int
	beforeDay   string
}

func (a Accountant) MoneyOnDay() int {
	nextUnixMonth, _ := time.Parse("2006-Jan-02", a.beforeDay)
	dayLeft := nextUnixMonth.Unix()
	currentDay := time.Now().Unix()
	remains := dayLeft - currentDay
	needDay := float64(remains / 86400)
	needDay = math.Round(needDay)
	leftDays := int(needDay)
	fmt.Print("Осталось дней: " + ConvectIntToStr(leftDays) + " ")
	saveMoney := a.salary - (a.target - a.startBase) - a.expenses
	return saveMoney / leftDays
}

func (a *Accountant) ReadInfoFromFile() {
	fileName := a.name + ".txt"

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	stringFileData := string(fileData)
	stringData := strings.Split(stringFileData, "\n")
	a.year = ConvectStrToInt(stringData[1])
	a.startBase = ConvectStrToInt(stringData[2])
	a.currentBase = ConvectStrToInt(stringData[3])
	a.expenses = ConvectStrToInt(stringData[4])
	a.salary = ConvectStrToInt(stringData[5])
	a.receipts = ConvectStrToInt(stringData[6])
	a.target = ConvectStrToInt(stringData[7])
	a.beforeDay = stringData[8]

}

func (a Accountant) WriteDataInFile() {
	file, err := os.OpenFile("./"+a.name+".txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err.Error() + " Ошибка записи в файл")
	}
	defer file.Close()
	data := CreatData(a)
	for i, v := range data {
		if i < (len(data) - 1) {
			file.WriteString(v + "\n")
		} else {
			file.WriteString(v)
		}
	}
}

func (a *Accountant) SetStartBase(data int) {
	a.startBase = data
}

func (a *Accountant) SetExpenses(data int) {
	a.expenses = data
}

func (a *Accountant) SetSalary(data int) {
	a.salary = data
}

func (a *Accountant) FoundCurrentBase() {
	a.currentBase = a.startBase - a.expenses
}

func (a *Accountant) SetNewPlus(data int) {
	a.startBase = a.startBase + data
}

func (a *Accountant) FoundReceipts() {
	a.receipts = a.salary - a.expenses
}

func (a *Accountant) SetBeforeDay(data string) {
	a.beforeDay = data
}

func (a *Accountant) SetTarget(data int) {
	a.target = data
}

func (a *Accountant) SetNewName(data string) {
	a.name = data
}
