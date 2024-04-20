package builder_pattern

import (
	"fmt"
	"strings"
)

type ComputerConfiguration struct {
	// CPU
	cpuName  string
	cpuType  string
	cpuPower float32

	// Cooler
	coolerName string

	//RAM
	ramSize int

	// Motherboard
	motherBoardModel string

	// GPU
	gpuType    string
	gpuName    string
	gpuRamSize int

	//HDD
	hddSize int
	hddName string
	hddType string

	// Case
	caseType  string
	caseName  string
	caseColor string

	//Power
	powerSupplyPower int
	powerSupplyName  string

	// Monitor
	monitorSize int
	monitorName string

	// OS
	osName       string
	applications []string
}

type ComputerConfigurationBuilder struct {
	computerConfiguration *ComputerConfiguration
}

func (b *ComputerConfigurationBuilder) NewCpuConfigurationBuilder() *CpuConfigurationBuilder {
	return &CpuConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewCoolerConfigurationBuilder() *CoolerConfigurationBuilder {
	return &CoolerConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewRAMConfigurationBuilder() *RAMConfigurationBuilder {
	return &RAMConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewMotherBoardConfigurationBuilder() *MotherboardConfigurationBuilder {
	return &MotherboardConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewGPUConfigurationBuilder() *GPUConfigurationBuilder {
	return &GPUConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewHDDConfigurationBuilder() *HDDConfigurationBuilder {
	return &HDDConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewCaseConfigurationBuilder() *CaseConfigurationBuilder {
	return &CaseConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewPowerSupplyConfigurationBuilder() *PowerSupplyConfigurationBuilder {
	return &PowerSupplyConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewMonitorConfigurationBuilder() *MonitorConfigurationBuilder {
	return &MonitorConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) NewOsConfigurationBuilder() *OSConfigurationBuilder {
	return &OSConfigurationBuilder{b}
}

func (b *ComputerConfigurationBuilder) Build() string {
	c := b.computerConfiguration
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("%13s %s %s %.1f\n", "CPU:", c.cpuType, c.cpuName, c.cpuPower))
	result.WriteString(fmt.Sprintf("%13s %s\n", "Cooler:", c.coolerName))
	result.WriteString(fmt.Sprintf("%13s %dGB\n", "RAM:", c.ramSize))
	result.WriteString(fmt.Sprintf("%13s %s\n", "Motherboard:", c.motherBoardModel))
	result.WriteString(fmt.Sprintf("%13s %s %s %dGB\n", "GPU:", c.gpuType, c.gpuName, c.gpuRamSize))
	result.WriteString(fmt.Sprintf("%13s %s %s %dGB\n", "HDD:", c.hddType, c.hddName, c.hddSize))
	result.WriteString(fmt.Sprintf("%13s %s %s %s\n", "Case:", c.caseType, c.caseName, c.caseColor))
	result.WriteString(fmt.Sprintf("%13s %s %dW\n", "Power Supply:", c.powerSupplyName, c.powerSupplyPower))
	result.WriteString(fmt.Sprintf("%13s %s %d\"\n", "Monitor:", c.monitorName, c.monitorSize))
	result.WriteString(fmt.Sprintf("%13s %s\n", "OS:", c.osName))
	result.WriteString(fmt.Sprintf("%13s\n", "Apps:"))
	for _, app := range c.applications {
		result.WriteString(fmt.Sprintf("%13s %s\n", " ", app))
	}

	return result.String()
}

func NewComputerConfigurationBuilder() *ComputerConfigurationBuilder {
	return &ComputerConfigurationBuilder{&ComputerConfiguration{}}
}

// CpuConfigurationBuilder setting all  CPU related fields
type CpuConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *CpuConfigurationBuilder) CpuName(name string) *CpuConfigurationBuilder {
	b.computerConfiguration.cpuName = name
	return b
}

func (b *CpuConfigurationBuilder) CpuType(cpuType string) *CpuConfigurationBuilder {
	b.computerConfiguration.cpuType = cpuType
	return b
}

func (b *CpuConfigurationBuilder) CpuPower(power float32) *CpuConfigurationBuilder {
	b.computerConfiguration.cpuPower = power
	return b
}

// CoolerConfigurationBuilder set cooler related fields
type CoolerConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *CoolerConfigurationBuilder) CoolerName(name string) *CoolerConfigurationBuilder {
	b.computerConfiguration.coolerName = name
	return b
}

// RAMConfigurationBuilder set ram parameters
type RAMConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *RAMConfigurationBuilder) RAMSize(ramSize int) *RAMConfigurationBuilder {
	b.computerConfiguration.ramSize = ramSize
	return b
}

// MotherboardConfigurationBuilder set mb model
type MotherboardConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *MotherboardConfigurationBuilder) MotherBoardModel(motherBoardModel string) *MotherboardConfigurationBuilder {
	b.computerConfiguration.motherBoardModel = motherBoardModel
	return b
}

// GPUConfigurationBuilder set GPU fields
type GPUConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *GPUConfigurationBuilder) GPUType(gpuType string) *GPUConfigurationBuilder {
	b.computerConfiguration.gpuType = gpuType
	return b
}

func (b *GPUConfigurationBuilder) GpuName(name string) *GPUConfigurationBuilder {
	b.computerConfiguration.gpuName = name
	return b
}

func (b *GPUConfigurationBuilder) GpuRamSize(ramSize int) *GPUConfigurationBuilder {
	b.computerConfiguration.gpuRamSize = ramSize
	return b
}

// HDDConfigurationBuilder set HDD fields
type HDDConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *HDDConfigurationBuilder) HDDSize(hddSize int) *HDDConfigurationBuilder {
	b.computerConfiguration.hddSize = hddSize
	return b
}

func (b *HDDConfigurationBuilder) HDDName(name string) *HDDConfigurationBuilder {
	b.computerConfiguration.hddName = name
	return b
}

func (b *HDDConfigurationBuilder) HDDType(hddType string) *HDDConfigurationBuilder {
	b.computerConfiguration.hddType = hddType
	return b
}

// CaseConfigurationBuilder set case fields
type CaseConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *CaseConfigurationBuilder) CaseType(caseType string) *CaseConfigurationBuilder {
	b.computerConfiguration.caseType = caseType
	return b
}

func (b *CaseConfigurationBuilder) CaseName(name string) *CaseConfigurationBuilder {
	b.computerConfiguration.caseName = name
	return b
}

func (b *CaseConfigurationBuilder) CaseColor(caseColor string) *CaseConfigurationBuilder {
	b.computerConfiguration.caseColor = caseColor
	return b
}

// PowerSupplyConfigurationBuilder set power supply fields
type PowerSupplyConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *PowerSupplyConfigurationBuilder) PowerSupplyPower(power int) *PowerSupplyConfigurationBuilder {
	b.computerConfiguration.powerSupplyPower = power
	return b
}

func (b *PowerSupplyConfigurationBuilder) PowerSupplyName(name string) *PowerSupplyConfigurationBuilder {
	b.computerConfiguration.powerSupplyName = name
	return b
}

// MonitorConfigurationBuilder set monitor fields
type MonitorConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *MonitorConfigurationBuilder) MonitorSize(size int) *MonitorConfigurationBuilder {
	b.computerConfiguration.monitorSize = size
	return b
}

func (b *MonitorConfigurationBuilder) MonitorName(name string) *MonitorConfigurationBuilder {
	b.computerConfiguration.monitorName = name
	return b
}

// OSConfigurationBuilder set os and apps fields
type OSConfigurationBuilder struct {
	*ComputerConfigurationBuilder
}

func (b *OSConfigurationBuilder) OSName(name string) *OSConfigurationBuilder {
	b.computerConfiguration.osName = name
	return b
}

func (b *OSConfigurationBuilder) AddApplication(appName string) *OSConfigurationBuilder {
	b.computerConfiguration.applications = append(b.computerConfiguration.applications, appName)
	return b
}
