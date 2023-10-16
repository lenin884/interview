package main

import "fmt"

/*
Допустим, у нас есть задача создать сложный объект "Компьютер" с разными компонентами,
такими как процессор, оперативная память,
жесткий диск и так далее, используя паттерн "Строитель".
*/

// Computer Продукт - компьютер
type Computer struct {
	Processor string
	RAM       int
	HDD       int
}

// ComputerBuilder Интерфейс строителя
type ComputerBuilder interface {
	SetProcessor()
	SetRAM()
	SetHDD()
	GetComputer() Computer
}

// GamingComputerBuilder Конкретный строитель
type GamingComputerBuilder struct {
	computer Computer
}

func (b *GamingComputerBuilder) SetProcessor() {
	b.computer.Processor = "Intel i9"
}

func (b *GamingComputerBuilder) SetRAM() {
	b.computer.RAM = 32
}

func (b *GamingComputerBuilder) SetHDD() {
	b.computer.HDD = 1000
}

func (b *GamingComputerBuilder) GetComputer() Computer {
	return b.computer
}

// ComputerDirector Директор
type ComputerDirector struct {
	builder ComputerBuilder
}

func (d *ComputerDirector) Construct() {
	d.builder.SetProcessor()
	d.builder.SetRAM()
	d.builder.SetHDD()
}

func main() {
	// Создание директора и конкретного строителя
	director := ComputerDirector{}
	builder := &GamingComputerBuilder{}

	// Собираем игровой компьютер
	director.builder = builder
	director.Construct()
	computer := builder.GetComputer()

	// Вывод информации о компьютере
	fmt.Printf("Игровой компьютер:\nПроцессор: %s\nRAM: %d ГБ\nHDD: %d ГБ\n", computer.Processor, computer.RAM, computer.HDD)
}

/*
В этом примере ComputerDirector управляет процессом пошагового конструирования компьютера с использованием конкретного
строителя GamingComputerBuilder. Каждый конкретный строитель реализует методы для установки различных компонентов
компьютера, и по завершении конструирования, мы получаем готовый объект "Компьютер" с нужными характеристиками.
*/
