package main

import (
	"fmt"
	"sync"
)

var givenArray = [5]int{2, 4, 6, 8, 10}

func squareFunc(n int, ch chan int, wg *sync.WaitGroup){
	ch <- n*n 
	defer wg.Done()
}


func main(){
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(5)

	for _, v := range givenArray{
		go squareFunc(v, ch, &wg)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for num := range ch{
		fmt.Println(num)
	} 

}
