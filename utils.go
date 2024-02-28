package main

import (
	"fmt"
	"io/fs"
)

func MakeAccountant(nameInput string, yearInput int) Accountant {
	var a Accountant
	a.name = "Acc_" + nameInput
	a.year = yearInput
	return a
}

func DisplayFile(infos []fs.FileInfo) {
	for _, v := range infos {
		fmt.Print(v.Name() + " ")
	}
}

func CheckingLengInfos(infos []fs.FileInfo) {
	if len(infos) == 0 {
		fmt.Println("В директории нет файлов учёта 'Acc_' ")
	}
}
