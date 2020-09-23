package gtds

import "C"
import "unsafe"

//export shouldClose
func shouldClose(ptr unsafe.Pointer) {
	w := Window{uintptr(ptr)}
	err := setData(w, func(data *windowData) {
		data.shouldClose = true
	})
	if err != nil {
		panic(err)
	}
}
