package gomoku

import (
	"fmt"
)

// 盤面の目の状態を表します。
// NONE  - 何も打たれていない。
// WHITE - 白石が打たれている。
// BLACK - 黒石が打たれている。

const (
	NONE  int = 0
	WHITE int = 1
	BLACK int = 2
)

// 碁盤です。
type Gomoku struct {
	turn   int
	board  [][]int
	record *Record
}

// 碁盤の初期化を行います。
func NewGomoku() *Gomoku {
	board := make([][]int, 19)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, 19)
	}
	return &Gomoku{
		turn:   BLACK,
		board:  board,
		record: NewRecord(),
	}
}

// 勝利判定
func (g *Gomoku) Check() bool {
	return g.check_row() || g.check_column() || g.check_slant()
}

// 指定座標に打ち込みます。
func (g *Gomoku) Put(x, y int) bool {
	if g.board[x][y] != NONE {
		return false
	}

	if g.board[x][y] = g.turn; g.turn == BLACK {
		g.turn = WHITE
	} else {
		g.turn = BLACK
	}
	return true
}

//碁盤の出力
func (g *Gomoku) Display() {
	for i := 0; i < len(g.board); i++ {
		fmt.Println(g.board[i])
	}
}

// 横方向(ー)で5目揃っているかを確認します。
func (g *Gomoku) check_row() bool {
	return false
}

// 縦方向(｜)で5目揃っているかを確認します。
func (g *Gomoku) check_column() bool {
	return false
}

// 斜め方向(＼)で5目揃っているかを確認します。
func (g *Gomoku) check_slant() bool {
	return false
}
