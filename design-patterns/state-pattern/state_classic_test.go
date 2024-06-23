package state_pattern

import "testing"

func TestClassicState(t *testing.T) {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
