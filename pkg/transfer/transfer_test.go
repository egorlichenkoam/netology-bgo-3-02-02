package transfer

import (
	"01/pkg/card"
	"01/pkg/money"
	"01/pkg/transaction"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		TransferSvc *Service
	}
	type args struct {
		from   string
		to     string
		amount money.Money
	}

	cardSvc := card.NewService("510621")
	transactionSvc := transaction.NewService()
	inBank := Commission{
		Percent: 0,
		Minimum: 0,
	}
	toDifferentBank := Commission{
		Percent: 0.5,
		Minimum: 10_00,
	}
	betweenDifferentBank := Commission{
		Percent: 1.5,
		Minimum: 30_00,
	}
	transferSvc := NewService(cardSvc, transactionSvc, inBank, toDifferentBank, betweenDifferentBank)

	cardSvc.NewCard("BABANK", 10_000_00, card.Rub, "5106212879499054")
	cardSvc.NewCard("BABANK", 22_433_00, card.Rub, "5106212548197220")
	cardSvc.NewCard("BABANK", 15_000_00, card.Rub, "5106211562724463")
	cardSvc.NewCard("BABANK", 30_000_00, card.Rub, "5106219146702939")
	cardSvc.NewCard("BABANK", 55_000_00, card.Rub, "5106218923315543")
	cardSvc.NewCard("BABANK", 10_500_00, card.Rub, "5106214088426217")
	cardSvc.NewCard("BABANK", 10_900_00, card.Rub, "5106217924694328")

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal money.Money
		wantError error
	}{
		{
			name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   cardSvc.Cards[0].Number,
				to:     cardSvc.Cards[1].Number,
				amount: 1_000_00,
			},
			wantTotal: 1_000_00,
			wantError: nil,
		}, {
			name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   cardSvc.Cards[2].Number,
				to:     cardSvc.Cards[3].Number,
				amount: 20_000_00,
			},
			wantTotal: 20_000_00,
			wantError: errNotEnoughMoney,
		}, {
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   cardSvc.Cards[4].Number,
				to:     "0200000000000000",
				amount: 20_000_00,
			},
			wantTotal: 0,
			wantError: errCardNumberInvalid,
		}, {
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   cardSvc.Cards[5].Number,
				to:     "4106217775856128",
				amount: 20_000_00,
			},
			wantTotal: 20_100_00,
			wantError: errCardToNotFound,
		}, {
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   "4106215234507001",
				to:     cardSvc.Cards[6].Number,
				amount: 20_000_00,
			},
			wantTotal: 20_000_00,
			wantError: errCardFromNotFound,
		}, {
			name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   "4106217734669026",
				to:     "0000000000000000",
				amount: 20_000_00,
			},
			wantTotal: 20_300_00,
			wantError: errCardFromNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, gotError := transferSvc.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			t.Log(gotTotal, tt.wantTotal, gotError, tt.wantError)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotError != tt.wantError {
				t.Errorf("Card2Card() got = %v, want %v", gotError, tt.wantError)
			}
		})
	}
}
