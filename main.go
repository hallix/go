package main

import (
	"sync"
	"time"
)

func main() {

	//Start time
	then := time.Now()

	bufferSize := 9000000000

	iterations := 3000

	consumerWorkers := 10

	producer := make(chan int)

	result := make(chan int, bufferSize)

	var wg sync.WaitGroup

	go produce(producer, iterations)

	wg.Add(consumerWorkers)

	for i := 0; i < consumerWorkers; i++ {
		go consumeAndCompute(producer, result, compute, &wg)
	}

	wg.Wait()
	close(result)

	wg.Add(1)
	go showResult(result, &wg)
	wg.Wait()

	//End time
	now := time.Now().Sub(then).Milliseconds()
	println("Time elapsed:", now, "ms")
}

func produce(p chan int, max int) {
	for i := 0; i < max; i++ {
		p <- i
	}
	close(p)
}

func compute(i int) int {
	n := 0
	limit := i * 20000
	for n < limit {
		n++
	}
	return n
}

func consumeAndCompute(p, r chan int, f func(num int) int, wg *sync.WaitGroup) {

	for i := range p {
		r <- f(i)
	}
	wg.Done()
}

func showResult(p chan int, wg *sync.WaitGroup) {
	var num int
	for i := range p {
		num = num + i
	}

	println("Result: ", num)
	wg.Done()
}
