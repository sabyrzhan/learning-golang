package visitor_pattern

import "testing"

func TestVisitorPattern(t *testing.T) {
	c := []MonitoringEnabled{RAM{32}, CPU{4},Motherboard{"v1.0.0"},HDD{4}}
	m := StdoutMonitoring{}
	for _, v := range c {
		v.Monitor(m)
	}
}
