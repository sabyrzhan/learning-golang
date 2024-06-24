package SRP

import (
	"fmt"
	"slices"
)

// SRP states that object must have single responsibility and one reason to change
// Here I have a Car and Car service center
// I divided Car and Service center into different structures
// because of Separation of Concern
// If I combine two responsibilities into single struct, it would become a God object

type Motor struct {
	Volume float32
}

type Wheel struct {
	Size int
}

type Rim struct {
	Size int
}

type Car struct {
	Make  string
	Model string
	Motor Motor
	Wheel Wheel
	Rim   Rim
}

func (c *Car) ScanRecords() string {
	return `Year: 2019
Mileage: 15000 KM/h
Bluetooth: Exists
Wifi: Exists
AppleCarplay: Exists
AndroidAuto: Exists
`
}

func (c *Car) TurnToServiceMode(enabled bool) {
	if enabled {
		fmt.Println("Service mode enabled")
	} else {
		fmt.Println("Service mode disabled")
	}
}

func NewCar(make string, model string, motorVolume float32) Car {
	return Car{
		Make:  make,
		Model: model,
		Motor: Motor{motorVolume},
	}
}

type ServiceCenter struct {
	CarMakes           []string
	WorkingHours       []int
	WorkingDays        []int
	AvailableBookSeats int
}

func NewServiceCenter() ServiceCenter {
	return ServiceCenter{
		CarMakes:           []string{"Toyota", "Honda", "Mercedes-Benz"},
		WorkingHours:       []int{9, 18},
		WorkingDays:        []int{1, 5},
		AvailableBookSeats: 10,
	}
}

func (c *ServiceCenter) DoesServiceMyCar(make string) bool {
	return slices.Index(c.CarMakes, make) != -1
}

func (c *ServiceCenter) CanBookForService(hour int, day int) bool {
	if day < c.WorkingDays[0] || day > c.WorkingDays[1] {
		return false
	}

	if hour < c.WorkingHours[0] || hour > c.WorkingHours[1] {
		return false
	}

	return c.AvailableBookSeats > 0
}

func (c *ServiceCenter) BookForService(car Car) {
	fmt.Println(fmt.Sprintf("You car %s %s was booked", car.Make, car.Model))
}
