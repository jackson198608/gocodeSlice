package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const systemargslen = 3

func main() {
	getInputFromCommandLine()
	getInputFromInteractiveUsBufio()
	getInputUsFmtScan()

}

func getInputFromInteractiveUsBufio() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		} else if strings.Compare("exit", text) == 0 {
			fmt.Println("see you")
			break
		}

	}

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

func getInputUsFmtScan() {
	for {
		var i int
		fmt.Scanf("%d", &i)
		if i != 0 {
			fmt.Println(i)
		} else {
			fmt.Println("bye")
			return
		}
	}
}
