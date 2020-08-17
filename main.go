package main

import "time"

func main() {
	Run(handler)
}

func handler() {
	CreateWindow(Window{"title", 200, 200, 0})
	time.Sleep(2 * time.Second)
	CreateWindow(Window{"window 2", 200, 200, 0})
}
