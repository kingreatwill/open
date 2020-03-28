package main

import (
	"fmt"
	"html2md/lib"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	switch len(os.Args) {
	case 2:
		do(os.Args[1], "", "")
	case 3:
		do(os.Args[1], os.Args[2], "")
	case 4:
		do(os.Args[1], os.Args[2], os.Args[3])
	default:
		fmt.Println("html2md.exe url file slector")
	}
}

func do(url, file, slector string) {
	suport := lib.New(url, slector, file)
	lib.Do(suport)
}
