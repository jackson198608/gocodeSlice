package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

func job3() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			fmt.Println("see you")
			return
		}

		//covert
		result := hex2int(text)
		fmt.Println(result)

	}

}
