package transfer

import (
	"01/pkg/card"
	"01/pkg/money"
	"01/pkg/transaction"
)

type Commission struct {
	Percent float64
	Minimum money.Money
}

type Service struct {
	CardSvc              *card.Service
	TransactionSvc       *transaction.Service
	InBank               Commission
	ToDifferentBank      Commission
	BetweenDifferentBank Commission
}

func NewService(cardSvc *card.Service, transactionSvc *transaction.Service, inBank Commission, toDifferentBank Commission, betweenDifferentBank Commission) *Service {
	return &Service{
		CardSvc:              cardSvc,
		TransactionSvc:       transactionSvc,
		InBank:               inBank,
		ToDifferentBank:      toDifferentBank,
		BetweenDifferentBank: betweenDifferentBank,
	}
}

func (s *Service) Card2Card(from, to string, amount money.Money) (total money.Money, ok bool) {
	ok = true
	cardFrom := s.CardSvc.ByNumber(from)
	cardTo := s.CardSvc.ByNumber(to)
	total = s.total(amount, s.commission(cardFrom, cardTo))
	if cardFrom != nil {
		ok = s.transfer(cardFrom, total, transaction.FROM)
	}
	if cardTo != nil && ok {
		ok = s.transfer(cardTo, amount, transaction.TO)
	}
	return
}

func (s *Service) commission(cardFrom, cardTo *card.Card) *Commission {
	if cardFrom == nil && cardTo == nil {
		return &s.BetweenDifferentBank
	}
	if cardFrom != nil && cardTo == nil {
		return &s.ToDifferentBank
	}
	return &s.InBank
}

func (s *Service) total(amount money.Money, commission *Commission) money.Money {
	internalCommission := money.Money(float64(amount) / 100 * commission.Percent)
	if internalCommission < commission.Minimum {
		internalCommission = commission.Minimum
	}
	return amount + internalCommission
}

func (s *Service) transfer(card *card.Card, amount money.Money, fromTo transaction.Type) (ok bool) {
	ok = false
	tx := s.TransactionSvc.CreateTransaction(amount, "", card, fromTo)
	if fromTo == transaction.FROM {
		if card.Balance >= amount {
			card.Balance -= amount
			tx.Status = transaction.Ok
			ok = true
		} else {
			tx.Status = transaction.Fail
		}
	} else if fromTo == transaction.TO {
		card.Balance += amount
		tx.Status = transaction.Ok
		ok = true
	}
	return
}
