package main

import "fmt"

var function = map[string]func(){
	"New": New,
}

func RouterCommand(command string) bool {
	exitFlag := false
	if command != "Exit" {
		v, ok := function[command]
		if ok {
			v()
		}
	} else {
		exitFlag = true
	}
	return exitFlag
}

func New() {
	fmt.Println("input New")
}
