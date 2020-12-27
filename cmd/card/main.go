package main

import (
	"01/pkg/card"
	"01/pkg/transaction"
	"01/pkg/transfer"
	"fmt"
)

func main() {
	cardSvc := card.NewService()
	transactionSvc := transaction.NewService()
	inBank := transfer.Commission{
		Percent: 0,
		Minimum: 0,
	}
	toDifferentBank := transfer.Commission{
		Percent: 0.5,
		Minimum: 10_00,
	}
	betweenDifferentBank := transfer.Commission{
		Percent: 1.5,
		Minimum: 30_00,
	}
	transferSvc := transfer.NewService(cardSvc, transactionSvc, inBank, toDifferentBank, betweenDifferentBank)

	cardSvc.NewCard("BANK", 10_000_00, card.RUB, "4263141548036728")
	cardSvc.NewCard("BANK", 20_000_00, card.RUB, "4759718447175045")
	cardSvc.NewCard("BANK", 30_000_00, card.RUB, "4806551844152926")

	printCards(cardSvc.Cards)
	printTransactions(transactionSvc.Transactions)

	fmt.Println(transferSvc.Card2Card("4806551844152926", "4759718447175045", 10_000_00))
	fmt.Println(transferSvc.Card2Card("4759718447175045", "4263141548036728", 10_000_00))
	fmt.Println(transferSvc.Card2Card("4263141548036728", "4450209454897335", 10_000_00))
	fmt.Println(transferSvc.Card2Card("4806551844152926", "4759718447175045", 5_000_00))
	fmt.Println(transferSvc.Card2Card("4806551844152926", "4759718447175045", 2_000_00))
	fmt.Println(transferSvc.Card2Card("4759718447175045", "4806551844152926", 8_000_00))
	fmt.Println(transferSvc.Card2Card("4806551844152926", "4759718447175045", 40_000_00))

	fmt.Println("")

	printCards(cardSvc.Cards)
	printTransactions(transactionSvc.Transactions)
	printCards(cardSvc.Cards)
}

func printCards(cards []card.Card) {
	for _, c := range cards {
		fmt.Println(c)
	}
	fmt.Println("")
}

func printTransactions(txs []transaction.Transaction) {
	for _, tx := range txs {
		fmt.Println(tx, tx.Card.Number)
	}
	fmt.Println("")
}
