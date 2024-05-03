package main

import "fmt"

func main() {
	var size int
	fmt.Println("Введите число - размер шахматной доски")
	fmt.Scanf("%d", &size)
	if size <2 {
		fmt.Println("Указаны некорректные данные")
	} else {
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				createchessline1(size)
			} else {
				createchessline2(size)
			}
		}
	}
}

func createchessline1(size int) {
	for i := 0; i < size; i++ {
		if i < size-1 {
			if i%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		} else {
			if i%2 == 0 {
				fmt.Println("#")
			} else {
				fmt.Println(" ")
			}
		}
	}
}

func createchessline2(size int) {
	for i := 0; i < size; i++ {
		if i < size-1 {
			if i%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		} else {
			if i%2 == 0 {
				fmt.Println(" ")
			} else {
				fmt.Println("#")
			}
		}
	}
}