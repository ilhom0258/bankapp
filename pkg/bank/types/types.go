package types

// Money представляют собой денежную сумму в минимальных единицах(центы, копейки, дирамы)
type Money int64

// Currency представляют код валюты
type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN - номер карты
type PAN string

//Card - платежная карта
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
