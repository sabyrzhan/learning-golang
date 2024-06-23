package strategy_pattern

import "fmt"

// Here we are executing our data on multiple platforms with variable core CPUs
// Strategy lets us change algorithm at run-time by specifying different implementations

type Compute interface {
	Compute(data interface{})
}

type SingleCoreCompute struct {}
type MultiCoreCompute struct {
	Cores int
}
type ArmCpuCompute struct {}

func (s SingleCoreCompute) Compute(data interface{}) {
	fmt.Println(fmt.Sprintf("Executing data on single core CPU machine"))
}

func (m MultiCoreCompute) Compute(data interface{}) {
	fmt.Println(fmt.Sprintf("Executing data on multi core CPU machine:"))
	for i := 0; i < m.Cores; i++ {
		fmt.Println(fmt.Sprintf("CORE: %d", i))
	}
}

func (a ArmCpuCompute) Compute(data interface{}) {
	fmt.Println(fmt.Sprintf("Executing data on arm CPU machine"))
}

type Computer struct {}
func (c Computer) ProcessData(data interface{}, strategy Compute) {
	strategy.Compute(data)
}