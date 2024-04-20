package builder_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildAllFields(t *testing.T) {
	expected := `         CPU: Intel Core Ultra 9 185H 5.1
      Cooler: CoolerMaster Core Ultra 9
         RAM: 64GB
 Motherboard: MSI 185H Deluxe Edition
         GPU: NVidia GeForce 4090 RTX 24GB
         HDD: SSD Samsung 960 4000GB
        Case: Tower AeroCool Tower Black
Power Supply: AeroCool 1000X 1000W
     Monitor: Dell Wide Gamer Edition 49"
          OS: Windows 11 Pro
        Apps:
              Office 2021 Pro
              ESET NOD32 Home Edition
              Google Chrome
              Adobe Photoshop 2024
              DaVinci Resolve Studio Edition
              Topaz Video AI
              Embarcadero RAD Studio 12
`
	computerBuilder := NewComputerConfigurationBuilder()
	result := computerBuilder.
		NewCpuConfigurationBuilder().
		CpuName("Core Ultra 9 185H").
		CpuType("Intel").
		CpuPower(5.1).
		NewCoolerConfigurationBuilder().
		CoolerName("CoolerMaster Core Ultra 9").
		NewRAMConfigurationBuilder().
		RAMSize(64).
		NewMotherBoardConfigurationBuilder().
		MotherBoardModel("MSI 185H Deluxe Edition").
		NewGPUConfigurationBuilder().
		GPUType("NVidia").
		GpuName("GeForce 4090 RTX").
		GpuRamSize(24).
		NewHDDConfigurationBuilder().
		HDDType("SSD").
		HDDName("Samsung 960").
		HDDSize(4000).
		NewCaseConfigurationBuilder().
		CaseType("Tower").
		CaseName("AeroCool Tower").
		CaseColor("Black").
		NewPowerSupplyConfigurationBuilder().
		PowerSupplyPower(1000).
		PowerSupplyName("AeroCool 1000X").
		NewMonitorConfigurationBuilder().
		MonitorSize(49).
		MonitorName("Dell Wide Gamer Edition").
		NewOsConfigurationBuilder().
		OSName("Windows 11 Pro").
		AddApplication("Office 2021 Pro").
		AddApplication("ESET NOD32 Home Edition").
		AddApplication("Google Chrome").
		AddApplication("Adobe Photoshop 2024").
		AddApplication("DaVinci Resolve Studio Edition").
		AddApplication("Topaz Video AI").
		AddApplication("Embarcadero RAD Studio 12").
		Build()

	assert.Equal(t, expected, result)
}
