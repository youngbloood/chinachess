package battle

import "chinachess/piece"

type Battle struct {
	piece.Map
	Black, Red  *Army
	isRed       bool
	isCheckmate bool
	Winside     piece.Side
	getLayout   func(piece.Map, map[string]interface{}) string
}

var GBattle *Battle

func InitBattle(isRed bool, getLayout func(piece.Map, map[string]interface{}) string) {
	GBattle = &Battle{
		Black:     Black,
		Red:       Red,
		Map:       Map,
		isRed:     isRed,
		getLayout: getLayout,
	}
}

/*

 */
func (b Battle) IsOver() bool {
	return b.Winside == piece.SideBlack || b.Winside == piece.SideRed
}

func (b *Battle) SideNow() piece.Side {
	if b.isRed {
		return piece.SideRed
	}
	return piece.SideBlack
}

func (b *Battle) Switch() {
	if b.IsOver() {
		return
	}

	if b.isRed {
		b.isCheckmate = b.Red.IsCheckmate(&b.Map)
		if b.isCheckmate {
			isOver := false
			for _, pn := range b.Black.PieceNow {
				points := pn.CanMove(&b.Map)
				originX := pn.X
				originY := pn.Y
				for _, point := range points {
					pn.MoveTo(point.X, point.Y)
					if b.Red.IsCheckmate(&b.Map) {
						pn.MoveTo(originX, originY)
						isOver = true
					} else {
						isOver = false
					}
				}
			}
			if isOver {
				// 结束
				b.Winside = piece.SideRed
			}
		}
	} else {
		b.isCheckmate = b.Black.IsCheckmate(&b.Map)
		if b.isCheckmate {
			isOver := false
			for _, pn := range b.Red.PieceNow {
				points := pn.CanMove(&b.Map)
				originX := pn.X
				originY := pn.Y
				for _, point := range points {
					pn.MoveTo(point.X, point.Y)
					if b.Black.IsCheckmate(&b.Map) {
						pn.MoveTo(originX, originY)
						isOver = true
					} else {
						isOver = false
					}
				}
			}
			if isOver {
				// 结束
				b.Winside = piece.SideBlack
			}
		}
	}
	b.isRed = !b.isRed
}

// 将p移动到pos
func (b *Battle) Move(pp *piece.PiecePoint, pos piece.Point) bool {
	if b.IsOver() ||
		pos.X < 0 || pos.X >= piece.Row || pos.Y < 0 || pos.Y >= piece.Col ||
		(b.isRed && pp.Side != piece.SideRed) ||
		(!b.isRed && pp.Side != piece.SideBlack) {
		return false
	}

	(&b.Map).SetPiece(pp.X, pp.Y, piece.PiecePointNil)
	pp.X = pos.X
	pp.Y = pos.Y
	(&b.Map).SetPiece(pp.X, pp.Y, pp)
	return true
}

func (b *Battle) GetLayout(label map[string]interface{}) string {
	if b.getLayout != nil {
		return b.getLayout(b.Map, label)
	}
	return b.Map.GetLayout()
}
