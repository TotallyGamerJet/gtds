package main

import (
	gtds "letsgo"
)

func main() {
	gtds.Run(handler)
}

func handler() {
	gtds.CreateWindow(gtds.Window{Title: "title", Width: 200, Height: 200})
}
