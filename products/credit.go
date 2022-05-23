package products

import (
	. "gitproject/pkg"
)

func CalculateCredit(amount Money, yearPercent float64) Money {
	var mounthPay = amount * Money((1+yearPercent/100)/12)
	return mounthPay
}

var (
	Str1 = "Other text"
	Str2 = "SOME TEXT"
)
