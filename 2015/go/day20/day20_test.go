package main

import (
	"fmt"
	"testing"
)

func TestFindHouseByPresents(t *testing.T) {
	ans := findHouseByPresentsNum(10)
	if ans != 1 {
		t.Error("incorrect")
	}

	presents := []int{10, 30, 40, 70, 60, 120, 80, 150, 130}

	for _, value := range presents {
		ans = findHouseByPresentsNum(value)
		fmt.Println(ans)
		//if ans != 1 {
		//	t.Error("incorrect")
		//}
	}

}
