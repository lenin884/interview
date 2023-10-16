package main

/*
Давайте представим, что у нас есть задача по созданию структуры, представляющей файловую систему. Файловая система может
содержать как файлы, так и папки. Мы можем использовать паттерн "Компоновщик" для создания иерархии файлов и папок.
*/

import "fmt"

// Component - интерфейс для всех компонентов (файлов и папок)
type Component interface {
	Name() string
	Size() int
}

// File - конкретная реализация файлов
type File struct {
	name string
	size int
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

// Folder - конкретная реализация папок
type Folder struct {
	name     string
	contents []Component
}

func (f *Folder) Name() string {
	return f.name
}

func (f *Folder) Size() int {
	totalSize := 0
	for _, c := range f.contents {
		totalSize += c.Size()
	}
	return totalSize
}

func main() {
	// Создание файлов
	file1 := &File{name: "file1.txt", size: 10}
	file2 := &File{name: "file2.txt", size: 20}

	// Создание папок и добавление файлов
	folder1 := &Folder{name: "Folder 1", contents: []Component{file1, file2}}
	folder2 := &Folder{name: "Folder 2", contents: []Component{file1, folder1}}

	// Вычисление размера папок
	fmt.Printf("Размер %s: %d байтов\n", folder1.Name(), folder1.Size())
	fmt.Printf("Размер %s: %d байтов\n", folder2.Name(), folder2.Size())
}

/*
В этом примере Component представляет общий интерфейс для файлов и папок. File и Folder - конкретные реализации.
Файлы имеют имена и размеры, а папки содержат список компонентов. Папки вычисляют размер,
суммируя размеры всех своих компонентов. Это позволяет нам строить иерархическую структуру
файловой системы и работать с ней, используя единый интерфейс.
*/
