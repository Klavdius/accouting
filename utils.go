package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"time"
)

var month = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

func PrintMonthHelp() {
	for _, v := range month {
		fmt.Print(v + " ")
	}
	fmt.Println()
}
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
	fmt.Print("->")
}

func CheckingLengInfos(infos []string) {
	if len(infos) == 0 {
		fmt.Println("В директории нет файлов учёта 'Acc_' ")
	}
}

func CreatFile(a Accountant) (string, error) {
	file, err := os.OpenFile("./mount/"+a.name+".txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return "", err
	}
	file.WriteString(CreatAccountantToString(a))
	return a.name, err
}

func DeleteFile(nameFile string) {
	reg, _ := regexp.MatchString("\\.txt", nameFile)
	if reg {
		nameFile = nameFile[:len(nameFile)-4]
	}

	err := os.Remove("./mount/" + nameFile + ".txt")
	if err != nil {
		fmt.Println("Не удалось удалить файл")
	} else {
		fmt.Println("Файл успешно удалён")
	}
}

func CreatAccountantToString(a Accountant) string {
	var line string
	line = a.name + "\n" +
		ConvectIntToStr(a.year) + "\n" +
		ConvectIntToStr(a.startBase) + "\n" +
		ConvectIntToStr(a.currentBase) + "\n" +
		ConvectIntToStr(a.expenses) + "\n" +
		ConvectIntToStr(a.salary) + "\n" +
		ConvectIntToStr(a.receipts) + "\n" +
		ConvectIntToStr(a.target) + "\n" +
		a.beforeDay
	return line
}

func New(newName string) string {
	var a Accountant
	t := time.Now()
	if newName == "" {
		a.name = t.Month().String()
	} else {
		a.name = newName
	}
	a.year = t.Year()
	a.beforeDay = "2024-Jan-01"
	newName, err := CreatFile(a)
	if err != nil {
		fmt.Println("файл не создан!")
	}
	return newName + ".txt"
}

func CreatData(a Accountant) []string {
	var data = []string{a.name, ConvectIntToStr(a.year), ConvectIntToStr(a.startBase), ConvectIntToStr(a.currentBase),
		ConvectIntToStr(a.expenses), ConvectIntToStr(a.salary), ConvectIntToStr(a.receipts), ConvectIntToStr(a.target), a.beforeDay}
	return data
}

func ConvectIntToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

func ConvectStrToInt(data string) int {
	intData, _ := strconv.Atoi(data)
	return intData
}

func HelpList() {
	for i, v := range helpMap {
		fmt.Println(i + " -> " + v)
	}
	fmt.Println()
}

func CreatListAboutFileInDirect(dir []fs.DirEntry) []string {
	infos := make([]string, 0, len(dir))
	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			fmt.Println("error")
		}
		reg, _ := regexp.MatchString("\\.txt", info.Name())
		if reg {
			infos = append(infos, info.Name())
		}
	}
	return infos
}

func NeedListOfFiles(r RouterCommand) bool {
	if r.account.name == "" {
		return true
	}
	return false
}
