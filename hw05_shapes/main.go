package main

import (
	"errors"
	"fmt"
	"math"
)

// Интерфейс Shape с методом для вычисления площади.

type Shape interface {
	Area() float64
}

// Структуры для геометрических фигур: круг, прямоугольник, треугольник.

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Base, Height float64
}

// calculateArea функция для вычисления площади фигуры ожидает на входе объект типа Shape и возвращает его площадь.

func calculateArea(s Shape) (float64, error) {
	if s == nil {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return s.Area(), nil
}

// Area метод для вычисления площади круга.

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area метод для вычисления площади прямоугольника.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area метод для вычисления площади треугольника.

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Реализация метода String для Circle.
func (c Circle) String() string {
	return fmt.Sprintf("Круг: радиус %.0f", c.Radius)
}

// Реализация метода String для Rectangle.
func (r Rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %.0f, высота %.0f", r.Width, r.Height)
}

// Реализация метода String для Triangle.
func (t Triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %.0f, высота %.0f", t.Base, t.Height)
}

func main() {
	var shapes []Shape

	// Ввод данных для круга

	var radius float64
	fmt.Print("Введите радиус круга: ")
	_, err := fmt.Scan(&radius)
	if err != nil || radius <= 0 {
		fmt.Println("Ошибка ввода: переданный объект не является фигурой.", err)
		return
	}

	// Создание объекта Circle с введенным радиусом.

	circle := Circle{Radius: radius}

	// Добавление объекта Circle в срез shapes.

	shapes = append(shapes, circle)

	// Ввод данных для прямоугольника

	var width, height float64
	fmt.Print("Введите ширину прямоугольника: ")
	_, err = fmt.Scan(&width)
	if err != nil || width <= 0 {
		fmt.Println("Ошибка ввода: переданный объект не является фигурой.", err)
		return
	}
	fmt.Print("Введите высоту прямоугольника: ")
	_, err = fmt.Scan(&height)
	if err != nil || height <= 0 {
		fmt.Println("Ошибка ввода: переданный объект не является фигурой.", err)
		return
	}
	rectangle := Rectangle{Width: width, Height: height}
	shapes = append(shapes, rectangle)

	// Ввод данных для треугольника

	var base, triangleHeight float64
	fmt.Print("Введите основание треугольника: ")
	_, err = fmt.Scan(&base)
	if err != nil || base <= 0 {
		fmt.Println("Ошибка ввода: переданный объект не является фигурой.", err)
		return
	}
	fmt.Print("Введите высоту треугольника: ")
	_, err = fmt.Scan(&triangleHeight)
	if err != nil || triangleHeight <= 0 {
		fmt.Println("Ошибка ввода: переданный объект не является фигурой.", err)
		return
	}

	triangle := Triangle{Base: base, Height: triangleHeight}
	shapes = append(shapes, triangle)

	// Вызов функции calculateArea для каждой фигуры и вывод результата
	for _, shape := range shapes {
		area, calcErr := calculateArea(shape)
		if calcErr != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			// Использование разного форматирования для площади
			if _, ok := shape.(Circle); ok {
				fmt.Printf("%s Площадь: %.14f\n", shape, area)
			} else {
				fmt.Printf("%s Площадь: %.0f\n", shape, area)
			}
		}
	}
}
