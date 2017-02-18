package gomoku

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
	Board [][]int
}

// 碁盤の初期化を行います。
func NewGomoku(xsize int, ysize int) *Gomoku {
	board := make([][]int, ysize)
	for i := 0; i < ysize; i++ {
		board[i] = make([]int, xsize)
	}
	gomoku := &Gomoku{
		Board: board,
	}
	return gomoku
}

// 横方向(ー)で5目揃っているかを確認します。
func (gomoku *Gomoku) check_0_degrees() bool {
	return false
}

// 斜め方向(／)で5目揃っているかを確認します。
func (gomoku *Gomoku) check_45_degrees() bool {
	return false
}

// 縦方向(｜)で5目揃っているかを確認します。
func (gomoku *Gomoku) check_90_degrees() bool {
	return false
}

// 斜め方向(＼)で5目揃っているかを確認します。
func (gomoku *Gomoku) check_135_degrees() bool {
	return false
}

// 勝利判定
func (gomoku *Gomoku) Check() bool {
	return gomoku.check_0_degrees() || gomoku.check_45_degrees() || gomoku.check_90_degrees() || gomoku.check_135_degrees()
}

// 指定座標に打ち込みます。
func (gomoku *Gomoku) Put(status int, x int, y int) bool {
	return false
}
