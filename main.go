package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func main(){

	var wg sync.WaitGroup

	wg.Add(1)
		go func (){
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("First func call:", i)
			time.Sleep(time.Duration(time.Duration(rand.Intn(100)) * time.Millisecond))
		}
		}()

		wg.Add(1)
		go func (){
			defer wg.Done()
			for i := 0; i < 100; i++ {
				fmt.Println("Second func call:", i)
				time.Sleep(time.Duration(time.Duration(rand.Intn(100)) * time.Millisecond))
			}
			}()

			wg.Wait()

	fmt.Println("All done")
}

