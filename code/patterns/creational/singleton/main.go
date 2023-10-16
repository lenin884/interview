package main

import (
	"fmt"
	"sync"
)

// Singleton - структура для реализации Singleton
type Singleton struct {
	Data int
}

var instance *Singleton
var once sync.Once

// GetInstance - функция, возвращающая экземпляр Singleton
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Data: 42}
	})
	return instance
}

func main() {
	// Создание двух ссылок на Singleton
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	// Убедимся, что это один и тот же экземпляр
	if singleton1 == singleton2 {
		fmt.Println("Это один и тот же экземпляр Singleton")
		fmt.Printf("Значение Data: %d\n", singleton1.Data)
	} else {
		fmt.Println("Это разные экземпляры Singleton")
	}
}

/*
В этом примере мы создаем Singleton с помощью пакета sync. Мы определяем структуру Singleton, в которой может
храниться какие-либо данные (в данном случае, просто Data с значением 42).

Функция GetInstance создает экземпляр Singleton с использованием sync.Once, который гарантирует,
что создание произойдет только один раз. Если экземпляр уже создан, он просто возвращает существующий.

В функции main, мы получаем два экземпляра Singleton с помощью GetInstance и убеждаемся,
что они ссылаются на один и тот же объект.
*/
