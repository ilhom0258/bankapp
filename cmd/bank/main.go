package main

<<<<<<< HEAD
func main() {

=======
import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	card := card.IssueCard(types.TJS, "black", "Infinity")
	fmt.Println(card)
>>>>>>> 8984773dea0a3995a7ba94c524ebaa8a3ade4843
}
