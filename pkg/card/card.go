package card

import (
	"01/pkg/money"
	"math/rand"
)

type Currency string

const (
	RUB Currency = "RUB"
)

type Card struct {
	Id       int64
	Issuer   string
	Balance  money.Money
	Currency Currency
	Number   string
	Icon     string
}

func (s *Service) NewCard(issuer string, balance money.Money, currency Currency, number string) *Card {
	return s.Add(Card{Id: rand.Int63(), Issuer: issuer, Balance: balance, Currency: currency, Number: number, Icon: ""})
}

type Service struct {
	Cards []Card
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Add(card Card) *Card {
	s.Cards = append(s.Cards, card)
	return &card
}

func (s *Service) ByNumber(number string) (card *Card) {
	for i, c := range s.Cards {
		if c.Number == number {
			return &s.Cards[i]
		}
	}
	return nil
}
