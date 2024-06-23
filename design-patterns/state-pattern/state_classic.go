package state_pattern

import "fmt"

// Class example of state with light switcher
// We define 2 states: OnState and OffState, where both of them have the default implementations from BaseState
// The point here is: the state of the object is changed by the states themselves
// In case of OnState: we override Off() method which sets switch's state to off.
// In case of OffState: we override On() method which sets switch's state to on.

type Switch struct {
	State State
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct {}

func (s *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

type OffState struct {
	BaseState
}

func NewOnState() *OnState {
	return &OnState{BaseState{}}
}

func NewOffState() *OffState {
	return &OffState{BaseState{}}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

func (os *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off")
	sw.State = NewOffState()
}

func (os *OffState) On(sw *Switch) {
	fmt.Println("Turning light on")
	sw.State = NewOnState()
}
