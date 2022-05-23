package products

import (
	"gitproject/pkg"
)

func Transfer(card types.Card, amount types.Money) types.Card {
	card.Balance += amount
	return card
}

func Withdraw(amount types.Money, card * types.Card) {
	if amount <= 0 {
		return
	}
	if !card.Activity {
		return
	}
	if card.Balance < amount {
		return
	}
	card.Balance -= amount

	transaction := types.Transaction{
		Amount: amount,
		Type:   "withdraw",
	}

	card.TrnHistory = append(card.TrnHistory, transaction)
}
