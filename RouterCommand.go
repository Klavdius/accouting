package main

import (
	"fmt"
	"os"
	"time"
)

type RouterCommand struct {
	fileList []string
	usingMap map[string]string
	account  Accountant
}

func (r *RouterCommand) InputCommand(command string) bool {
	exitFlag := false
	if command != "Exit" {
		v, ok := r.usingMap[command]
		if ok {
			switch v {
			case "New":
				nameNewFile := New()
				r.AppendNameFileToSlice(nameNewFile)
				DisplayFile(r.fileList)

			case "Select":
				var newCommand string
				fmt.Println("Введите название файла")
				fmt.Fscan(os.Stdin, &newCommand)
				r.account = MakeAccountant(newCommand, time.Now().Year())
				r.account.ReadInfoFromFile()
			case "s":
				r.account.MoneyOnDay()
			}
		}
	} else {
		exitFlag = true
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
	fmt.Println("Текущий баланс -- " + ConvectIntToStr(r.account.currentBase))
	fmt.Println("Доступно в день -- ")
}
