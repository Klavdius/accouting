package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

func main() {
	entries, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("error")
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println("error")
		}
		reg, _ := regexp.MatchString("Acc_", info.Name())
		if reg {
			infos = append(infos, info)
		}
	}
	CheckingLengInfos(infos)
	DisplayFile(infos)

	for {
		var in *bufio.Reader
		var out *bufio.Writer
		in = bufio.NewReader(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
		defer out.Flush()

		var (
			command  string
			exitFlag bool
		)
		fmt.Fscan(in, &command)
		exitFlag = RouterCommand(command)
		if exitFlag {
			break
		}
	}
}
