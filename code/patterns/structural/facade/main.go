package main

import "fmt"

/*
Допустим, у вас есть сложная система, предоставляющая несколько сервисов, таких как отправка сообщений,
обработка платежей и управление пользователями. Каждый сервис имеет свой уникальный интерфейс и методы.
Вы можете создать фасад, чтобы упростить работу с этой системой
*/

// MessagingService Сервис отправки сообщений
type MessagingService struct{}

func (ms *MessagingService) SendMessage(message string) {
	fmt.Println("Отправка сообщения:", message)
}

// PaymentService Сервис обработки платежей
type PaymentService struct{}

func (ps *PaymentService) ProcessPayment(amount float64) {
	fmt.Println("Обработка платежа на сумму:", amount)
}

// UserService Сервис управления пользователями
type UserService struct{}

func (us *UserService) RegisterUser(username string) {
	fmt.Println("Регистрация пользователя:", username)
}

// Facade Фасад
type Facade struct {
	messagingService *MessagingService
	paymentService   *PaymentService
	userService      *UserService
}

func NewFacade() *Facade {
	return &Facade{
		messagingService: &MessagingService{},
		paymentService:   &PaymentService{},
		userService:      &UserService{},
	}
}

func (f *Facade) PerformPurchase(username, message string, amount float64) {
	f.userService.RegisterUser(username)
	f.messagingService.SendMessage(message)
	f.paymentService.ProcessPayment(amount)
}

func main() {
	facade := NewFacade()

	username := "user123"
	message := "Спасибо за покупку!"
	amount := 50.0

	fmt.Println("Начало покупки.")
	facade.PerformPurchase(username, message, amount)
	fmt.Println("Покупка завершена.")
}

/*
В этом примере создан фасад Facade, который инкапсулирует сложную систему из трех разных сервисов.
Фасад предоставляет единственный метод PerformPurchase, который выполняет ряд операций,
связанных с регистрацией пользователя, отправкой сообщения и обработкой платежа.
*/
