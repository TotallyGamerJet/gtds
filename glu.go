package gtds

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

type WindowStyle int

const (
	Borderless WindowStyle = 1 << iota
	Titled
	Closable
	Resizable
	Hideable
	Fullscreen
)

type Window struct {
	ptr uintptr
}

type WindowConfig struct {
	Title         string
	Width, Height int
	Style         WindowStyle
}

type windowData struct {
	window      Window
	shouldClose bool
}

func CreateWindow(w WindowConfig) Window {
	if w.Width < 0 || w.Height < 0 {
		panic("Improper Window Dimensions")
	}
	return platformCreateWindow(w)
}

func PollEvents() {
	platformPollEvents()
}

func (w Window) ShouldClose() bool {
	return getData(w).shouldClose
}

func Run(run func() error) error {
	platformInit()
	return run()
}
