package main

var allMap = map[string]int{
	"newBalance":   0,
	"newTarget":    1,
	"Plus":         2,
	"Minus":        3,
	"newBeforeDay": 4,
	"list":         5,
	"nbd":          4,
}

var serviceMap = map[int]func(){
	5: ShowList,
}

var mapMonth = map[string]string{
	"January":   "1",
	"February":  "2",
	"March":     "3",
	"April":     "4",
	"May":       "5",
	"June":      "6",
	"July":      "7",
	"August":    "8",
	"September": "9",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}
