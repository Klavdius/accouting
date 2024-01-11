package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func WriteLog(lines []string) {
	logFile, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()

	oBal, oldMid, err := LastParam()
	if err != nil {
		fmt.Println(err)
	}

	save := GetStringDataLine(lines, "saveMoney")
	expenses := GetStringDataLine(lines, "expenses")
	MID := GetDataLine(lines, "moneyInDay")
	strMID := GetStringDataLine(lines, "moneyInDay")
	day := time.Now()

	analyses := " Хорошо"
	if oldMid > MID {
		analyses = " Превышение!"
	}
	strDay := day.Format("2006-01-02 15:04:05")
	record := strDay + " Осталось: " + save + " текущие расходы: " + expenses + " денег в день: " + strMID + analyses

	if oBal != save {
		WriteInLog(*logFile, record)
	}
}

func WriteInLog(file os.File, str string) {
	_, err := file.WriteString(str + "\n")
	if err != nil {
		return
	}
}

// get latest data form log file for analisys
func LastParam() (string, int, error) {
	fileData, err := os.ReadFile("./log.txt")
	if err != nil {
		fmt.Println("Ошибка при чтение log.txt")
		fmt.Println(err)
	}

	strFileData := string(fileData)
	stringData := strings.Split(strFileData, "\n")
	lengSplit := len(stringData)
	if lengSplit != 1 {
		lastString := stringData[lengSplit-2]
		sliceLastString := strings.Split(lastString, ": ")
		bal1 := strings.Split(sliceLastString[1], " ")
		oldBal := bal1[0]
		m1 := strings.Split(sliceLastString[3], " ")
		m := m1[0]
		intM, _ := strconv.Atoi(m)
		return oldBal, intM, nil
	} else {
		return "_", 0, err
	}
}
