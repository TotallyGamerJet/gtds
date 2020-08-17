package gtds

var cmds = make(chan command, 10)

type command int

const (
	cmdCreateWindow command = iota
)

func Run(handler Handler) {
	go handler()
	platformRun()
}

type Window struct {
	Title         string
	Width, Height int
	Style         int
}

type windowData struct {
	window Window
}

type Handler func()

func CreateWindow(w Window) {
	select {
	case cmds <- cmdCreateWindow:
	}
}
