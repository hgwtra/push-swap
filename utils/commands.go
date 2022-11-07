package utils

import (
	"fmt"
	"strconv"
)

func IsSortAEmptyB(a []int, b []int) bool {

	isResSorted := isSortedA(a)
	return len(b) == 0 && isResSorted
}

func isSortedA(a []int) bool {

	for index, item := range a {
		if index > 0 {
			if item < a[index-1] {
				return false
			}
		}
	}
	return true
}

func GetStack(inputAStack []string) (a []int, err error) {
	for _, item := range inputAStack {
		tempItem, err := strconv.Atoi(item)
		if err == nil {
			a = append(a, tempItem)
		} else {
			return a, err
		}
	}

	return a, err
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ExecuteCommand(a *[]int, b *[]int, command string) {
	switch command {
	case "pa":
		pushPop(a, b)
	case "pb":
		pushPop(b, a)
	case "sa":
		swapOne(a)
	case "sb":
		swapOne(b)
	case "ss":
		swapOne(a)
		swapOne(b)
	case "ra":
		rotate(a)
	case "rb":
		rotate(b)
	case "rr":
		rotate(a)
		rotate(b)
	case "rra":
		rotateReverse(a)
	case "rrb":
		rotateReverse(b)
	case "rrr":
		rotateReverse(a)
		rotateReverse(b)
	}
}

func PrintAndExecute(a *[]int, b *[]int, command string) {
	ExecuteCommand(a, b, command)
	fmt.Println(command)
}

/*
a []int - array of indexes of sorted a
*/
func PrintInstructions(a []int, b []int, initialLenA int) {

	//done condition (break recursion)
	if IsSortAEmptyB(a, b) {
		return
	}

	// if the greatest element in the top put it to the bottom
	for a[0] > a[len(a)-1] {
		PrintAndExecute(&a, &b, "ra")

		if a[1] > a[len(a)-1] && isSwapNeeded(a, true) {
			PrintAndExecute(&a, &b, "sa")
		}
	}

	//if the smalest element in the bottom put it to the top
	if len(a) > 1 && a[len(a)-1] == 0 {
		PrintAndExecute(&a, &b, "rra")
	}

	//if a is sorted push element from b
	if isSortedA(a) {
		PrintAndExecute(&a, &b, "pa")
	} else {
		if isSwapNeeded(b, false) && isSwapNeeded(a, true) {
			PrintAndExecute(&a, &b, "ss")
		} else {
			if isSwapNeeded(b, false) {
				PrintAndExecute(&a, &b, "sb")
			}
			if isSwapNeeded(a, true) {
				PrintAndExecute(&a, &b, "sa")
			}
		}

		if !isSortedA(a) {
			PrintAndExecute(&a, &b, "pb")
		}
	}

	PrintInstructions(a, b, initialLenA)
}

func isSwapNeeded(a []int, isAsc bool) bool {
	isNeeded := false
	if len(a) > 1 {
		if isAsc {
			isNeeded = a[0] > a[1]
		} else {
			isNeeded = a[0] < a[1]
		}
	}

	return isNeeded
}

// push the top first element of stack b to stack a
func pushPop(a *[]int, b *[]int) {
	if len(*b) > 0 {
		*a = append([]int{(*b)[0]}, (*a)...)
		*b = (*b)[1:]
	}
}

// swap first 2 elements of stack
func swapOne(a *[]int) {
	if len(*a) > 1 {
		(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
	}
}

// rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func rotate(a *[]int) {
	if len(*a) > 1 {
		*a = append((*a)[1:], (*a)[0])
	}
}

// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func rotateReverse(a *[]int) {
	if len(*a) > 1 {
		*a = append([]int{(*a)[len(*a)-1]}, (*a)[:(len(*a)-1)]...)
	}
}

func CheckDuplicates(a []int) bool {
	allKeys := make(map[int]bool)
	for _, item := range a {
		if !allKeys[item] {
			allKeys[item] = true
		} else {
			return true
		}
	}
	return false
}
