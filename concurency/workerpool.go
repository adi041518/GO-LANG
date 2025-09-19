package main

import (
	"fmt"
	"time"
)

func work(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println(id, "worker is assigned to job", job)
		time.Sleep(time.Second)
		results <- job

	}
}
func main() {
	results := make(chan int, 5)
	jobs := make(chan int, 5)

	for i := 0; i < 3; i++ {
		go work(i, jobs, results)
	}

	for j := 0; j < 5; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Println("Result received:", result)
	}

}
