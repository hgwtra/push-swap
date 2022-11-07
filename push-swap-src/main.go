package main

import (
	"fmt"
	"os"
	"push-swap/algo"
	"push-swap/utils"
	"strings"
)

func main() {

	args := os.Args
	if len(args) == 2 {
		inputString := os.Args[1]

		aStr := strings.Split(inputString, " ")
	
		a, err := utils.GetStack(aStr)
		if err != nil {
			fmt.Println("Error")
			return
		}

		if utils.CheckDuplicates(a) {
			fmt.Println("Error")
			return
		}

		result := algo.SortIndex(a)

		var b []int

		if len(a) < 50 {
			utils.PrintInstructions(result, b, len(result))
		} else {
			algo.RadixSortForTwoStack(result, b)
		}	
	}
}
