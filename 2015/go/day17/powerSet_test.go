// +build ignore

package main

import (
	"fmt"
	"testing"
)

func TestPowerSet(t *testing.T) {

	results := [][]int(nil)
	powerSet([]int{1, 2, 3}, []int{}, 0, &results)

	fmt.Println(results)
}
