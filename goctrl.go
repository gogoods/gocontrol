package main

import (
	"fmt"
	"os"

	xfile "github.com/gogoods/x/file"
)

const (
	Usage = `
Usage:

	goctrl <appname>
	
	`
	SrcFile = "control.tmpl"
	DstFile = "control"

	OldContent = "<appname>"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println(Usage)
	} else {
		if err := xfile.Copy(SrcFile, DstFile); err == nil {
			err = xfile.ReplaceContent(DstFile, OldContent, os.Args[1])

			if err == nil {
				fmt.Println("Done!")
			} else {
				fmt.Println("Failed:", err.Error())
			}

		} else {
			fmt.Println("Failed:", err.Error())
		}

	}
}
