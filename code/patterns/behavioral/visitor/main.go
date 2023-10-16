package main

/*
Рассмотрим пример, где у нас есть структура данных для представления различных элементов, таких как числа и строки,
и мы хотим выполнить операцию над каждым элементом с помощью посетителя.
*/

import "fmt"

// Интерфейс элемента, который принимает посетителя
type Element interface {
	Accept(visitor Visitor)
}

// Конкретные элементы
type IntegerElement struct {
	Value int
}

func (ie *IntegerElement) Accept(visitor Visitor) {
	visitor.VisitIntegerElement(ie)
}

type StringElement struct {
	Value string
}

func (se *StringElement) Accept(visitor Visitor) {
	visitor.VisitStringElement(se)
}

// Интерфейс посетителя
type Visitor interface {
	VisitIntegerElement(element *IntegerElement)
	VisitStringElement(element *StringElement)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitIntegerElement(element *IntegerElement) {
	fmt.Printf("Посетитель обработал целое число: %d\n", element.Value)
}

func (cv *ConcreteVisitor) VisitStringElement(element *StringElement) {
	fmt.Printf("Посетитель обработал строку: %s\n", element.Value)
}

func main() {
	elements := []Element{
		&IntegerElement{Value: 42},
		&StringElement{Value: "Hello, World!"},
	}

	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}

/*
В этом примере у нас есть элементы IntegerElement и StringElement, которые реализуют интерфейс Element и могут
принимать посетителя с помощью метода Accept. У нас также есть интерфейс Visitor,
который объявляет методы для каждого типа элемента, и конкретный посетитель ConcreteVisitor,
который реализует эти методы.

При выполнении программы мы создаем элементы и передаем их посетителю. Посетитель выполняет
соответствующие операции над каждым элементом, не изменяя их классы.
*/
