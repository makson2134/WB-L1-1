package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func example1() {
	fmt.Println("1. Выход по условию:")
	var wg sync.WaitGroup
	stop := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if stop {
				fmt.Println("Горутина остановлена по флагу")
				return
			}
			fmt.Printf("Горутина выполняет итерацию: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(250 * time.Millisecond)
	stop = true
	wg.Wait()
	fmt.Println()
}

func example2() {
	fmt.Println("2. Остановка через канал уведомления:")
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Горутина остановлена через канал")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(300 * time.Millisecond)
	close(done)
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func example3() {
	fmt.Println("3. Остановка через контекст:")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена через контекст")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func example4() {
	fmt.Println("4. Остановка через runtime.Goexit():")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Горутина завершена через Goexit")

		for i := 0; i < 5; i++ {
			fmt.Printf("Горутина выполняет итерацию: %d\n", i)
			if i == 2 {
				runtime.Goexit()
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println()
}

func example5() {
	fmt.Println("5. Остановка через закрытие канала данных:")
	ch := make(chan int)

	go func() {
		for num := range ch {
			fmt.Printf("Горутина получила значение: %d\n", num)
		}
		fmt.Println("Горутина остановлена - канал закрыт")
	}()

	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}

	close(ch)
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
}
