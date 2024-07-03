package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

// SensorData симулирует считывание данных с сенсора.
func SensorData(dataCh chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 10; i++ { // Уменьшено количество итераций до 10
		<-ticker.C
		// Симуляция считывания данных с использованием crypto/rand.
		n, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			n = big.NewInt(0) // значение по умолчанию в случае ошибки.
		}
		data := float64(n.Int64())
		dataCh <- data
	}
	close(dataCh)
}

// ProcessData обрабатывает данные с сенсора и вычисляет средние значения.
func ProcessData(dataCh <-chan float64, processedCh chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	dataPoints := make([]float64, 0, 10)

	for data := range dataCh {
		dataPoints = append(dataPoints, data)
		if len(dataPoints) == 10 {
			// Вычисление среднего значения.
			sum := 0.0
			for _, val := range dataPoints {
				sum += val
			}
			average := sum / 10
			processedCh <- average
			dataPoints = dataPoints[:0] // сброс dataPoints.
		}
	}
	close(processedCh)
}

func main() {
	dataCh := make(chan float64, 10)
	processedCh := make(chan float64, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go SensorData(dataCh, &wg)

	wg.Add(1)
	go ProcessData(dataCh, processedCh, &wg)

	// WaitGroup для основной горутины, которая выводит обработанные данные.
	var printWg sync.WaitGroup
	printWg.Add(1)

	// Основная горутина для получения и вывода обработанных данных.
	go func() {
		defer printWg.Done()
		for avg := range processedCh {
			fmt.Printf("Обработанное среднее значение: %.2f\n", avg)
		}
	}()

	wg.Wait()
	printWg.Wait() // Ждем завершения вывода всех данных.
}
