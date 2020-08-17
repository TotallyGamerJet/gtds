package main

import "fmt"

var cmds = make(chan int, 10)

func Run(handler Handler) {
	go handler()
	platformRun()
}

type Window struct {
	title         string
	width, height int
	style         int
}

type Handler func()

func CreateWindow(w Window) {
	select {
	case cmds <- 0:
	default:
		fmt.Println("I didn't expect this")
	}
}
