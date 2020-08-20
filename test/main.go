package main

import (
	"letsgo"
)

func main() {
	gtds.Run(handler)
}

func handler() {
	gtds.CreateWindow(gtds.WindowConfig{Title: "from main", Width: 1080, Height: 720})
}
