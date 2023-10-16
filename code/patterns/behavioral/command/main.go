package main

/*
Рассмотрим пример, где у нас есть электронное устройство (например, светильник), и мы хотим создать пульт
дистанционного управления, который будет отправлять команды для включения и выключения света.
*/

import "fmt"

// Light - Receiver - получатель команд
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Свет включен")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Свет выключен")
}

// Command - Command - интерфейс для команд
type Command interface {
	Execute()
}

// TurnOnLightCommand - ConcreteCommand - конкретная реализация команды
type TurnOnLightCommand struct {
	light *Light
}

func (c *TurnOnLightCommand) Execute() {
	c.light.TurnOn()
}

type TurnOffLightCommand struct {
	light *Light
}

func (c *TurnOffLightCommand) Execute() {
	c.light.TurnOff()
}

// RemoteControl - Invoker - инициатор команд
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}
	turnOnCommand := &TurnOnLightCommand{light: light}
	turnOffCommand := &TurnOffLightCommand{light: light}

	remote := &RemoteControl{}

	// Включение света
	remote.command = turnOnCommand
	remote.PressButton()

	// Выключение света
	remote.command = turnOffCommand
	remote.PressButton()
}

/*
В этом примере у нас есть получатель команды (Light), который представляет светильник, и две конкретные команды
(TurnOnLightCommand и TurnOffLightCommand), которые инкапсулируют действия включения и выключения света.

Пульт дистанционного управления (RemoteControl) представляет инициатор команд,
который может отправлять команды получателю. Пульт может быть настроен на выполнение разных команд.
*/
