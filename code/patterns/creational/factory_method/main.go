package main

type ShapeFactory interface {
	CreateShape() Shape
}

type Shape interface {
	Draw()
}

type Rectangle struct{}

type Circle struct{}

type RectangleFactory struct{}

type CircleFactory struct{}

func (r *Rectangle) Draw() {
	println("draw rectangle")
}

func (c *Circle) Draw() {
	println("draw circle")
}

func (r *RectangleFactory) CreateShape() Shape {
	return &Rectangle{}
}

func (c *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

func main() {
	var factory ShapeFactory
	factory = &RectangleFactory{}
	shape := factory.CreateShape()
	shape.Draw()

	factory = &CircleFactory{}
	shape = factory.CreateShape()
	shape.Draw()
}

/*
В этом примере есть интерфейс ShapeFactory, который предоставляет метод CreateShape для создания объектов типа Shape.
Конкретные реализации CircleFactory и SquareFactory реализуют этот интерфейс и создают объекты Circle и Square,
соответственно. Каждый из этих объектов также реализует интерфейс Shape и имеет свой собственный метод Draw,
который определяет, как рисовать соответствующую фигуру.
*/
