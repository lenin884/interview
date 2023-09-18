package traveler

type Ticket struct {
	From string
	To   string
}

/*
findRoute метод для поиска маршрута
tickets - массив билетов.
Возвращает массив строк, где каждая строка - это пункт откуда путешественник пришел

Попробуем решить либо хеш таблицей
*/
func findRoute(tickets []Ticket) []string {
	// Создадим хеш таблицу
	hashTableFrom := make(map[string]string)
	hashTableTo := make(map[string]string)

	// Заполним хеш таблицу куда - откуда
	for _, ticket := range tickets {
		hashTableFrom[ticket.From] = ticket.To
		hashTableTo[ticket.To] = ticket.From
	}

	// Создадим массив для результата
	result := make([]string, len(tickets)+1)

	// Найдем конечную точку и пойдем по ней
	var last string
	for _, ticket := range tickets {
		if _, ok := hashTableFrom[ticket.To]; !ok {
			last = ticket.From
			break
		}
	}

	for i := len(tickets) - 1; i >= 0; i-- {
		result[i+1] = hashTableFrom[last]
		if _, ok := hashTableTo[last]; ok {
			last = hashTableTo[last]
		}
	}
	result[0] = last

	return result
}

type List struct {
	Val  string
	Next *List
	Prev *List
}

// findRouteFromAlexey метод для поиска маршрута с использование хеш таблицы и связного списка
func findRouteFromAlexey(tickets []Ticket) []string {
	// Создадим хеш таблицу
	hashMap := make(map[string]int, 0)
	hashInvertMap := make(map[string]string, 0)

	// Создадим связный список
	list := &List{}
	// Найдем точку начала и конца путешествия
	for _, ticket := range tickets {
		hashMap[ticket.From] -= 1
		hashMap[ticket.To] += 1
		hashInvertMap[ticket.From] = ticket.To
	}

	for _, ticket := range tickets {
		if hashMap[ticket.From] == -1 {
			list.Val = ticket.From
			list.Next = &List{Val: ticket.To, Prev: list}
			break
		}
	}

	for list.Next != nil {
		list = list.Next
		if hashMap[list.Val] == 1 {
			break
		}
		list.Next = &List{Val: hashInvertMap[list.Val], Prev: list}
	}

	result := make([]string, len(tickets)+1)

	index := len(tickets)
	for list != nil {
		result[index] = list.Val
		list = list.Prev
		index--
	}

	return result
}
