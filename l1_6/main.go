package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// --- 1. Завершение по условию (Естественный выход) ---
func workerByCondition(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		if i >= 3 {
			fmt.Println("Worker 1 (Condition): Условие выполнено (i=3), выход из горутины.")
			return
		}
		fmt.Printf("Worker 1 (Condition): Работает, итерация %d\n", i)
		time.Sleep(300 * time.Millisecond)
		i++
	}
}

// --- 2. Завершение через канал уведомления (Signal Channel) ---
func workerByChannel(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			fmt.Println("Worker 2 (Channel): Получен сигнал остановки через канал, выход.")
			return
		default:
			fmt.Println("Worker 2 (Channel): Работает...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// --- 3. Завершение через context.WithCancel (Идиоматический способ) ---
func workerByContextCancel(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker 3 (Context Cancel): Получен сигнал отмены. Причина: %v, выход.\n", ctx.Err())
			return
		default:
			fmt.Println("Worker 3 (Context Cancel): Работает...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// --- 4. Завершение через context.WithTimeout ---
func workerByContextTimeout(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker 4 (Context Timeout): Сработал таймаут. Причина: %v, выход.\n", ctx.Err())
			return
		default:
			fmt.Println("Worker 4 (Context Timeout): Работает, ждет таймаута...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// --- 5. Прекращение работы runtime.Goexit() ---
// Goexit() заставляет текущую горутину немедленно завершиться, выполняя все defer'ы.
func workerWithGoexit(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Worker 5 (Goexit): [DEFERRED] Отложенный вызов сработал.")

	fmt.Println("Worker 5 (Goexit): Шаг 1: Начинаем работу.")
	time.Sleep(300 * time.Millisecond)

	fmt.Println("Worker 5 (Goexit): Шаг 2: Вызываем runtime.Goexit().")
	runtime.Goexit()

}

// --- 6. Завершение через общую атомарную переменную (Shared Flag) ---
// Используется atomic для безопасного изменения флага без мьютекса.
var stopFlag int32

func workerByAtomicFlag(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if atomic.LoadInt32(&stopFlag) == 1 {
			fmt.Println("Worker 6 (Atomic Flag): Обнаружен флаг остановки, выход.")
			return
		}

		fmt.Println("Worker 6 (Atomic Flag): Работает, проверяет флаг...")
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	delay := 500 * time.Millisecond

	fmt.Println("\n1. Завершение по условию")
	wg.Add(1)
	go workerByCondition(&wg)
	wg.Wait()

	fmt.Println("\n2. Завершение через канал уведомления")
	stopCh := make(chan struct{})
	wg.Add(1)
	go workerByChannel(stopCh, &wg)

	time.Sleep(delay)
	fmt.Println("Main: Отправка сигнала остановки в канал.")
	close(stopCh)
	wg.Wait()

	fmt.Println("\n3. Завершение через context.WithCancel")
	ctxCancel, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go workerByContextCancel(ctxCancel, &wg)

	time.Sleep(delay)
	fmt.Println("Main: Вызов функции отмены (cancel()).")
	cancel()
	wg.Wait()

	fmt.Println("\n4. Завершение через context.WithTimeout")
	timeout := 1200 * time.Millisecond
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), timeout)
	defer cancelTimeout()
	wg.Add(1)
	go workerByContextTimeout(ctxTimeout, &wg)

	wg.Wait()
	fmt.Println("\n5. Прекращение работы runtime.Goexit()")
	wg.Add(1)
	go workerWithGoexit(&wg)
	wg.Wait()

	fmt.Println("\n6. Завершение через общую атомарную переменную")
	atomic.StoreInt32(&stopFlag, 0)
	wg.Add(1)
	go workerByAtomicFlag(&wg)

	time.Sleep(delay)
	fmt.Println("Main: Установка атомарного флага остановки.")
	atomic.StoreInt32(&stopFlag, 1)

	wg.Wait()

	fmt.Println("\nВсе горутины успешно завершены.")
}
