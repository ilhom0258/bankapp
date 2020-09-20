package payment

import "bank/pkg/bank/types"

//Max - finding maximum amount if payment
func Max(payments []types.Payment) types.Payment {
	for i := range payments {
		for j := range payments {
			if payments[i].Amount > payments[j].Amount {
				payments[i], payments[j] = payments[j], payments[i]
			}
		}
	}
	return payments[0]
}
