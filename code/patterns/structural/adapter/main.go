package main

import "fmt"

/*
Предположим, у нас есть структура LegacyPrinter с методом PrintLegacy.
*/

type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintLegacy(s string) {
	fmt.Println("Legacy Printer: " + s)
}

/*
И у нас есть структура ModernPrinter с методом PrintModern.
*/

type ModernPrinter struct{}

func (mp *ModernPrinter) PrintModern(s string) {
	fmt.Println("Modern Printer: " + s)
}

/*
Теперь, предположим, что у нас есть клиентский код, который ожидает работу с ModernPrinter, но у нас есть только
LegacyPrinter. Мы можем создать адаптер, который позволит LegacyPrinter работать с интерфейсом ModernPrinter.
*/

type PrinterAdapter struct {
	OldPrinter *LegacyPrinter
	Msg        string
}

func (pa *PrinterAdapter) PrintModern(s string) {
	pa.OldPrinter.PrintLegacy(pa.Msg)
}

/*
Теперь мы можем использовать PrinterAdapter, чтобы сделать LegacyPrinter
совместимым с интерфейсом ModernPrinter.
*/

func main() {
	oldPrinter := &LegacyPrinter{}
	adapter := &PrinterAdapter{OldPrinter: oldPrinter, Msg: "Hello from Legacy Printer"}

	modernPrinter := &ModernPrinter{}

	modernPrinter.PrintModern("Hello from Modern Printer")
	adapter.PrintModern("Hello from Adapter")
}

/*
Это позволяет нам использовать LegacyPrinter как ModernPrinter, благодаря адаптеру. Adapter - это прослойка
между клиентским кодом и объектом, который не соответствует интерфейсу, ожидаемому клиентом.
Это упрощает интеграцию старых компонентов в новые системы или работу с библиотеками, имеющими разные интерфейсы.
*/
