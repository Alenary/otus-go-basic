package main

import (
	"fmt"
	"sync"
)

// incrementCounter увеличивает значение счетчика с использованием Mutex для синхронизации.
func incrementCounter(mu *sync.Mutex, counter *int) {
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
			mu.Lock()
			fmt.Printf("Горутина %d увеличила счетчик, текущее значение: %d\n", id, *counter)
			mu.Unlock()
		}(i)
	}

	// Ожидает, пока счетчик не станет равным нулю. Ждем завершения всех горутин.
	wg.Wait()
}

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
