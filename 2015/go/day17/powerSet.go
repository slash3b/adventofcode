package main

func PowerSetWrapper(input []int) [][]int {
	results := [][]int(nil)

	powerSet(input, []int{}, 0, &results)

	return results
}

func powerSet(input, set []int, index int, result *[][]int) {

	if index == len(input) {
		*result = append(*result, set)
		return
	}

	elem := input[index]

	index++

	// do not append here
	powerSet(input, append([]int(nil), set...), index, result)
	set = append(set, elem)

	powerSet(input, append([]int(nil), set...), index, result)
}

//func has(set *[][]int, subset []int) bool {
//
//	for _, subset := range *set {
//
//		for _, value := range subset {
//
//		}
//	}
//
//	return false
//}

//func powerSet(input, set []int, result *[][]int) {
//
//	if len(input) == 0 {
//		*result = append(*result, set)
//		return
//	}
//
//	for i := 0; i < len(input); i++ {
//		powerSet(append([]int(nil), input[i+1:]...), set, result)
//		set = append(set, input[i])
//
//		secondCopySet := append([]int(nil), set...)
//		powerSet(append([]int(nil), input[i+1:]...), secondCopySet, result)
//	}
//
//}
