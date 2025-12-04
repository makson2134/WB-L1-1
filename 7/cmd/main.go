package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/makson2134/WB-L1-7/safemap"
)

func main() {
	sm := safemap.NewSafeMap()
	var n, m int

	fmt.Printf("Num of goroutines: ")
	fmt.Scanf("%d", &n)

	fmt.Printf("Num of digitals to write: ")
	fmt.Scanf("%d", &m)

	jobs := make(chan int, m)
	for i := 0; i < m; i++ {
		jobs <- i + 1
	}
	close(jobs)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for num := range jobs {
				sm.SetValue(num, num)
				fmt.Printf("Goroutine %d wrote %d in safe map\n", id, num)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("|----------------------Now reading----------------------|")
	time.Sleep(3 * time.Second)

	keys := make(chan int, m)
	for i := 0; i < m; i++ {
		keys <- i + 1
	}
	close(keys)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for key := range keys {
				num, ok := sm.GetValue(key)
				if !ok {
					fmt.Printf("Error: gourutine %d couldn't get value with key: %d\n", id, num)
				} else {
					fmt.Printf("Goroutine %d read %d from safe map\n", id, num)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("|--------------------------------END OF THE PROGRAMM---------------------------------|")

}
