package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitChannels() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("A")

		// wg.Add(-1)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("B")
	}()

	wg.Wait()
	fmt.Println("C")
}
