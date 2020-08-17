package gtds

import "C"
import "fmt"

//appUpdate is call everytime the app receives a notification
//it does not wait for any commands
//export appUpdate
func appUpdate() {
	select {
	case c := <-cmds:
		switch c {
		case cmdCreateWindow:
			platformCreateWindow(Window{"Title", 200, 200, 0})
		default:
			fmt.Println("NO")
		}
	default: //don't wait for commands
	}
}
