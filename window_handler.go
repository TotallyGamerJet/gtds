package gtds

import (
	"fmt"
)

func init() {
	go windowHandler()
}

var windowChan = make(chan func(map[Window]windowData))

var (
	ErrWindowNoExist = fmt.Errorf("window doesn't exist")
)

func windowHandler() {
	windowList := make(map[Window]windowData)
	for op := range windowChan {
		op(windowList)
	}
}

func setData(w Window, f func(*windowData)) error {
	err := make(chan error)
	windowChan <- func(m map[Window]windowData) {
		_, ok := m[w]
		if !ok {
			err <- ErrWindowNoExist
			return
		}
		data := m[w]
		f(&data)
		m[w] = data
		err <- nil
	}
	return <-err
}

func getData(win Window) windowData {
	ret := make(chan windowData)
	windowChan <- func(m map[Window]windowData) {
		d, ok := m[win]
		if !ok { //If it doesn't exist make a new one
			m[win] = windowData{window: win}
			ret <- m[win]
			return
		}
		ret <- d
	}
	return <-ret
}

func removeData(win Window) error {
	err := make(chan error)
	windowChan <- func(m map[Window]windowData) {
		_, ok := m[win]
		if !ok {
			err <- ErrWindowNoExist
			return
		}
		delete(m, win)
		err <- nil
	}
	return <-err
}
