package main

import (
	"./core"
)

//go:generate ragel -Z core/remote_ctrl.rl

func main() {
	core.Command("turn on incr incr \n")
	core.PrintCurrentState()
	core.Command("incr\n")
	core.Command("turn off")
	core.PrintCurrentState()
}
