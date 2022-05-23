package products

import "fmt"

func CalculateDeposit(amount int64, currency string) (minAmount, maxAmount int64) {

	if amount <= 0 {
		return 0, 0
	}

	if bonus := int64(100_00); amount >= 1000_00 {
		amount += bonus
	}

	minPercent, maxPercent := percent4Currency(currency)
	minAmount = amount * int64(minPercent) / 100
	maxAmount = amount * int64(maxPercent) / 100
	return
}

func percent4Currency(curr string) (min, max int64) {
	// "TJS", "RUB", "USD"
	switch curr {
	case "RUB":
		fmt.Println("Поздравляем, у вас рубли!")
		return 4, 6
	case "USD":
		return 2, 4
	case "TJS":
		return 14, 17
	default:
		return 0, 0
	}
}
