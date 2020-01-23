package main

import "fmt"

func main() {
	// fmt.Println(get_max_profit([]int{10, 7, 5, 8, 11, 9}))
	// fmt.Println(get_max_profit([]int{}))
	fmt.Println(get_max_profit([]int{99, 5}))

	// fmt.Println(get_max_profit([]int{5, 9, 11, 12, 4, 9, 12}))
	// fmt.Println(get_max_profit([]int{5, 9, 11, 12, 4, 9, 10}))
	// fmt.Println(get_max_profit([]int{9, 6, 3, 1}))

}

type price struct {
	value int
	index int
}

func get_max_profit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var maxPrice int
	var minPrice int

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if price > maxPrice {
			maxPrice = price
		}
	}
	var maxProfit = minPrice - maxPrice

	currentMin := &price{prices[0], 0}
	currentMax := &price{prices[0], 0}
	for index, p := range prices[1:] {
		if p < currentMin.value {
			currentMin = &price{p, index + 1}
			currentMax = &price{p, index + 1}
		} else if p > currentMax.value {
			currentMax = &price{p, index + 1}
		}

		if currentMin.index != currentMax.index {
			fmt.Println(currentMax)
			profit := currentMax.value - currentMin.value

			if profit > maxProfit {
				maxProfit = profit
			}
		}

	}
	return maxProfit
}
