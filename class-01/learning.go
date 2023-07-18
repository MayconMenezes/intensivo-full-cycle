package learning

type Car struct {
	Model string
	Color string
}

// sum function
func sum(a int, b int) int {
	return a + b
}

// method
func (c Car) Start() { // (c Car) is a receiver
	println("Starting the car", c.Model)
}

func (c Car) ChangeColor(color string) {
	c.Color = color
	println("Changing the car color to", c.Color)
}

func (c *Car) ChangeColorPointer(color string) {
	c.Color = color
	println("Changing the car color to", c.Color)
}

func learning() {
	car := Car{ // := is a short variable declaration and assignment
		Model: "Ferrari",
		Color: "Red",
	}
	println(car.Model, car.Color, sum(1, 2))

	car.Start()
	car.ChangeColor("Blue")
	println("Old color is", car.Color, "without pointer method")
	car.ChangeColorPointer("Purple")
	println("New color is", car.Color)

	a := 10
	b := a // it copied the value of a to b, but b created its one space in memory
	b = 20
}
