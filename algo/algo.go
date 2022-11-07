package algo

import (
	"push-swap/utils"
	"sort"
)
type Element struct {
	index int
	value int
}

func insertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > temp; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}

func Sort(array []int) []int {
	sorted := make([]int, len(array))

	copy(sorted, array)
	sort.Ints(sorted)
	return sorted
}

func SortIndex(array []int) []int {

	var initialStack []Element

	for index, element := range array {
		initialStack = append(initialStack, Element{index, element})
	}

	sorted := Sort(array)

	sortedOrder := make([]int, len(array))
	for sortedIndex, sortedElement := range sorted {
		for _, initialElement := range initialStack {
			if sortedElement == initialElement.value {
				sortedOrder[initialElement.index] = sortedIndex
				break
			}
		}
	}

	return sortedOrder
}

func RadixSortForTwoStack(a []int, b[]int) {

	size := len(a);
	maxNum := size - 1; // the biggest number in a is stack_size - 1
	maxBits := 0; // how many bits for maxNum 

	for (maxNum >> maxBits) != 0  {
		maxBits++;
	}

	// repeat for maxBits times
	for i := 0; i < maxBits ; i++ {
		for j := 0 ; j < size ; j++ {
			num := a[0]; // top number of A
			if (num >> i)&1 == 1 {
				utils.PrintAndExecute(&a,&b, "ra"); 
				// if the (i + 1)-th bit is 1, leave in stack a
			} else {
				utils.PrintAndExecute(&a,&b, "pb");
				// otherwise push to stack b
			}
		}
		// put into boxes done
		for len(b) != 0 { 
			utils.PrintAndExecute(&a,&b, "pa"); // while stack b is not empty, do pa
		}
		
		// connect numbers done
	}
}