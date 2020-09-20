package card

import (
	"bank/pkg/bank/types"
)

// IssueCard - creating a card
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

// Withdraw - money withdraw
func Withdraw(card types.Card, amount types.Money) types.Card {

	if card.Active == true && card.Balance >= amount && amount > 0 && amount <= 20_000_00 {
		card.Balance -= amount
	}
	return card

}
