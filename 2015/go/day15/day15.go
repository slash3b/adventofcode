package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ingredient struct {
	name    string
	cap     int
	dur     int
	flav    int
	texture int
	cal     int
}

func newIngredient(s string) ingredient {

	line := strings.Split(s, " ")

	//fmt.Println("LINE:", line)

	name := strings.TrimRight(line[0], ":")
	capacity, _ := strconv.Atoi(strings.TrimRight(line[2], ", "))
	durability, _ := strconv.Atoi(strings.TrimRight(line[4], ", "))
	flavour, _ := strconv.Atoi(strings.TrimRight(line[6], ", "))
	texture, _ := strconv.Atoi(strings.TrimRight(line[8], ", "))
	calories, _ := strconv.Atoi(strings.TrimRight(line[10], ", "))

	//fmt.Println(name, capacity, durability, flavour, texture, calories)

	return ingredient{
		name:    name,
		cap:     capacity,
		dur:     durability,
		flav:    flavour,
		texture: texture,
		cal:     calories,
	}
}

func main() {

	a := `Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
	Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
	Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
	Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1`
	//b := `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	//Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

	input := strings.Split(a, "\n")
	var ingredients []ingredient
	for _, v := range input {
		ingredients = append(ingredients, newIngredient(v))
	}

	/*
		Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
		Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3

		Then, choosing to use 44 teaspoons of butterscotch and 56 teaspoons of cinnamon (because the amounts of each ingredient must add up to 100) would result in a cookie with the following properties:

		    A capacity of 44*-1 + 56*2 = 68
		    A durability of 44*-2 + 56*3 = 80
		    A flavor of 44*6 + 56*-2 = 152
		    A texture of 44*3 + 56*-1 = 76

		Multiplying these together (68 * 80 * 152 * 76, ignoring calories for now) results in a total score of 62842880,
	*/
	calcScore := func(i, j, k, c int) int {

		totalCapacity := i*ingredients[0].cap + j*ingredients[1].cap + k*ingredients[2].cap + c*ingredients[3].cap
		totalDuration := i*ingredients[0].dur + j*ingredients[1].dur + k*ingredients[2].dur + c*ingredients[3].dur
		totalFlavor := i*ingredients[0].flav + j*ingredients[1].flav + k*ingredients[2].flav + c*ingredients[3].flav
		totalTexture := i*ingredients[0].texture + j*ingredients[1].texture + k*ingredients[2].texture + c*ingredients[3].texture
		//totalCalories := i*ingredients[0].cal + j*ingredients[1].cal + k*ingredients[2].cal + c*ingredients[3].cal

		if totalCapacity <= 0 || totalDuration <= 0 || totalFlavor <= 0 || totalTexture <= 0 {
			return 0
		}

		return totalCapacity * totalDuration * totalFlavor * totalTexture
	}

	max := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				for c := 0; c <= 100; c++ {
					if i+j+k+c >= 100 {
						continue
					}
					result := calcScore(i, j, k, c)
					if result > max {
						max = result
					}

				}
			}
		}
	}

	//fmt.Printf("Ingredients: %+v \n", ingredients)
	fmt.Printf("MAX: %+v \n", max)
	fmt.Println("incorrects", 3458145636)

}
