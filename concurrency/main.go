package main

import (
	"fmt"
	"sync"
)

func main() {
	// asyncHelloWorld()
	// asyncHelloWorld2()

	bufferedChannel()
	//thead safe
	// mutexExample()
}

func asyncHelloWorld() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
	fmt.Println("world")
}

//with anonymous function for a cleaner hello function
func asyncHelloWorld2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		hello2()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("world")
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}

func hello2() {
	fmt.Println("hello")
}

//without mutex the final value of c is not warantied
//comment out the mutex code and see it yourself
func mutexExample() {
	c := 0
	n := 200
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			c++
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(c)
}

func bufferedChannel() {
	c := make(chan string, 2)
	//this call is blocking unless you have a bufeered channel
	c <- "hello"
	c <- "world"

	// simple alternative to receive messages
	// msg := <-c
	// fmt.Println(msg)
	// msg = <-c
	// fmt.Println(msg)

	close(c)

	for msg := range c {
		fmt.Println(msg)
	}
}
