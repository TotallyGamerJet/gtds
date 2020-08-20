package gtds

import "C" //required for
import "fmt"

//appUpdate is called everytime the app receives a notification
//it will attempt to select any incoming command but does not wait
//for any to arrive
//export appUpdate
func appUpdate() {
	select {
	case c := <-cmds:
		switch c.code {
		case cmdCreateWindow:
			platformCreateWindow(c.data.(WindowConfig))
		default:
			fmt.Println("NO")
		}
	default: //don't wait for commands
	}
}
