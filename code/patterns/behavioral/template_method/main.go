package main

/*
Давайте рассмотрим пример реализации паттерна "Метод шаблона" для создания чайных и кофейных напитков.
У нас будет абстрактный класс Beverage, который определяет общий шаблон для приготовления напитков,
и два конкретных подкласса Tea и Coffee, которые реализуют специфические шаги приготовления каждого напитка.
*/

import "fmt"

// Beverage Абстрактный класс
type Beverage interface {
	PrepareBeverage()
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

// Tea Конкретный класс: Чай
type Tea struct{}

// NewTea Конструктор
func NewTea() Beverage {
	return &Tea{}
}

func (t *Tea) PrepareBeverage() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}

func (t *Tea) BoilWater() {
	fmt.Println("- Кипятим воду")
}

func (t *Tea) Brew() {
	fmt.Println("- Завариваем чай")
}

func (t *Tea) PourInCup() {
	fmt.Println("- Переливаем чай в чашку")
}

func (t *Tea) AddCondiments() {
	fmt.Println("- Добавляем лимон")
}

// Coffee Конкретный класс: Кофе
type Coffee struct{}

// NewCoffee Конструктор
func NewCoffee() Beverage {
	return &Coffee{}
}

func (c *Coffee) PrepareBeverage() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

func (c *Coffee) BoilWater() {
	fmt.Println("- Кипятим воду")
}

func (c *Coffee) Brew() {
	fmt.Println("- Завариваем кофе")
}

func (c *Coffee) PourInCup() {
	fmt.Println("- Переливаем кофе в чашку")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("- Добавляем сахар и молоко")
}

func main() {
	tea := NewTea()
	coffee := NewCoffee()

	fmt.Println("Приготовление чая:")
	tea.PrepareBeverage()

	fmt.Println("\nПриготовление кофе:")
	coffee.PrepareBeverage()
}

/*
В этом примере Beverage представляет абстрактный класс с методами, которые представляют шаги процесса приготовления
напитка. Tea и Coffee - это конкретные подклассы, которые реализуют специфические детали приготовления чая и кофе.

Общий шаблон приготовления напитков остается неизменным, но каждый подкласс может переопределить методы Brew и
AddCondiments для определения собственных шагов.
*/
