package main

import "C"
import "fmt"

//appUpdate is call everytime the app receives a notification
//it does not wait for any commands
//export appUpdate
func appUpdate() {
	select {
	case c := <-cmds:
		if c == 0 {
			platformCreateWindow(Window{"title", 200, 200, 0})
		} else {
			fmt.Println("NO")
		}
	default: //don't wait for commands
	}
}
