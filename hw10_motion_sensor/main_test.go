package main

import (
	"sync"
	"testing"
)

// TestSensorData тестирует функцию SensorData.
func TestSensorData(t *testing.T) {
	dataCh := make(chan float64, 10) // Уменьшено количество данных до 10
	var wg sync.WaitGroup

	wg.Add(1)
	go SensorData(dataCh, &wg)

	// Ждем завершения горутины.
	wg.Wait()

	// Проверяем, что было сгенерировано правильное количество данных.
	count := 0
	for range dataCh {
		count++
	}

	if count != 10 { // Ожидание 10 данных вместо 60
		t.Errorf("ожидалось 10 данных, получено %d", count)
	}
}

// TestProcessData тестирует функцию ProcessData.
func TestProcessData(t *testing.T) {
	dataCh := make(chan float64, 10)
	processedCh := make(chan float64, 10)
	var wg sync.WaitGroup

	// Отправляем тестовые данные в dataCh.
	go func() {
		for i := 0; i < 10; i++ {
			dataCh <- float64(i)
		}
		close(dataCh)
	}()

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
