package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	fmt.Println("Начало выполнения:", time.Now().Format("15:04:05.000"))

	sleep(2 * time.Second)

	fmt.Println("Продолжение выполнения после 2 секунд:", time.Now().Format("15:04:05.000"))

	sleep(500 * time.Millisecond)

	fmt.Println("Конец выполнения:", time.Now().Format("15:04:05.000"))
}
