package battle

import "chinachess/piece"

type Army struct {
	Side      piece.Side
	PieceNow  []*piece.PiecePoint
	PieceDead []piece.Piece
	TimeUse   int64 // 用时
	Label     map[string]string
}

type Battler interface {
	Side() piece.Side
}

// 重置棋盘
func ResetBoard() {
	resetBlack()
	resetRed()
	resetMap()
}

// 是否将军
func (a *Army) IsCheckmate(m *piece.Map) bool {
	for _, point := range a.PieceNow {
		cans := point.CanMove(m)
		for _, can := range cans {
			if m.GetPiece(can.X, can.Y).Type() == piece.PieceShuai {
				return true
			}
		}
	}
	return false
}
