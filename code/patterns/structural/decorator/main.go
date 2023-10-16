package main

import "fmt"

/*
Представьте, у вас есть интерфейс Coffee и его реализация SimpleCoffee, которая представляет обычный кофе
*/

type Coffee interface {
	Price() int
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) Price() int {
	return 5
}

/*
Теперь мы хотим добавить возможность украсить этот кофе молоком и/или сахаром, не изменяя исходный класс.
Мы можем использовать паттерн "Декоратор" для этой цели
*/

type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) Price() int {
	return m.coffee.Price() + 2
}

type SugarDecorator struct {
	coffee Coffee
}

func (s *SugarDecorator) Price() int {
	return s.coffee.Price() + 1
}

/*
Теперь мы можем создать разные комбинации кофе и декораторов
*/
func main() {
	coffee := &SimpleCoffee{}
	fmt.Printf("Простой кофе: $%d\n", coffee.Price())

	milkCoffee := &MilkDecorator{coffee: coffee}
	fmt.Printf("Кофе с молоком: $%d\n", milkCoffee.Price())

	sweetCoffee := &SugarDecorator{coffee: coffee}
	fmt.Printf("Кофе с сахаром: $%d\n", sweetCoffee.Price())

	milkAndSugarCoffee := &MilkDecorator{coffee: sweetCoffee}
	fmt.Printf("Кофе с молоком и сахаром: $%d\n", milkAndSugarCoffee.Price())
}

/*
В этом примере мы создаем декораторы MilkDecorator и SugarDecorator, каждый из которых реализует интерфейс Coffee,
добавляя соответственно молоко и сахар к базовому кофе. Мы можем комбинировать декораторы для создания разных
видов кофе с разной функциональностью.
*/
