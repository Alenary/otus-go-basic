package main

import (
	"fmt"
)

func Board(size int) string {
	var board string
	if size < 2 {
		return "Указаны некорректные данные"
	}

	for i := 0; i < size; i++ {
		for n := 0; n < size; n++ {
			if (i+n)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	return board
}

func main() {
	var size int
	fmt.Println("Введите число - размер шахматной доски")
	fmt.Scanf("%d", &size)
	fmt.Print(Board(size))
}
