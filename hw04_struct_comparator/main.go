package main

import "fmt"

//  Cтруктура Book с неэкспортируемыми полями: ID, Title, Author, Year, Size, Rate (может быть дробным).

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// Конструктор создания новой книги с возвратом указателя на этот экземпляр.

func NewBook(id int, title, author string, year, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

// Методы для установки (Set) и получения (Get) полей структуры.

// Pointer receiver.

func (b *Book) SetID(id int) {
	b.id = id
}

// Value receiver.

func (b *Book) ID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) Rate() float64 {
	return b.rate
}

// BookCompare Определение пользовательского типа для сравнения.
type BookCompare int

// Константы для определения режимов сравнения книг по году, размеру и рейтингу.
const (
	ByYear BookCompare = iota
	BySize
	ByRate
)

// BookComparator Структура для сравнения книг по разным полям.
type BookComparator struct {
	a BookCompare
}

// NewBookComparator - создание нового экземпляра структуры BookComparator.
func NewBookComparator(mode BookCompare) *BookComparator {
	return &BookComparator{a: mode}
}

// Compare - метод для сравнения двух книг.
func (bc *BookComparator) Compare(book1, book2 *Book) bool {
	switch bc.a {
	case ByYear:
		return book1.Year() > book2.Year()
	case BySize:
		return book1.Size() > book2.Size()
	case ByRate:
		return book1.Rate() > book2.Rate()
	default:
		fmt.Println("Второй аргумент больше первого")
		return false
	}
}

func main() {
	book1 := NewBook(1, "Title1", "Author1", 2024, 100, 4.9)
	book2 := NewBook(2, "Title2", "Author2", 2023, 99, 4.5)

	comparatorByYear := NewBookComparator(ByYear)
	fmt.Println("Сравнение по году:", comparatorByYear.Compare(book1, book2))

	comparatorBySize := NewBookComparator(BySize)
	fmt.Println("Сравнение по размеру:", comparatorBySize.Compare(book1, book2))

	comparatorByRate := NewBookComparator(ByRate)
	fmt.Println("Сравнение по рейтингу:", comparatorByRate.Compare(book1, book2))
}
