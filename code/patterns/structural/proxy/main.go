package main

/*
Допустим, у вас есть интерфейс Image и его реализация RealImage, которая представляет загрузку и отображение
изображений. Мы можем создать прокси-объект ProxyImage, который будет загружать и
отображать изображение только при необходимости.
*/

import "fmt"

// Image - интерфейс для работы с изображением
type Image interface {
	Display()
}

// RealImage - реальная реализация изображения
type RealImage struct {
	filename string
}

func (ri *RealImage) Display() {
	fmt.Println("Отображение изображения:", ri.filename)
}

func NewRealImage(filename string) *RealImage {
	fmt.Println("Загрузка изображения:", filename)
	return &RealImage{filename: filename}
}

// ProxyImage - прокси для изображения
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func (pi *ProxyImage) Display() {
	if pi.realImage == nil {
		pi.realImage = NewRealImage(pi.filename)
	}
	pi.realImage.Display()
}

func NewProxyImage(filename string) Image {
	return &ProxyImage{filename: filename}
}

func main() {
	// Создание прокси для изображения
	image := NewProxyImage("image.jpg")

	// Изображение отображается только при вызове Display
	image.Display()

	// Изображение не загружается повторно
	image.Display()
}

/*
В этом примере RealImage представляет реальную реализацию загрузки и отображения изображений.
ProxyImage - это прокси-класс, который откладывает создание реального объекта RealImage до тех пор, пока вызов
Display() не будет сделан. Таким образом, изображение загружается и отображается только при необходимости.
*/
