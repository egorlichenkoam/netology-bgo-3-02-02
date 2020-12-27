package transaction

import (
	"01/pkg/card"
	"01/pkg/money"
	"math/rand"
	"time"
)

type Type int

const (
	FROM Type = iota
	TO
)

type Status string

const (
	Ok   Status = "Ok"
	Fail        = "Fail"
	Wait        = "Wait"
)

type Transaction struct {
	Id       int64
	Amount   money.Money
	Datetime int64
	Mcc      Mcc
	Status   Status
	Card     *card.Card
	Type     Type
}

type Service struct {
	Transactions []Transaction
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateTransaction(amount money.Money, mcc Mcc, card *card.Card, fromTo Type) *Transaction {
	tx := Transaction{
		Id:       rand.Int63(),
		Amount:   amount,
		Datetime: time.Now().Unix(),
		Mcc:      mcc,
		Status:   Wait,
		Card:     card,
		Type:     fromTo,
	}
	s.Transactions = append(s.Transactions, tx)
	return s.ById(tx.Id)
}

func (s *Service) ById(id int64) *Transaction {
	for i, tx := range s.Transactions {
		if tx.Id == id {
			return &s.Transactions[i]
		}
	}
	return nil
}
func (s *Service) ByCard(card *card.Card) []Transaction {
	result := make([]Transaction, 0)
	for _, transaction := range s.Transactions {
		if transaction.Card.Id == card.Id {
			result = append(result, transaction)
		}
	}
	return result
}

func (s *Service) LastNTransactions(card *card.Card, n int) []Transaction {
	transactions := s.ByCard(card)
	if len(transactions) < n {
		n = len(transactions)
	}
	n = len(transactions) - n
	transactions = transactions[n:]
	for i := len(transactions)/2 - 1; i >= 0; i-- {
		flipIdx := len(transactions) - 1 - i
		transactions[i], transactions[flipIdx] = transactions[flipIdx], transactions[i]
	}
	return transactions
}

func (s *Service) SumByMcc(card *card.Card, mccs []Mcc) money.Money {
	var result money.Money = 0
	transactions := filterTransactionsByMcc(s.ByCard(card), mccs)
	for _, transaction := range transactions {
		result = result + transaction.Amount
	}
	return result
}

func filterTransactionsByMcc(transactions []Transaction, mccs []Mcc) []Transaction {
	result := make([]Transaction, 0)
	for _, transaction := range transactions {
		for _, mcc := range mccs {
			if transaction.Mcc == mcc {
				result = append(result, transaction)
			}
		}
	}
	return result
}

func (s *Service) TranslateMcc(code Mcc) string {
	result := "Категория не указана"
	value, ok := Mccs()[code]
	if ok {
		result = value
	}
	return result
}
