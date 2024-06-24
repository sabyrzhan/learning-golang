package SRP

import (
	"fmt"
	"testing"
)

func TestSRP(t *testing.T) {
	car := NewCar("Toyota", "RAV4", 2.7)
	carService := NewServiceCenter()
	doesService := carService.DoesServiceMyCar(car.Make)
	fmt.Println(fmt.Sprintf("Car make and model: %s %s", car.Make, car.Model))
	fmt.Println(fmt.Sprintf("Does servie my car: %t", doesService))
	canBook := carService.CanBookForService(13, 5)
	fmt.Println(fmt.Sprintf("Can book at time=%d and day=%d : %t", 13, 5, canBook))
	carService.BookForService(car)
}
