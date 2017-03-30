package main

import (
	"./gomoku"
	_"fmt"
)

func main() {
	g := gomoku.NewGomoku()
	g.Put(1,6)
	g.Display()
}
