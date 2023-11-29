package main

import (
	"os"
)

func Write(file os.File, list []string) {
	for i, v := range list {
		if i < 7 {
			file.WriteString(v + "\n")
		} else {
			file.WriteString(v)
		}
	}
}
