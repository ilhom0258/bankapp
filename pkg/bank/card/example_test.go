package card

import (
	"bank/pkg/bank/types"
	"fmt"
	"bank/pkg/bank/payment"
)

func ExampleWithdraw() {
	card := IssueCard(types.TJS, "black", "Ilhom")
	card.Balance = 10_000_00
	Withdraw(&card, 1_000_00)
	fmt.Println(card.Balance)
	//Output:
	//900000
}

func ExampleDeposit() {
	card := IssueCard(types.TJS, "black", "Ilhom")
	card.Balance = 10_000_00
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	card.Active = false
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	card.Active = true
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	//Output:
	//2000000
	//2000000
	//2000000
}

func ExampleAddBonus() {
	card := IssueCard(types.TJS, "black", "Ilhom")

	card.Balance = 0
	card.MinBalance = 10_000_00
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	card.Active = false
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	card.Active = true

	card.Balance = -5_000_00
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	card.Balance = 10_000_00

	card.MinBalance = 10_000_000_00
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	//Output:
	//2465
	//2465
	//-500000
	//1000000
}
func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	//Output:
	//3000000
}

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     99,
			Amount: 50000,
		},
		{
			ID:     1,
			Amount: 2,
		},
	}
	result := payment.Max(payments)
	fmt.Println(result.ID)
	//Output:
	//99
}

func ExamplePaymentSources() {
	cards := []types.Card{}
	for i := 0; i < 10; i++ {
		card := IssueCard(types.TJS, "black", "Ilhom")
		cards = append(cards,card)
	}
	paymentSources := PaymentSources(cards)
	for _, paymentSource := range paymentSources {
		fmt.Println(paymentSource.Number, paymentSource.Balance)
	}
	//Output:
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	// 5058 xxxx xxxx 0001 0
	
}
