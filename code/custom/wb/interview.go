package wb

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Представь, что у тебя есть функция, работающая неопределенно долго (условный сетевой запрос)
// Нужно к ней написать функцию обертку, которая будет ограничивать её время работы в 1 секунду
// Если уложились - возвращаем результат функции, если нет - ошибка
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Изменить
func predictable() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ch := make(chan int64)
	go func() {
		ch <- unpredictableFunc()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return 0, errors.New("timeout")
	case v := <-ch:
		return v, nil
	}
}

//--------------------------------------------------------------------------------

func ReadFromChan() {
	c := make(chan int)
	go func() {
		for {
			c <- 1
		}
	}()
	time.Sleep(time.Millisecond * 100)
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}

//--------------------------------------------------------------------------------

func GetValue() (x int) {
	defer func(n int) {
		fmt.Printf("x as param: %d\n", n)
		fmt.Printf("x after return: %d\n", x)
		x *= 10
	}(x)

	x = 7
	return x
}
