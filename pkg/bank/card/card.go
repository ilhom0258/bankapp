package card

import (
	"bank/pkg/bank/types"
	"math"
)

//IssueCard - card issuing
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

//Withdraw - witdrawing money
func Withdraw(card *types.Card, amount types.Money) {

	const withdrawLimit = 20_000_00

	if !card.Active {
		return
	}

	if amount > withdrawLimit {
		return
	}

	if amount < 0 {
		return
	}

	if card.Balance < amount {
		return
	}

	card.Balance -= amount
}

// Deposit - deposit money to card
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if amount > 50_000_00 {
		return
	}

	if amount < 0 {
		return
	}

	card.Balance += amount
}

// AddBonus - bonus adding
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}

	if card.Balance < 0 {
		return
	}

	if card.MinBalance < 0 {
		return
	}

	if card.Balance+card.MinBalance < 0 {
		return
	}
	bonus := types.Money(math.Floor(float64(card.MinBalance) * float64(percent) / 100.0 * float64(daysInMonth) / float64(daysInYear)))

	if bonus > 5_000_00 {
		return
	}

	card.Balance += bonus
}

// Total вычисляет общую сумму средств на всех картах
// Отрицательные суммы и суммы на заблокированных картах игнорируются
func Total(cards []types.Card) types.Money {
	total := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance < 0 {
			continue
		}
		total += card.Balance
	}
	return total
}
