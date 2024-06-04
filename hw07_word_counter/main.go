package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Функция countWords принимает строку текста и возвращает мапу (ключ - слово, значение - количество упоминаний).
func countWords(text string) map[string]int {
	// Создание мапы для хранения количества упоминаний каждого слова.
	wordCount := make(map[string]int)

	// Преобразуем текст в нижний регистр.
	text = strings.ToLower(text)

	// Разделяем текст на слова с помощью strings.FieldsFunc.
	words := strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	// Подсчет упоминаний каждого слова.
	for _, word := range words {
		wordCount[word]++ // map типа map[string]int
	}

	return wordCount
}

func main() {
	// Ввод текста.
	fmt.Println("Введите текст:")

	// Ридер для чтения текста из стандартного ввода.
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text) // Убираем пробелы и символы новой строки по краям

	// Мапа с количеством упоминаний каждого слова.
	wordCount := countWords(text)

	// Вывод результата.
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
