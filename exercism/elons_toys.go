package elon

import "fmt"

func main() {
	/*distance := 100
	track := NewTrack(distance)*/

	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)

	//car = Drive(car)
	//car := NewCar(speed, batteryDrain)
	//fmt.Println("car is now %s", car.Drive())
	car.Drive()

	fmt.Printf("Speed %d, BatteryDrain %d, Battery %d, Distance %d\n", car.speed, car.batteryDrain, car.battery, car.distance)

	//canFinish := CanFinish(car, track)
	//fmt.Printf("\nCarrera finalizada: %t\n", canFinish)
	fmt.Println(car.DisplayDistance()) // 0??
	fmt.Println(car.DisplayBattery())  // 100%

	c := Car{speed: 5, batteryDrain: 5, battery: 100}
	fmt.Println(c.CanFinish(101))
	//car.CanFinish(track)

}

// TODO: define the 'Car' type struct
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
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

// CanFinish checks if a car is able to finish a certain track.
/*func CanFinish(car Car, track Track) bool {
	if (car.battery/car.batteryDrain)*car.speed >= track.distance {
		return true
	}
	return false
}*/

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery = c.battery - c.batteryDrain
		c.distance += c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	return (c.battery/c.batteryDrain)*c.speed >= trackDistance
}
