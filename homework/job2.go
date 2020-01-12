package main

import (
	"fmt"
	"sort"
)

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func job2() {
	for {
		//get input n
		var inputN int
		fmt.Scanf("%d", &inputN)
		if inputN == 0 {
			break
		}
		var inputArr []int = make([]int, inputN)

		//get input arr
		for i := 0; i < inputN; i++ {
			fmt.Scanf("%d", &inputArr[i])
		}

		//sort
		sort.Ints(inputArr)

		//uniq
		inputArr = unique(inputArr)

		//print result
		for _, e := range inputArr {
			fmt.Println(e)
		}
	}

}
