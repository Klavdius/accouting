package main

import (
	"os"
)

func Write(file os.File, list []string) {
	for _, v := range list {
		file.WriteString(v + "\n")
	}
}
