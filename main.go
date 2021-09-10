package main

import (
	"github.com/walkerxiong/ragel_example/core"
)

//go:generate ragel -Z core/remote_ctrl.rl

func main() {
	core.Command("turn on incr incr \n")
	core.PrintCurrentState()
	core.Command("incr\n")
	core.Command("mod")
	core.Command("turn off")
	core.Command("incr")
	core.Command("mod")
	core.PrintCurrentState()
}
