package main

/*
Допустим, у вас есть список задач (to-do list), и вы хотите создать итератор для перебора элементов этого списка.
*/

import "fmt"

// Iterator - интерфейс для итератора
type Iterator interface {
	HasNext() bool
	Next() string
}

// TodoList - коллекция задач
type TodoList struct {
	tasks []string
}

func NewTodoList() *TodoList {
	return &TodoList{}
}

func (tl *TodoList) AddTask(task string) {
	tl.tasks = append(tl.tasks, task)
}

// TodoListIterator - конкретная реализация итератора
type TodoListIterator struct {
	todoList *TodoList
	index    int
}

func NewTodoListIterator(todoList *TodoList) Iterator {
	return &TodoListIterator{todoList: todoList, index: 0}
}

func (tli *TodoListIterator) HasNext() bool {
	return tli.index < len(tli.todoList.tasks)
}

func (tli *TodoListIterator) Next() string {
	if tli.HasNext() {
		task := tli.todoList.tasks[tli.index]
		tli.index++
		return task
	}
	return ""
}

func main() {
	todoList := NewTodoList()
	todoList.AddTask("Посадить цветы")
	todoList.AddTask("Купить продукты")
	todoList.AddTask("Прочитать книгу")

	iterator := NewTodoListIterator(todoList)

	fmt.Println("Список задач:")
	for iterator.HasNext() {
		task := iterator.Next()
		fmt.Println(task)
	}
}

/*
В этом примере у нас есть TodoList, представляющий коллекцию задач. TodoListIterator -
это конкретная реализация итератора для этой коллекции.

Мы добавляем задачи в TodoList и затем используем TodoListIterator для итерации по списку задач.
Итератор позволяет нам перебирать элементы коллекции без необходимости знать её внутреннее представление.
*/
