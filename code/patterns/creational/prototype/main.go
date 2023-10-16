package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeColor string
}

func (p *Person) Clone() *Person {
	// Создаем копию текущего объекта
	clone := *p
	return &clone
}

func main() {
	// Создаем прототип человека
	prototype := &Person{
		Name:     "John",
		Age:      30,
		EyeColor: "Brown",
	}

	// Создаем нового человека на основе прототипа
	person1 := prototype.Clone()
	person1.Name = "Alice"

	// Создаем еще одного нового человека на основе прототипа
	person2 := prototype.Clone()
	person2.Name = "Bob"
	person2.Age = 25

	// Вывод информации о созданных людях
	fmt.Printf("Person 1: Name - %s, Age - %d, Eye Color - %s\n", person1.Name, person1.Age, person1.EyeColor)
	fmt.Printf("Person 2: Name - %s, Age - %d, Eye Color - %s\n", person2.Name, person2.Age, person2.EyeColor)
}

/*
В этом примере Person представляет объект, который мы хотим клонировать. Мы добавили метод Clone(), который создает
глубокую копию текущего объекта.

Затем мы создаем прототип человека prototype и создаем два новых человека (person1 и person2) на основе этого прототипа.
Мы изменяем их свойства, но они все равно остаются независимыми объектами. Это демонстрирует принцип работы паттерна
"Прототип" - создание новых объектов путем клонирования существующих, чтобы избежать дополнительной сложности создания
объектов с нуля.
*/
