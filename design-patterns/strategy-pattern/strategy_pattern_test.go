package strategy_pattern

import "testing"

func TestCompute(t *testing.T) {
	c := Computer{}
	c.ProcessData("testData", SingleCoreCompute{})
	c.ProcessData("testData", MultiCoreCompute{Cores: 4})
	c.ProcessData("testData", ArmCpuCompute{})
}
