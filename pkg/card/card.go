package card

import (
	"01/pkg/money"
	"math/rand"
	"strconv"
	"strings"
)

type Currency string

const (
	Rub Currency = "RUB"
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
	IssuerId string
	Cards    []Card
}

func NewService() *Service {
	return &Service{
		IssuerId: "510621",
	}
}

func (s *Service) Add(card Card) *Card {
	s.Cards = append(s.Cards, card)
	return &card
}

func (s *Service) ByNumber(number string) (card *Card) {
	card = nil
	if s.isOurCard(number) {
		for i, c := range s.Cards {
			if c.Number == number {
				card = &s.Cards[i]
			}
		}
		if card == nil {
			card = s.NewCard("", 0, Rub, number)
		}
	}
	return
}

func (s *Service) isOurCard(number string) bool {
	if strings.HasPrefix(number, "510621") {
		return true
	}
	return false
}

func (s *Service) CheckByLuna(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	numberInString := strings.Split(number, "")
	numberInNumders := make([]int, 0)
	for s := range numberInString {
		if n, e := strconv.Atoi(numberInString[s]); e == nil {
			numberInNumders = append(numberInNumders, n)
		} else {
			return false
		}
	}
	sum := 0
	for n := range numberInNumders {
		if (n+1)%2 > 0 {
			numberInNumders[n] = numberInNumders[n] * 2
			if numberInNumders[n] > 9 {
				numberInNumders[n] = numberInNumders[n] - 9
			}
		}
		sum += numberInNumders[n]
	}
	if (((sum % 10) - 10) * -1) == 10 {
		return true
	}
	return false
}
