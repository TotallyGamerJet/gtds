package gtds

import (
	"github.com/faiface/mainthread"
)

type WindowStyle int

const (
	Borderless WindowStyle = 1 << iota
	Titled
	Closable
	Resizable
	Minimizable
)

func Run(handler Handler) {
	mainthread.Run(handler)
	platformRun()
}

type Window struct{}

type WindowConfig struct {
	Title         string
	Width, Height int
	Style         WindowStyle
}

type windowData struct {
	window WindowConfig
}

type Handler func()

func CreateWindow(w WindowConfig) Window {
	return platformCreateWindow(w)
}
