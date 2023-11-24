package main

var allMap = map[string]int{
	"newBalance":    0,
	"newTarget":     1,
	"newSalary":     2,
	"newExpenses":   3,
	"newCurrentDay": 4,
	"newBeforeDay":  5,
	"list":          6,
}

var serviceMap = map[int]func(){
	6: ShowList,
}
