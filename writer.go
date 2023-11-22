package main

import (
	"fmt"
	"os"
)

func Write(file os.File, list []string) {
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	//err = file.Truncate(0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	for _, v := range list {
		file.WriteString(v + "\n")
	}
}
