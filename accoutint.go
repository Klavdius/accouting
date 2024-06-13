package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"regexp"
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

func CreatListAboutFileInDirect(dir []fs.DirEntry) []string {
	infos := make([]string, 0, len(dir))
	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			fmt.Println("error")
		}
		reg, _ := regexp.MatchString("\\.txt", info.Name())
		if reg {
			infos = append(infos, info.Name())
		}
	}
	return infos
}

func NeedListOfFiles(r RouterCommand) bool {
	if r.account.name == "" {
		return true
	}
	return false
}
