package main

import (
	"fmt"
	"log"
	"os"

	xfile "github.com/gogoods/x/file"
)

const (
	Usage = `
Usage:

	goctrl <appname>
	
	`
	SrcFile = "control.tmpl"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println(Usage)
	} else {
		log.Println(xfile.Copy(SrcFile, "./control"))
	}
}
