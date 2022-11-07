package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"push-swap/utils"
	"strings"
)

func main() {

	args := os.Args
	if len(args) == 2 {
		inputString := os.Args[1]
		var b []int
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
		
		reader := bufio.NewReader(os.Stdin)

		isEndOfInstructions := false
		for !isEndOfInstructions {
			command, err := reader.ReadString('\n')
			if err == io.EOF {
				if utils.IsSortAEmptyB(a, b) {
					fmt.Println("OK")
				} else {
					fmt.Println("KO")
				}
				break
			}

			command = strings.TrimRight(command, "\r\n")
			command = strings.TrimRight(command, "\n")
			command = strings.TrimRight(command, "\r")
			if command == "" { //now handles newline at end of commandsequence like "echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | go run main.go "0 9 1 8 2"" without error
				continue
			} else {
				commands := []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr", "\\n"}
				if utils.StringInSlice(command, commands) {
					utils.ExecuteCommand(&a, &b, command)
				} else {
					fmt.Println("Error")
				}
			}
		}
	}
}
