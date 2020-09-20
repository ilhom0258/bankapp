package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	card := card.IssueCard(types.TJS, "black", "Infinity")
	fmt.Println(card)
}
