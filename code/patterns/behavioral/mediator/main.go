package main

/*
Рассмотрим пример простой системы чата, где несколько пользователей могут отправлять
сообщения друг другу с помощью медиатора.
*/

import "fmt"

// Mediator - интерфейс медиатора
type Mediator interface {
	RegisterUser(user *User)
	SendMessage(sender *User, message string)
}

// User - пользователь системы
type User struct {
	name     string
	mediator Mediator
}

func NewUser(name string, mediator Mediator) *User {
	return &User{name: name, mediator: mediator}
}

func (u *User) Send(message string) {
	u.mediator.SendMessage(u, message)
}

func (u *User) ReceiveMessage(message string) {
	fmt.Printf("%s получил сообщение: %s\n", u.name, message)
}

// ChatMediator - конкретный медиатор
type ChatMediator struct {
	users []*User
}

func NewChatMediator() Mediator {
	return &ChatMediator{}
}

func (cm *ChatMediator) RegisterUser(user *User) {
	cm.users = append(cm.users, user)
}

func (cm *ChatMediator) SendMessage(sender *User, message string) {
	for _, user := range cm.users {
		if user != sender {
			user.ReceiveMessage(message)
		}
	}
}

func main() {
	mediator := NewChatMediator()

	user1 := NewUser("Пользователь 1", mediator)
	user2 := NewUser("Пользователь 2", mediator)
	user3 := NewUser("Пользователь 3", mediator)

	mediator.RegisterUser(user1)
	mediator.RegisterUser(user2)
	mediator.RegisterUser(user3)

	user1.Send("Привет, всем!")
	user2.Send("Привет!")
}

/*
В этом примере у нас есть медиатор ChatMediator, который регистрирует пользователей и обеспечивает коммуникацию между
ними. Каждый пользователь (User) может отправлять сообщения с помощью метода Send, и
медиатор рассылает сообщения всем остальным пользователям. Компоненты (пользователи) не зависят
напрямую друг от друга, а взаимодействуют через медиатор.
*/
