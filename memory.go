package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
	
}

func sendData(ch chan string) {
	ch <- "Wasington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Shanghai"
}

func getData(ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Printf("%s ", input)
	}
}
