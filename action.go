package main

import (
	"fmt"
	"os"
	"strings"
)

func MainAction(newLines []string, num int) {
	if num < 6 {
		ActionNew(newLines, num)
	}

	iter, ok := serviceMap[num]
	if ok != false {
		iter()
	}

}

func ActionNew(newLines []string, num int) {
	fmt.Println(message[num])
	var str string
	var targetLine int
	var res bool
	fmt.Fscan(os.Stdin, &str)
	for i, v := range newLines {
		res = strings.Contains(v, teg[num])
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = teg[num] + " -> " + str

	newLines = CheckDays(newLines)
}

func ActionAdd(newLines []string, line string, data string) {
	var res bool
	var targetLine int
	for i, v := range newLines {
		res = strings.Contains(v, line)
		if res == true {
			targetLine = i
		}
	}
	newLines[targetLine] = line + " -> " + data
}

func CreatEmptyList(slice []string) []string {
	for _, teg := range fullTeg {
		slice = append(slice, teg+" -> 0")
	}
	ActionAdd(slice, "currentDay", "1")

	return slice
}
