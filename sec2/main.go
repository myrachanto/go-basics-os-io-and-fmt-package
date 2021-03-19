package main

import (
	"fmt"
	"sync"
)
var (
	data = []string{"hello world!", "this is a very nice", "how cool"}
)
func main(){
	var wg sync.WaitGroup
	ch := make(chan string)
	producer(&wg, ch)
	consumer(&wg, ch)
	wg.Wait()
}
func consumer(wg *sync.WaitGroup, in <- chan string){
	wg.Add(1)
	go func(){
		for s := range in {
			fmt.Println(s)
		}
		wg.Done()
	}()
}
func producer(wg *sync.WaitGroup, out chan<- string){
	wg.Add(1)
	go func(){
		for _, s := range data {
			out<- s
		}
		close(out)
		wg.Done()
	}()
} 