package main

import (
	"fmt"
)

func main() {
	var size int
	var board string
	fmt.Println("Введите число - размер шахматной доски")
	fmt.Scanf("%d", &size)
	if size < 2 {
		fmt.Println("Указаны некорректные данные")
	} else {
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
		fmt.Println(board)
	}
}
