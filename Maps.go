package main

type fn func(*RouterCommand)

var function = map[string]fn{
	"New":       NewFile,
	"Select":    Select,
	"s":         s,
	"Minus":     Minus,
	"Plus":      Plus,
	"Base":      Base,
	"Salary":    Salary,
	"Target":    Target,
	"BeforeDay": BeforeDay,
	"Save":      Save,
	"Rename":    Rename,
	"Cancel":    Cancel,
}
