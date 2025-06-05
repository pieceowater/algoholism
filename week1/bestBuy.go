package week1

import "math"

func BestBuy(stocks []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for i := 0; i < len(stocks); i++ {
		if stocks[i] < minPrice {
			minPrice = stocks[i]
		} else {
			profit := stocks[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
