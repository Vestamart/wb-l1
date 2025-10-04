package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int) {
	fmt.Printf("Worker %d запущен.\n", id)
	defer fmt.Printf("Worker %d завершен.\n", id)

	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Worker %d: Делаю работу...\n", id)
		case <-ctx.Done():
			fmt.Printf("Worker %d: Получен сигнал отмены\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i)
	}

	fmt.Println("Программа запущена")

	<-sigCh

	fmt.Println("\nПолучен сигнал прерывания")
	cancel()

	fmt.Println("Ожидаю завершения воркеров...")
	time.Sleep(2 * time.Second)

	fmt.Println("Все воркеры завершены. ")
}
