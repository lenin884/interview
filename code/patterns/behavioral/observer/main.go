package main

/*
Давайте создадим простой пример, где есть издатель (субъект), который публикует новости, и подписчики (наблюдатели),
которые получают уведомления о новых новостях.
*/

import "fmt"

// Observer Интерфейс для наблюдателя
type Observer interface {
	Update(news string)
}

// Subject Интерфейс для субъекта
type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers(news string)
	PublishNews(news string)
}

// Publisher Издатель (субъект)
type Publisher struct {
	observers []Observer
}

func NewPublisher() Subject {
	return &Publisher{}
}

func (p *Publisher) RegisterObserver(observer Observer) {
	p.observers = append(p.observers, observer)
}

func (p *Publisher) RemoveObserver(observer Observer) {
	for i, o := range p.observers {
		if o == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *Publisher) NotifyObservers(news string) {
	for _, observer := range p.observers {
		observer.Update(news)
	}
}

func (p *Publisher) PublishNews(news string) {
	fmt.Printf("Новая новость: %s\n", news)
	p.NotifyObservers(news)
}

// Subscriber Подписчик (наблюдатель)
type Subscriber struct {
	name string
}

func NewSubscriber(name string) Observer {
	return &Subscriber{name: name}
}

func (s *Subscriber) Update(news string) {
	fmt.Printf("%s получил новость: %s\n", s.name, news)
}

func main() {
	publisher := NewPublisher()

	subscriber1 := NewSubscriber("Подписчик 1")
	subscriber2 := NewSubscriber("Подписчик 2")

	publisher.RegisterObserver(subscriber1)
	publisher.RegisterObserver(subscriber2)

	publisher.PublishNews("Важное событие!")

	publisher.RemoveObserver(subscriber1)

	publisher.PublishNews("Еще одна новость!")
}

/*
В этом примере Publisher представляет субъект, который может регистрировать, удалять и уведомлять наблюдателей о новых
новостях. Subscriber представляет наблюдателя, который получает новости и реагирует на них.

Когда новость публикуется с помощью PublishNews, все зарегистрированные наблюдатели уведомляются о новости.
Если наблюдатель больше не интересуется новостями, его можно удалить из списка наблюдателей с помощью RemoveObserver.
*/
