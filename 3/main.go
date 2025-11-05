package main

import (
	"fmt"
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

	ch := make(chan int, n)
	

	for i := 1; i <= n; i++{
		go func() {
			for data := range ch{
				fmt.Printf("Горутина  %d: %d\n", i, data)
			}
		}()
	}

	for i := 0; ;i++{
		ch <- i
		time.Sleep(333 * time.Millisecond)
	}

}