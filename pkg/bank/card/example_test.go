package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := IssueCard(types.TJS, "black", "Ilhom")
	card.Balance += 400000
	result := Withdraw(card, 40000)
	fmt.Println(result.Balance)
	//Output:
	//360000
}

func ExampleWithdraw_noMoney(){
	card := IssueCard(types.USD, "black", "Acer")
	result := Withdraw(card,30000)
	fmt.Println(result.Balance)
	//Output:
	//0
}

func ExampleWithdraw_inactive(){
	card := IssueCard(types.USD, "black", "Acer")
	card.Active = false
	result := Withdraw(card,30000)
	fmt.Println(result.Balance)
	//Output:
	//0
}

func ExampleWithdraw_limit(){
	card := IssueCard(types.USD, "black", "Acer")
	card.Balance = 20_000_000_00
	result := Withdraw(card,3_000_000_00)
	fmt.Println(result.Balance)
	//Output:
	//2000000000
}