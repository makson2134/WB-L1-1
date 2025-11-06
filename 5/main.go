package main

import (
	"context"
	"fmt"
	"time"
)

func sendNums(ch chan int, ctx context.Context, i int){
	for{
		select{
		case <- ctx.Done():
			fmt.Println("Функция отправки закончила работу")
			return
		default:
			time.Sleep(100 * time.Millisecond)
			ch <- i
			i++
		}
		
	}
}

func getNums(ch chan int, ctx context.Context){
	for{
		select{
		case <- ctx.Done():
			fmt.Println("Функция считывания закончила работу")
			return
		case num := <- ch:
			fmt.Printf("Функция получила значение: %d;\n", num)
		}
		
	}
}

func main(){
	var n int
	fmt.Println("N: ")
	fmt.Scan(&n)

	ch := make(chan int, 1)

	timeout := time.Duration(n) * time.Second
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	go sendNums(ch, ctx, 0)
	go getNums(ch, ctx)
	
	<-ctx.Done()

	time.Sleep(1 * time.Second)


	
}
