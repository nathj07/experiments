package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// task1()
	// task2()
	task2a()
	// task3()
	//task4()
}

// Square takes a float and returns the square
func Square(x float64) float64 {
	time.Sleep(1 * time.Second)
	return x * x
}

func task1() {
	var wg sync.WaitGroup
	var i float64
	for i = 0; i <= 9; i++ {
		wg.Add(1)
		go func(j float64) {
			defer wg.Done()
			fmt.Println(Square(j))
		}(i)
	}
	wg.Wait()
}

func task2() {
	input := make(chan float64)
	var i float64
	var wg sync.WaitGroup
	for j := 0; j <= 9; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(Square(<-input))
		}()
	}
	for i = 0; i < 10; i++ {
		input <- i
	}
	wg.Wait()
}

func task2a() {
	input := make(chan float64)
	var i float64
	var wg sync.WaitGroup
	for j := 0; j <= 9; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for inp := range input {
				fmt.Println(Square(inp))
			}
		}()
	}
	for i = 0; i < 50; i++ {
		input <- i
	}
	close(input)
	wg.Wait()

}

func task3() {
	fmt.Println(SumSquares(float64(10)))
}

// SumSquares sums the squares of the numbers from 0 to n-1
func SumSquares(n float64) float64 {
	input := make(chan float64)
	var i, j, sum float64
	var wg sync.WaitGroup
	for j = 0; j <= n; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum = sum + Square(<-input)
		}()
	}
	for i = 0; i <= n; i++ {
		input <- i
	}
	wg.Wait()
	return sum
}

func task4() {
	var wg sync.WaitGroup
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println(SumSquares(float64(10 * j)))
		}(i)
	}
	wg.Wait()
}
