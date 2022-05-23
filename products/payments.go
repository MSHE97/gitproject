package products

import (
	"gitproject/pkg"
	"fmt"
)

func PrintCardTransactions(card types.Card) {
	for _, transaction := range card.TrnHistory {
		fmt.Printf("Сумма: %v, тип: %v\n", transaction.Amount, transaction.Type)
	}
}
