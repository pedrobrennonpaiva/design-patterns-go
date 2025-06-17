package factory_method

import "fmt"

func Run() {
	car, _ := getCar("Prius")
	car2, _ := getCar("Corolla")

	fmt.Println("Car 1: ", car.GetName())
	fmt.Println("Car 1: ", car.GetColor())
	fmt.Println("Car 1: ", car.GetEngine())

	fmt.Println("Car 2: ", car2.GetName())
	fmt.Println("Car 2: ", car2.GetColor())
	fmt.Println("Car 2: ", car2.GetEngine())
}

func getCar(name string) (ICar, error) {
	switch name {
	case "Prius":
		return newPrius(), nil
	case "Corolla":
		return newCorolla(), nil
	}

	return nil, fmt.Errorf("invalid car: %s", name)
}

type ICar interface {
	setName(name string)
	setColor(color string)
	setEngine(engine string)
	GetName() string
	GetColor() string
	GetEngine() string
}

type Car struct {
	name   string
	color  string
	engine string
}

func (c *Car) setName(name string) {
	c.name = name
}

func (c *Car) setColor(color string) {
	c.color = color
}

func (c *Car) setEngine(engine string) {
	c.engine = engine
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetColor() string {
	return c.color
}

func (c *Car) GetEngine() string {
	return c.engine
}

// Prius
type Prius struct {
	Car
}

func newPrius() ICar {
	return &Prius{
		Car: Car{
			name:   "Prius",
			color:  "Red",
			engine: "Electric",
		},
	}
}

// Corolla
type Corolla struct {
	Car
}

func newCorolla() ICar {
	return &Corolla{
		Car: Car{
			name:   "Corolla",
			color:  "Blue",
			engine: "Gasoline",
		},
	}
}
