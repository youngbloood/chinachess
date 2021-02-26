package piece

// 棋子
type Piece int

const (
	PieceNil         = 0         // 1
	PieceBing  Piece = 1 << iota // 1
	PiecePao                     // 2
	PieceJu                      // 4
	PieceMa                      // 8
	PieceXiang                   // 16
	PieceShi                     // 32
	PieceShuai                   // 64
)

func (p Piece) Type() Piece { return p }

func (p Piece) CanMove(m Map) []Point {
	return nil
}
