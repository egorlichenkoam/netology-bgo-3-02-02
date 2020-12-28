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

	cardSvc := card.NewService()
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

	cardSvc.NewCard("BABANK", 10_000_00, card.RUB, "4018682190154150")
	cardSvc.NewCard("BABANK", 5_000_00, card.RUB, "4105733741399564")
	cardSvc.NewCard("BABANK", 15_000_00, card.RUB, "4922876603093402")
	cardSvc.NewCard("BABANK", 30_000_00, card.RUB, "4084227961096153")
	cardSvc.NewCard("BABANK", 55_000_00, card.RUB, "4879888487800649")
	cardSvc.NewCard("BABANK", 10_500_00, card.RUB, "4772438185495983")
	cardSvc.NewCard("BABANK", 10_900_00, card.RUB, "4409713590773955")

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
				from:   "4018682190154150",
				to:     "4105733741399564",
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
				from:   "4922876603093402",
				to:     "4084227961096153",
				amount: 20_000_00,
			},
			wantTotal: 20_000_00,
			wantError: ErrorNotEnoughMoney,
		}, {
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   "4879888487800649",
				to:     "0000000000000000",
				amount: 20_000_00,
			},
			wantTotal: 20_100_00,
			wantError: nil,
		}, {
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   "4772438185495983",
				to:     "0000000000000000",
				amount: 20_000_00,
			},
			wantTotal: 20_100_00,
			wantError: ErrorNotEnoughMoney,
		}, {
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   "0000000000000000",
				to:     "4772438185495983",
				amount: 20_000_00,
			},
			wantTotal: 20_000_00,
			wantError: nil,
		}, {
			name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				TransferSvc: transferSvc,
			},
			args: args{
				from:   "0000000000000000",
				to:     "0000000000000000",
				amount: 20_000_00,
			},
			wantTotal: 20_300_00,
			wantError: nil,
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
