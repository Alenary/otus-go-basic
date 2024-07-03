package main

import (
	"sync"
	"testing"
)

// TestSensorData тестирует функцию SensorData.
func TestSensorData(t *testing.T) {
	dataCh := make(chan float64, 60)
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

	if count != 60 {
		t.Errorf("ожидалось 60 данных, получено %d", count)
	}
}

// TestProcessData тестирует функцию ProcessData.
func TestProcessData(t *testing.T) {
	dataCh := make(chan float64, 10)
	processedCh := make(chan float64, 10)
	var wg sync.WaitGroup

	// Отправляем тестовые данные в dataCh.
	go func() {
		for i := 0; i < 60; i++ {
			dataCh <- float64(i)
		}
		close(dataCh)
	}()

	wg.Add(1)
	go ProcessData(dataCh, processedCh, &wg)

	// Ждем завершения горутины.
	wg.Wait()

	// Собираем обработанные данные.
	results := make([]float64, 0, 6) // Предварительное выделение памяти для среза results.
	for avg := range processedCh {
		results = append(results, avg)
	}

	// Ожидаемые средние значения для 0-9, 10-19, 20-29, 30-39, 40-49, 50-59.
	expected := []float64{4.5, 14.5, 24.5, 34.5, 44.5, 54.5}

	if len(results) != len(expected) {
		t.Fatalf("ожидалось %d результатов, получено %d", len(expected), len(results))
	}

	for i, avg := range results {
		if avg != expected[i] {
			t.Errorf("ожидалось среднее %.1f, получено %.1f", expected[i], avg)
		}
	}
}
