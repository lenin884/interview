package main

import "fmt"

// Абстрактные интерфейсы для фабрик

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

// Конкретные реализации интерфейсов для фабрик и продуктов

type WinFactory struct{}

func (w WinFactory) CreateButton() Button {
	return WinButton{}
}

func (w WinFactory) CreateCheckbox() Checkbox {
	return WinCheckbox{}
}

type WinButton struct{}

func (w WinButton) Paint() {
	fmt.Println("Отрисовка кнопки в стиле Windows")
}

type WinCheckbox struct{}

func (w WinCheckbox) Paint() {
	fmt.Println("Отрисовка флажка в стиле Windows")
}

type OSXFactory struct{}

func (o OSXFactory) CreateButton() Button {
	return OSXButton{}
}

func (o OSXFactory) CreateCheckbox() Checkbox {
	return OSXCheckbox{}
}

type OSXButton struct{}

func (o OSXButton) Paint() {
	fmt.Println("Отрисовка кнопки в стиле macOS")
}

type OSXCheckbox struct{}

func (o OSXCheckbox) Paint() {
	fmt.Println("Отрисовка флажка в стиле macOS")
}

// CreateGUI Функция, которая создает GUI элементы с помощью абстрактной фабрики
func CreateGUI(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.Paint()
	checkbox.Paint()
}

func main() {
	// Создание элементов в стиле Windows
	winFactory := WinFactory{}
	CreateGUI(winFactory)

	// Создание элементов в стиле macOS
	osxFactory := OSXFactory{}
	CreateGUI(osxFactory)
}

/*
В данном примере существуют две абстрактные фабрики: WinFactory и OSXFactory, каждая из которых реализует интерфейс
GUIFactory. Каждая фабрика имеет два метода для создания кнопок (CreateButton) и флажков (CreateCheckbox).

Конкретные реализации продуктов, такие как WinButton, WinCheckbox, OSXButton и OSXCheckbox, реализуют интерфейсы
Button и Checkbox для каждой из фабрик.

Функция CreateGUI позволяет создать GUI элементы с помощью абстрактной фабрики, и в зависимости от
выбранной фабрики будут созданы элементы в соответствующем стиле (Windows или macOS). Этот паттерн
позволяет легко поддерживать разные стили интерфейса без изменения кода, создающего эти элементы.
*/
