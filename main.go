package main
import "./gomoku"
import "fmt"

func main() {
  g := gomoku.NewGomoku(15, 15)
  g.Put(gomoku.BLACK, 5, 5)
  fmt.Println(g.Board)
  fmt.Println(g.Check())
}


