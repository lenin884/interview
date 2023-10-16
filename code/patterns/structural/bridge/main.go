package main

/*
Рассмотрим ситуацию, где у нас есть формы и системы окраски для этих форм.
Мы хотим, чтобы каждая форма могла быть нарисована с использованием разных систем окраски.
Мы можем использовать паттерн "Мост" для разделения абстракции формы и реализации системы окраски.
*/

import "fmt"

// ColorImplementor Интерфейс для системы окраски
type ColorImplementor interface {
	Colorize() string
}

// RedColor Конкретная реализация системы окраски
type RedColor struct{}

func (r *RedColor) Colorize() string {
	return "красным цветом"
}

type GreenColor struct{}

func (g *GreenColor) Colorize() string {
	return "зеленым цветом"
}

// Shape Интерфейс для формы
type Shape interface {
	Draw() string
	SetColor(ColorImplementor)
}

// Circle Конкретная реализация формы
type Circle struct {
	color ColorImplementor
}

func (c *Circle) SetColor(color ColorImplementor) {
	c.color = color
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Рисуем круг %s", c.color.Colorize())
}

func main() {
	redColor := &RedColor{}
	greenColor := &GreenColor{}

	circle := &Circle{}

	// Рисуем круг красным цветом
	circle.SetColor(redColor)
	fmt.Println(circle.Draw())

	// Рисуем круг зеленым цветом
	circle.SetColor(greenColor)
	fmt.Println(circle.Draw())
}

/*
В этом примере мы разделили интерфейс Shape на абстракцию и конкретную реализацию формы.
Затем мы внедрили интерфейс ColorImplementor для системы окраски, позволяя форме иметь разные системы окраски.
Circle - это конкретная реализация формы, и мы можем динамически изменять систему окраски для этой формы, не меняя код формы.
Это демонстрирует принцип паттерна "Мост" - разделение абстракции и реализации для уменьшения связанности и повышения гибкости кода.
*/
