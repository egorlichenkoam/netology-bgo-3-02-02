package main

import (
	"01/pkg/card"
	"fmt"
	"time"
)

func main() {
	master := &card.Card{
		Id:       1,
		Issuer:   "Gipermarket",
		Balance:  1000000,
		Currency: "",
		Number:   "7643-1237-2383-9386",
		Icon:     "masterCard.png",
		Transactions: []card.Transaction{
			{
				Id:       1,
				Amount:   -73555,
				Datetime: time.Now().Unix(),
				Mcc:      "5411",
				Status:   "DONE",
			},
			{
				Id:       2,
				Amount:   200000,
				Datetime: time.Now().Unix(),
				Mcc:      "0000",
				Status:   "DONE",
			},
			{
				Id:       3,
				Amount:   -120391,
				Datetime: time.Now().Unix(),
				Mcc:      "5812",
				Status:   "DONE",
			}},
	}
	transaction := &card.Transaction{
		Id:       4,
		Amount:   99900,
		Datetime: time.Now().Unix(),
		Mcc:      "5555",
		Status:   "DONE",
	}
	card.AddTransaction(master, transaction)
	fmt.Println(card.SumByMcc(master.Transactions, []card.Mcc{"5411", "0000"}))
	fmt.Println(card.TranslateMcc(transaction.Mcc))
	fmt.Println(card.LastNTransactions(master, 2))
	fmt.Println(master.Transactions)
	fmt.Println(card.LastNTransactions2(master, 2))
	fmt.Println(master.Transactions)
}
