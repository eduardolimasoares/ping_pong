package main

import (
	"fmt"
	"time"
)

func Ping(ping <-chan string, pong chan<- string) {
	for m := range ping {
		printAndDelay(m)
		pong <- "pong"
	}
}

func Pong(ping chan<- string, pong <-chan string) {
	for m := range pong {
		printAndDelay(m)
		ping <- "ping"
	}
}

func printAndDelay(msg string) {
	fmt.Println(msg)
	time.Sleep(time.Second)
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go Ping(ping, pong)

	go Pong(ping, pong)

	ping <- "ping"

	for {
	}
}
