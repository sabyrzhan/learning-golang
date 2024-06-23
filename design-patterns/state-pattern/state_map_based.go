package state_pattern

import "fmt"

// Here we are doing money transfer from one account to another one using P2P or International transfer type
// In case of International transfer we have multiple states of the transfer because each state is checked by bank for fraud
// In case of P2P it is possible to directly send from account to another one without any checks because of domesticity
// Main point of using map for states here is, current state points to next state(s). We can set our next action by
// selecting the needed state as next step.

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