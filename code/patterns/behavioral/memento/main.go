package main

import "fmt"

// Memento - интерфейс снимка
type Memento interface {
	GetState() string
}

// ConcreteMemento - конкретная реализация снимка
type ConcreteMemento struct {
	state string
}

func (cm *ConcreteMemento) GetState() string {
	return cm.state
}

// Editor - Originator - создатель, объект, состояние которого можно сохранять и восстанавливать
type Editor struct {
	content string
}

func (e *Editor) Write(text string) {
	e.content += text
}

func (e *Editor) Display() {
	fmt.Println("Содержимое редактора:", e.content)
}

func (e *Editor) Save() Memento {
	return &ConcreteMemento{state: e.content}
}

func (e *Editor) Restore(memento Memento) {
	e.content = memento.GetState()
}

// History - Caretaker - управляющий, который сохраняет и восстанавливает состояние редактора
type History struct {
	mementos []Memento
}

func (h *History) Push(memento Memento) {
	h.mementos = append(h.mementos, memento)
}

func (h *History) Pop() Memento {
	if len(h.mementos) == 0 {
		return nil
	}
	memento := h.mementos[len(h.mementos)-1]
	h.mementos = h.mementos[:len(h.mementos)-1]
	return memento
}

func main() {
	editor := &Editor{}
	history := &History{}

	editor.Write("Это начальное состояние. ")
	history.Push(editor.Save())

	editor.Write("Добавим некоторый текст. ")
	history.Push(editor.Save())

	editor.Write("Продолжим редактирование. ")
	editor.Display()

	fmt.Println("Отменяем последнее действие.")
	editor.Restore(history.Pop())
	editor.Display()

	fmt.Println("Отменяем еще одно действие.")
	editor.Restore(history.Pop())
	editor.Display()
}

/*
В этом примере Editor представляет объект, состояние которого можно сохранять и восстанавливать.
History представляет объект, который отвечает за сохранение и восстановление снимков состояния редактора.
Снимки состояния редактора сохраняются в ConcreteMemento.

Мы добавляем текст в редактор, сохраняем состояние после каждой операции и можем отменять операции,
восстанавливая предыдущие состояния.
*/
