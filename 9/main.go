package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)
func main(){
	var  n int
	fmt.Print("Enter numbers count: ")
	fmt.Scan(&n)
	

	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	//Generation
	wg.Add(1)
	go func(){
		defer wg.Done()
		defer close(input)
		for i := 0; i < n; i ++{
			num := rand.Intn(250)
			input <- num
			fmt.Printf("[Generation] --- Goroutine set %d in first chanel\n", num)
		}
	}()

	//Calculating
	wg.Add(1)
	go func(){
		defer wg.Done()
		defer close (output)
		for inputNum := range input{
			outputNum := inputNum * 2
			output <- outputNum
			fmt.Println("[Calculating] --- Goroutine set new calculated num in second chanel")
		}
	}()

	//Writing in StdOut
	wg.Add(1)
	go func(){
		defer wg.Done()
		for result := range output{
			fmt.Fprintf(os.Stdout, "[Output] --- Goroutine wrote %d in StdOut\n", result)
		}
	}()
	

	wg.Wait()
}