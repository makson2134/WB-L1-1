package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	var n int
	fmt.Println("Количество горутин: ")
	_, err := fmt.Scan(&n)

	if err != nil{ 
		fmt.Println("Ошибка ввода", err)
		return
	}

	//Я забыл добавить это в предыдущем задании
	if n < 1{
		fmt.Println("Количество горутин должно быть больше 0")
		return
	}

	ch := make(chan int, n)
	done := make(chan struct{})

	for i := 1; i <= n; i++ {
    go func() {
        for {
            select {
            case data := <-ch:
                fmt.Printf("Горутина %d: %d\n", i, data)
            case <-done: 
				fmt.Printf("Горутина %d закончила работу\n", i)
                return
            }
        }
    }()
}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
	for i := 0; ;i++{
			select{
			case ch <- i:
				time.Sleep(333 * time.Millisecond)
			case <- done:
				return
			}
		}
	}()

	<-sigChan
	fmt.Println("\nИзящное завершение...")

	close(done)
	time.Sleep(1 * time.Second)
	close(ch)
	fmt.Println("Работа завершена")
}