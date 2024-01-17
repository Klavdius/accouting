package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var CurrentDay = time.Now()
var stringCurrentDay = CurrentDay.Format("2006-01-02 15:04:05")

func WriteLog(lines []string) {
	logFile, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()

	save := GetStringDataFromFieldSlice("saveMoney")
	expenses := GetStringDataFromFieldSlice("expenses")
	strMID := GetStringDataFromFieldSlice("moneyInDay")

	analyses, err := GetAnalysesLog(strMID)
	if err != nil {
		fmt.Println(err)
	}

	record := stringCurrentDay + " Осталось: " + save + " текущие расходы: " + expenses + " денег в день: " + strMID + " " + analyses

	oldSaveMoney := GetSaveMoneyLastLine()
	ChangeDay := CheckChangeDay()
	if oldSaveMoney != save || ChangeDay == "Change" {
		WriteInLog(*logFile, record)
	}
}

func WriteInLog(file os.File, str string) {
	_, err := file.WriteString(str + "\n")
	if err != nil {
		return
	}
}

func GetAnalysesLog(MID string) (result string, err error) {
	lengSplit := GetLengLogFile()
	allowance, _ := strconv.Atoi(MID)
	if lengSplit != 1 {
		lastLine := GetSlacesLastLineFormLogFile()

		lastSectorToday := strings.Split(lastLine[3], " ")
		stringAllowanceYesterday := lastSectorToday[0]
		allowanceYesterday, _ := strconv.Atoi(stringAllowanceYesterday)
		if allowance >= allowanceYesterday {
			return "Good", nil
		} else {
			return "Bad", nil
		}

	} else {
		return "_", err
	}
}

func GetSlacesLastLineFormLogFile() []string {
	var numberLine = 2
	sliceLastDayInLog := GetSliceLineFromLog(numberLine)

	return sliceLastDayInLog
}

func GetSliceLineFromLog(numberLine int) []string {
	lengSplit := GetLengLogFile()
	allData := GetAllStringSliceDataLogFile()
	lastLineInLog := allData[lengSplit-numberLine]
	sliceString := strings.Split(lastLineInLog, ": ")

	return sliceString
}

func GetLengLogFile() int {
	SliceData := GetAllStringSliceDataLogFile()
	lengSplit := len(SliceData)

	return lengSplit
}

func GetAllStringSliceDataLogFile() []string {
	fileData, err := os.ReadFile("./log.txt")
	if err != nil {
		fmt.Println("Ошибка при чтение log.txt")
		fmt.Println(err)
	}

	strFileData := string(fileData)
	stringData := strings.Split(strFileData, "\n")

	return stringData
}

func GetSaveMoneyLastLine() string {
	numberLine := 2
	sliceLine := GetSliceLineFromLog(numberLine)
	sectorLine := strings.Split(sliceLine[1], " ")
	saveMoney := sectorLine[0]

	return saveMoney
}

func CheckChangeDay() string {
	numberLine := 2
	var result string
	newSlice := GetSliceLineFromLog(numberLine)
	sliceLog := strings.Split(newSlice[0], " ")
	dayInLog := sliceLog[0]
	sliceToday := strings.Split(stringCurrentDay, " ")
	dayToday := sliceToday[0]
	if dayInLog == dayToday {
		result = "No change"
	} else {
		result = "Change"
	}

	return result
}
