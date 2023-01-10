package speed

import "fmt"

func main() {
	distance := 100
	track := NewTrack(distance)

	speed := 5
	batteryDrain := 3
	car := NewCar(speed, batteryDrain)

	car = Drive(car)

	fmt.Printf("Speed %d, BatteryDrain %d, Battery %d, Distance %d\n", car.speed, car.batteryDrain, car.battery, car.distance)

	canFinish := CanFinish(car, track)
	fmt.Printf("\nCarrera finalizada: %t", canFinish)
}

// TODO: define the 'Car' type struct
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery == 0 || car.batteryDrain > car.battery {
		car.distance = car.distance
	} else {
		car.battery = car.battery - car.batteryDrain
		car.distance += car.speed
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if (car.battery/car.batteryDrain)*car.speed >= track.distance {
		return true
	}
	return false
}
