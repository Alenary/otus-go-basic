package main

import (
	"fmt"
	"sync"
)

// incrementCounter увеличивает значение счетчика с использованием Mutex для синхронизации.
func incrementCounter(mu *sync.Mutex, counter *int) {
	// Синхронизируем доступ к счетчику.
	// Блокировка позволяет одной горутине получить эксклюзивный доступ к ресурсу,
	// чтобы предотвратить одновременные изменения, которые могут привести к некорректным данным или гонкам данных.
	mu.Lock()
	*counter++
	mu.Unlock()
}

// runGoroutines запускает указанное количество горутин, каждая из которых увеличивает счетчик.
func runGoroutines(numGoroutines int, wg *sync.WaitGroup, mu *sync.Mutex, counter *int) {
	// Устанавливаем значение счетчика WaitGroup.
	wg.Add(numGoroutines)

	// Запускаем горутины.
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			// Метод Done() уменьшает внутренний счетчик WaitGroup на 1.
			// Этот метод вызывается внутри каждой горутины, когда она завершает свою работу.
			defer wg.Done()

			// Синхронизируем доступ к счетчику.
			incrementCounter(mu, counter)
			fmt.Printf("Горутина %d увеличила счетчик, текущее значение: %d\n", id, *counter)
		}(i)
	}

	// Ожидает, пока счетчик не станет равным нулю. Ждем завершения всех горутин.
	wg.Wait()
}

// Функция main - точка входа в программу.
func main() {
	var counter int

	// Создаем WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Создаем Mutex для синхронизации доступа к счетчику.
	var mu sync.Mutex

	// Определяем количество горутин.
	numGoroutines := 5

	// Запускаем горутины и ждем их завершения.
	runGoroutines(numGoroutines, &wg, &mu, &counter)

	// Выводим окончательное значение счетчика.
	fmt.Printf("Окончательное значение счетчика: %d\n", counter)
}
