package gtds

var cmds = make(chan command, 10)

type command struct {
	code commandCode
	data interface{}
}

type commandCode int

const (
	cmdCreateWindow commandCode = iota
)

func Run(handler Handler) {
	go handler()
	platformRun()
}

type Window struct{}

type WindowConfig struct {
	Title         string
	Width, Height int
	Style         int
}

type windowData struct {
	window WindowConfig
}

type Handler func()

func CreateWindow(w WindowConfig) Window {
	select {
	case cmds <- command{code: cmdCreateWindow, data: w}:
	}
	return Window{}
}
