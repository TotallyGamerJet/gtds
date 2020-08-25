package main

import (
	"letsgo"
)

func main() {
	gtds.Run(handler)
}

func handler() {
	gtds.CreateWindow(gtds.WindowConfig{Title: "from main", Style: gtds.Titled | gtds.Closable, Width: 720, Height: 480})
}
