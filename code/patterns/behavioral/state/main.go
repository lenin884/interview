package main

/*
Рассмотрим пример управления вентилятором с различными режимами - выключен, низкая скорость,
средняя скорость и высокая скорость - используя паттерн "Состояние".
*/

import "fmt"

// State Интерфейс для состояния вентилятора
type State interface {
	HandleRequest(fan *Fan)
}

// Fan Контекст - вентилятор
type Fan struct {
	state State
}

func (f *Fan) Request() {
	f.state.HandleRequest(f)
}

func NewFan() *Fan {
	return &Fan{state: NewOffState()}
}

// OffState Конкретное состояние вентилятора - выключен
type OffState struct{}

func NewOffState() State {
	return &OffState{}
}

func (s *OffState) HandleRequest(fan *Fan) {
	fmt.Println("Переключение вентилятора на низкую скорость.")
	fan.state = NewLowState()
}

// LowState Конкретное состояние вентилятора - низкая скорость
type LowState struct{}

func NewLowState() State {
	return &LowState{}
}

func (s *LowState) HandleRequest(fan *Fan) {
	fmt.Println("Переключение вентилятора на среднюю скорость.")
	fan.state = NewMediumState()
}

// MediumState Конкретное состояние вентилятора - средняя скорость
type MediumState struct{}

func NewMediumState() State {
	return &MediumState{}
}

func (s *MediumState) HandleRequest(fan *Fan) {
	fmt.Println("Переключение вентилятора на высокую скорость.")
	fan.state = NewHighState()
}

// HighState Конкретное состояние вентилятора - высокая скорость
type HighState struct{}

func NewHighState() State {
	return &HighState{}
}

func (s *HighState) HandleRequest(fan *Fan) {
	fmt.Println("Выключение вентилятора.")
	fan.state = NewOffState()
}

func main() {
	fan := NewFan()

	fan.Request()
	fan.Request()
	fan.Request()
	fan.Request()
}

/*
В этом примере мы создаем состояния вентилятора (OffState, LowState, MediumState, HighState), каждое из
которых реализует интерфейс State с методом HandleRequest, который отвечает за изменение состояния вентилятора.

Контекст вентилятора (Fan) содержит текущее состояние. Когда вызывается метод Request, контекст делегирует запрос
текущему состоянию. Это позволяет динамически менять состояние вентилятора и его поведение, например,
переключать скорости или выключать.
*/
