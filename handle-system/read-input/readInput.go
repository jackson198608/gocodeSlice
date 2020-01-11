package main

import (
	"fmt"
	"os"
)

const systemargslen = 3

func main() {
	getInputFromCommandLine()

}

func getInputFromCommandLine() {
	if len(os.Args) != systemargslen {
		fmt.Println("not enough argument")
		return
	}
	first := os.Args[0] //executable programe name
	second := os.Args[1]
	third := os.Args[2]
	fmt.Println("args: ", first, " ", second, " ", third)

}
