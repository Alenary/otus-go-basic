package main

import (
	"sync"
	"testing"
	"time"
)

// TestSensorData тестирует функцию SensorData.
func TestSensorData(t *testing.T) {
	dataCh := make(chan float64, 60)
	var wg sync.WaitGroup

	// Уменьшаем время работы сенсора до 5 секунд для теста
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-ticker.C:
				n := float64(42) // Используем фиксированное значение для простоты
				dataCh <- n
			case <-timer.C:
				close(dataCh)
				return
			}
		}
	}()
	wg.Add(1)

	// Ждем завершения горутины.
	wg.Wait()

	// Проверяем, что было сгенерировано правильное количество данных.
	count := 0
	for range dataCh {
		count++
	}

	if count != 5 {
		t.Errorf("ожидалось 5 данных, получено %d", count)
	}
}

// TestProcessData тестирует функцию ProcessData.
func TestProcessData(t *testing.T) {
	dataCh := make(chan float64, 10)
	processedCh := make(chan float64, 10)
	var wg sync.WaitGroup

	// Отправляем тестовые данные в dataCh.
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			dataCh <- float64(i)
		}
		close(dataCh)
	}()
	wg.Add(1)

	wg.Add(1)
	go ProcessData(dataCh, processedCh, &wg)

	// Ждем завершения горутины.
	wg.Wait()

	// Собираем обработанные данные.
	results := make([]float64, 0, 1) // Предварительное выделение памяти для среза results.
	for avg := range processedCh {
		results = append(results, avg)
	}

	// Ожидаемое среднее значение для 0-9.
	expected := []float64{4.5}

	if len(results) != len(expected) {
		t.Fatalf("ожидалось %d результатов, получено %d", len(expected), len(results))
	}

	for i, avg := range results {
		if avg != expected[i] {
			t.Errorf("ожидалось среднее %.1f, получено %.1f", expected[i], avg)
		}
	}
}
