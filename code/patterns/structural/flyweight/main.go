package main

/*
Допустим, у вас есть приложение, которое работает с большим количеством символов текста,
и вы хотите оптимизировать использование памяти. Вы можете использовать паттерн "Легковес",
чтобы разделить общие символы и хранить их внутри легковесов.
Каждый символ будет представлен как объект, но общие символы будут разделяться между объектами.
*/

import "fmt"

// Flyweight - интерфейс для символов
type Flyweight interface {
	Display()
}

// ConcreteFlyweight - конкретная реализация символа
type ConcreteFlyweight struct {
	character rune
}

func (cf *ConcreteFlyweight) Display() {
	fmt.Printf("Символ: %c\n", cf.character)
}

// FlyweightFactory - фабрика легковесов
type FlyweightFactory struct {
	flyweights map[rune]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[rune]Flyweight),
	}
}

func (ff *FlyweightFactory) GetFlyweight(character rune) Flyweight {
	if flyweight, exists := ff.flyweights[character]; exists {
		return flyweight
	}

	// Если символ еще не существует, создаем новый легковес
	flyweight := &ConcreteFlyweight{character: character}
	ff.flyweights[character] = flyweight
	return flyweight
}

func main() {
	factory := NewFlyweightFactory()

	text := "Hello, World!"
	for _, char := range text {
		flyweight := factory.GetFlyweight(char)
		flyweight.Display()
	}
}

/*
В этом примере FlyweightFactory создает и управляет легковесами (конкретными символами).
Когда несколько символов из одного текста совпадают, они используют один и тот же легковес.
Это позволяет экономить память и уменьшать накладные расходы на хранение символов в памяти.
*/
