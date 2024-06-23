package state_pattern

import "fmt"

type MoneyTransferState int
const (
	Created MoneyTransferState = iota
	OnHold
	Approved
	Sent
	Failed
)

func (m MoneyTransferState) String() string {
	switch m {
	case Created:
		return "Created"
	case OnHold:
		return "OnHold"
	case Approved:
		return "Approved"
	case Sent:
		return "Sent"
	case Failed:
		return "Failed"
	default:
		return "Unknown state"
	}
}

type MoneyTransferType int
const (
	International MoneyTransferType = iota
	P2P
)

func (m MoneyTransferType) String() string {
	switch m {
	case International:
		return "International"
	case P2P:
		return "P2P"
	default:
		return "Unsupported type"
	}
}

type MoneyTransfer struct {
	Type MoneyTransferType
	State MoneyTransferState
	Amount int
	From string
	To string
}

var internationalStateMap = map[MoneyTransferState][]MoneyTransferState{
	Created: []MoneyTransferState {
		OnHold,
		Failed,
	},
	OnHold: []MoneyTransferState {
		Approved,
		Failed,
	},
	Approved: []MoneyTransferState {
		Sent,
		Failed,
	},
}

var p2pStateMap = map[MoneyTransferState][]MoneyTransferState{
	Created: []MoneyTransferState{
		Sent, Failed,
	},
}

func (m *MoneyTransfer) Send() {
	state := m.State
	fmt.Println(fmt.Sprintf("Sending money %d amount from %s to %s account using %s type", m.Amount, m.From, m.To, m.Type))
	failMessage := ""
	for state != Failed && state != Sent {
		var states []MoneyTransferState
		if m.Type == International {
			states = internationalStateMap[state]
		} else {
			states = p2pStateMap[state]
		}

		fmt.Println(fmt.Sprintf("%s: %d", state, m.Amount))
		if m.Amount < 200 {
			state = states[1]
			failMessage = "(Minimum amount must be 200)"
		} else {
			state = states[0]
		}
	}
	fmt.Println(fmt.Sprintf("%s: %d %s", state, m.Amount, failMessage))
	fmt.Println("Money transfer complete")
}