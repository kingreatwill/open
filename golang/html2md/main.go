package main

import (
	"fmt"
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
		fmt.Println("html2md.exe url file selector")
	}
}

func do(url, file, selector string) {
	New(url, file, selector).Do()
}
