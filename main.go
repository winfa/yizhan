package main

import (
	"syscall/js"
	"yizhan/calculate"
)

func main() {
	window := js.Global()
	window.Set("calculate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return calculate.ExcuteJinjueZhan()
	}))
	select {}
}
