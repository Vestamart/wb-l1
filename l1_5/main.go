package main

import (
	"fmt"
	"time"
)

const N = 5

func sender(ch chan<- int, stop chan struct{}) {
	i := 0
	for {
		select {
		case <-stop:
			fmt.Println("Отправитель: Получен сигнал остановки, завершаю работу.")
			return
		default:
			ch <- i
			fmt.Printf("Отправитель: Отправлено значение %d\n", i)
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func receiver(ch <-chan int) {
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Получатель: Канал закрыт, завершаю работу.")
			return
		}
		fmt.Printf("Получатель: Получено значение %d\n", val)
	}
}

func main() {
	fmt.Printf("Программа будет работать в течение %d секунд.\n", N)

	dataCh := make(chan int)
	stopCh := make(chan struct{}, 1)

	go sender(dataCh, stopCh)
	go receiver(dataCh)

	select {
	case <-time.After(N * time.Second):
		fmt.Println("\nИстекло время работы")

		stopCh <- struct{}{}

		time.Sleep(100 * time.Millisecond)

		close(dataCh)

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Основная программа завершена.")
}
