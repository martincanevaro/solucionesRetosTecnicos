package main

import "fmt"

type Car struct {
	make  string
	model string
	color string
	doors int
}

type CarBuilder interface {
	SetMake(make string)
	SetModel(model string)
	SetColor(color string)
	SetDoors(doors int)
	Build() Car
}

type BasicCarBuilder struct {
	car *Car
}

func (b *BasicCarBuilder) SetMake(make string) {
	b.car.make = make
}

func (b *BasicCarBuilder) SetModel(model string) {
	b.car.model = model
}

func (b *BasicCarBuilder) SetColor(color string) {
	b.car.color = color
}

func (b *BasicCarBuilder) SetDoors(doors int) {
	b.car.doors = doors
}

func (b *BasicCarBuilder) Build() Car {
	return *b.car
}

func main() {
	builder := &BasicCarBuilder{
		car: &Car{
			make:  "Toyota",
			model: "Corolla",
			color: "Red",
			doors: 4,
		},
	}

	builder.SetMake("Honda")
	builder.SetModel("Civic")
	builder.SetColor("Blue")
	builder.SetDoors(2)

	car := builder.Build()

	fmt.Println(car)
}
