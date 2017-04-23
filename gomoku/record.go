package gomoku

//記録用
type Record struct {
	turns []int
	x     []int
	y     []int
}

func NewRecord() *Record {
	return &Record{
		turns: make([]int, 0),
		x:     make([]int, 0),
		y:     make([]int, 0),
	}
}

func (r *Record) Push(turn, x, y int) {
	r.turns = append(r.turns, turn)
	r.x = append(r.x, x)
	r.y = append(r.y, y)
}

func (r *Record) Last() (int, int) {
	return r.x[len(r.x)-1], r.y[len(r.y)-1]
}

func (r *Record) Winner() int {
	return r.turns[len(r.turns)-1]
}
