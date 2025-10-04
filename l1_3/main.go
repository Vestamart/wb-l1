package l1_3

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Воркер %d обработал: %d\n", id, j)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Printf("Воркер %d завершил работу.\n", id)
}

func main() {
	numWorkers := 3
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil && n > 0 {
			numWorkers = n
		}
	}
	fmt.Printf("Запустили %d воркеров\n", numWorkers)

	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	go func() {
		fmt.Println("Начали запись данных в канал...")
		for i := 1; ; i++ {
			jobs <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	fmt.Println("\nПолучили сигнал завершения")

	close(jobs)
	fmt.Println("Канал заданий закрыт.")

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершается.")
}
