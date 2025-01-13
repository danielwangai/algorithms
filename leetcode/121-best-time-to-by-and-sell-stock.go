package leetcode

func maxProfit(prices []int) int {
	return maxProfit1(prices)
	//return maxProfit2(prices)
}

// Time: O(n²)
// Space: O(1)
func maxProfit1(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if (prices[j] - prices[i]) > max {
				max = prices[j] - prices[i]
			}
		}
	}

	if max < 0 {
		return 0
	}

	return max
}

// Time: O(n)
// Space: O(1)
func maxProfit2(prices []int) int {
	/*
	   loop updating the lowest price
	*/
	lowest := prices[0]
	max := 0
	for i, price := range prices {
		if i == 0 {
			continue
		}
		if price < lowest {
			lowest = price
			continue
		}
		if (price - lowest) > max {
			// max = int(math.Max(float64(max), float64(price - lowest)))
			max = price - lowest
		}
	}

	return max
}
