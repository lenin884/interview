package main

/*
Давайте создадим пример, где у нас есть несколько обработчиков, каждый из которых может проверить и обработать запрос,
и если необходимо, передать его следующему обработчику в цепи. Мы будем рассматривать запрос на одобрение кредита.
*/

import "fmt"

// LoanHandler Обработчик запроса на кредит
type LoanHandler interface {
	SetNext(LoanHandler)
	HandleLoanRequest(float64)
}

// BaseLoanHandler Базовый обработчик
type BaseLoanHandler struct {
	next LoanHandler
}

func (b *BaseLoanHandler) SetNext(handler LoanHandler) {
	b.next = handler
}

// SmallLoanHandler Обработчик маленьких кредитов
type SmallLoanHandler struct {
	BaseLoanHandler
}

func (s *SmallLoanHandler) HandleLoanRequest(amount float64) {
	if amount <= 10000 {
		fmt.Printf("Маленький кредит на сумму %.2f одобрен\n", amount)
	} else if s.next != nil {
		s.next.HandleLoanRequest(amount)
	}
}

// MediumLoanHandler Обработчик средних кредитов
type MediumLoanHandler struct {
	BaseLoanHandler
}

func (m *MediumLoanHandler) HandleLoanRequest(amount float64) {
	if amount <= 50000 {
		fmt.Printf("Средний кредит на сумму %.2f одобрен\n", amount)
	} else if m.next != nil {
		m.next.HandleLoanRequest(amount)
	}
}

// LargeLoanHandler Обработчик больших кредитов
type LargeLoanHandler struct {
	BaseLoanHandler
}

func (l *LargeLoanHandler) HandleLoanRequest(amount float64) {
	if amount <= 100000 {
		fmt.Printf("Большой кредит на сумму %.2f одобрен\n", amount)
	} else {
		fmt.Printf("Кредит на сумму %.2f не может быть одобрен\n", amount)
	}
}

func main() {
	// Создаем цепь обработчиков
	smallLoanHandler := &SmallLoanHandler{}
	mediumLoanHandler := &MediumLoanHandler{}
	largeLoanHandler := &LargeLoanHandler{}

	smallLoanHandler.SetNext(mediumLoanHandler)
	mediumLoanHandler.SetNext(largeLoanHandler)

	// Тестируем обработку запросов на кредит
	loanRequests := []float64{5000, 25000, 75000, 150000}
	for _, request := range loanRequests {
		fmt.Println("---------------------------")
		smallLoanHandler.HandleLoanRequest(request)
	}
	fmt.Println("---------------------------")
}

/*
В этом примере мы создали цепь обработчиков для запросов на кредит разных размеров.
Каждый обработчик проверяет, может ли он одобрить кредит, и если не может, передает запрос следующему обработчику
в цепи. Это позволяет нам гибко настраивать порядок обработки запросов и избегать жесткой привязки
отправителя запроса к его получателю.
*/
