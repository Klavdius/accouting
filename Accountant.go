package main

import (
	"fmt"
	"os"
	"strings"
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
}

func (a *Accountant) ReadInfoFromFile() {
	fileName := a.name + ".txt"

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	stringFileData := string(fileData)
	stringData := strings.Split(stringFileData, "\n")
	a.name = a.name
	a.year = ConvectStrToInt(stringData[1])
	a.startBase = ConvectStrToInt(stringData[2])
	a.currentBase = ConvectStrToInt(stringData[3])
	a.expenses = ConvectStrToInt(stringData[4])
	a.salary = ConvectStrToInt(stringData[5])
	a.receipts = ConvectStrToInt(stringData[6])
	a.target = ConvectStrToInt(stringData[7])
}
