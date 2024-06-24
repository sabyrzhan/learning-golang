package visitor_pattern

import "fmt"

// The example here is monitoring of PC component. Each component notifying monitoring system about itself after PC startup.
// Visitor is implemented using Double Dispatch technique.
// Visitor here is StdoutMonitoring which implements Monitoring interface.
// All the PC components accept visitor object by implementing MonitoringEnabled interface.

type MonitoringEnabled interface {
	Monitor(monitoring Monitoring)
}
type RAM struct {
	Size int
}
type CPU struct {
	Cores int
}
type Motherboard struct {
	BiosVersion string
}
type HDD struct {
	Size int
}

type Monitoring interface {
	MonitorRam(ram RAM)
	MonitorCPU(cpu CPU)
	MonitorMotherboard(motherboard Motherboard)
	MonitorHDD(hdd HDD)
}

type StdoutMonitoring struct {}
func (s StdoutMonitoring) MonitorRam(ram RAM) {
	fmt.Println(fmt.Sprintf("RAM size: %d GB", ram.Size))
}

func (s StdoutMonitoring) MonitorCPU(cpu CPU) {
	fmt.Println(fmt.Sprintf("CPU cores: %d", cpu.Cores))
}

func (s StdoutMonitoring) MonitorMotherboard(motherboard Motherboard) {
	fmt.Println(fmt.Sprintf("MB biosversion: %s", motherboard.BiosVersion))
}

func (s StdoutMonitoring) MonitorHDD(hdd HDD) {
	fmt.Println(fmt.Sprintf("HDD size: %d TB", hdd.Size))
}

func (r RAM) Monitor(monitoring Monitoring) {
	monitoring.MonitorRam(r)
}

func (c CPU) Monitor(monitoring Monitoring) {
	monitoring.MonitorCPU(c)
}

func (m Motherboard) Monitor(monitoring Monitoring) {
	monitoring.MonitorMotherboard(m)
}

func (m HDD) Monitor(monitoring Monitoring) {
	monitoring.MonitorHDD(m)
}