package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	entries, err := os.ReadDir("./mount")
	if err != nil {
		fmt.Println("error")
	}

	listFile := CreatListAboutFileInDirect(entries)
	CheckingLengInfos(listFile)

	var r RouterCommand
	r.fileList = listFile
	r.usingMap = function

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		command  string
		exitFlag bool
	)

	for {
		r.DisplayInfoAboutAccount()
		checkNeedList := NeedListOfFiles(r)
		if checkNeedList {
			DisplayFile(r.fileList)
		}
		command, _ = in.ReadString('\n')
		command = strings.TrimSpace(command)
		if command == "" {
			command, _ = in.ReadString('\n')
			command = strings.TrimSpace(command)
		}
		exitFlag = r.InputCommand(command)
		if exitFlag {
			break
		}
	}
}
