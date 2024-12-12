package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Отправляю сообщение в канал...")
		ch <- 42
	}()
	go func() {
		msg := <-ch
		fmt.Println("Получено сообщение:", msg)
	}()
	time.Sleep(time.Second)
}
