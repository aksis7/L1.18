package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

// Инкрементируем счетчик
func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Получаем текущее значение счетчика
func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	// Количество горутин и инкрементов
	numGoroutines := 100
	numIncrements := 1000

	// Запускаем 100 горутин
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
}
