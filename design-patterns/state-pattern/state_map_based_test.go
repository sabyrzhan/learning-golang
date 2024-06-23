package state_pattern

import "testing"

func TestMoneyTransferInternational(t *testing.T) {
	mt := MoneyTransfer {
		Type: International,
		State: Created,
		Amount: 200,
		From: "John Doe",
		To: "Jane Doe",
	}

	mt.Send()
}

func TestMoneyTransferP2P(t *testing.T) {
	mt := MoneyTransfer {
		Type: P2P,
		State: Created,
		Amount: 200,
		From: "John Doe",
		To: "Jane Doe",
	}

	mt.Send()
}