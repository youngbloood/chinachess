package piece

// 基础点
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// 棋子所在点
type PiecePoint struct {
	Point
	P     Piece
	Side  Side
	Name  string
	Label map[string]string
}

var PiecePointNil = &PiecePoint{}

func (pp PiecePoint) Type() Piece { return pp.P }

func (pp PiecePoint) CanMove(m *Map) []Point {
	switch pp.P {
	case PieceNil:
		return nil
	case PieceBing:
		return pp.MoveBing(m)
	case PiecePao:
		return pp.MovePao(m)
	case PieceJu:
		return pp.MoveJu(m)
	case PieceMa:
		return pp.MoveMa(m)
	case PieceXiang:
		return pp.MoveXiang(m)
	case PieceShi:
		return pp.MoveShi(m)
	case PieceShuai:
		return pp.MoveShuai(m)
	}
	return nil
}

func (pp *PiecePoint) IsBlank() bool {
	return pp == PiecePointNil
}

func (pp *PiecePoint) MoveBing(m *Map) (result []Point) {
	if pp.P != PieceBing {
		return nil
	}
	switch pp.Side {
	case SideRed:
		if pp.IsInMapAfterMove(0, 1) && m.GetPiece(pp.X, pp.Y+1).IsEnemy(pp) {
			result = append(result, Point{X: pp.X, Y: pp.Y + 1})
		}
		if pp.IsInMapAfterMove(1, 0) && m.GetPiece(pp.X+1, pp.Y).IsEnemy(pp) && !pp.IsInRed() {
			result = append(result, Point{X: pp.X + 1, Y: pp.Y})
		}
		if pp.IsInMapAfterMove(-1, 0) && m.GetPiece(pp.X-1, pp.Y).IsEnemy(pp) && !pp.IsInRed() {
			result = append(result, Point{X: pp.X - 1, Y: pp.Y})
		}
	case SideBlack:
		if pp.IsInMapAfterMove(0, -1) && m.GetPiece(pp.X, pp.Y-1).IsEnemy(pp) {
			result = append(result, Point{X: pp.X, Y: pp.Y - 1})
		}
		if pp.IsInMapAfterMove(1, 0) && m.GetPiece(pp.X+1, pp.Y).IsEnemy(pp) && pp.IsInRed() {
			result = append(result, Point{X: pp.X + 1, Y: pp.Y})
		}
		if pp.IsInMapAfterMove(-1, 0) && m.GetPiece(pp.X-1, pp.Y).IsEnemy(pp) && pp.IsInRed() {
			result = append(result, Point{X: pp.X - 1, Y: pp.Y})
		}
	}
	return
}

func (pp *PiecePoint) MovePao(m *Map) (result []Point) {
	if pp.Type() != PiecePao {
		return nil
	}

	pivot := false
	for x := pp.X + 1; x < Row; x++ {
		piece := m.GetPiece(x, pp.Y)
		if piece.IsBlank() && !pivot {
			result = append(result, Point{X: x, Y: pp.Y})
		}
		if !piece.IsBlank() && !pivot {
			pivot = true
			continue
		}
		if !piece.IsBlank() && pivot && piece.IsEnemy(pp) {
			result = append(result, Point{X: x, Y: pp.Y})
			break
		}
	}

	pivot = false
	for x := pp.X - 1; x >= 0; x-- {
		piece := m.GetPiece(x, pp.Y)
		if piece.IsBlank() && !pivot {
			result = append(result, Point{X: x, Y: pp.Y})
		}
		if !piece.IsBlank() && !pivot {
			pivot = true
			continue
		}
		if !piece.IsBlank() && pivot && piece.IsEnemy(pp) {
			result = append(result, Point{X: x, Y: pp.Y})
			break
		}
	}

	pivot = false
	for y := pp.Y + 1; y < Col; y++ {
		piece := m.GetPiece(pp.X, y)
		if piece.IsBlank() && !pivot {
			result = append(result, Point{X: pp.X, Y: y})
		}
		if !piece.IsBlank() && !pivot {
			pivot = true
			continue
		}
		if !piece.IsBlank() && pivot && piece.IsEnemy(pp) {
			result = append(result, Point{X: pp.X, Y: y})
			break
		}
	}

	pivot = false
	for y := pp.Y - 1; y >= 0; y-- {
		piece := m.GetPiece(pp.X, y)
		if piece.IsBlank() && !pivot {
			result = append(result, Point{X: pp.X, Y: y})
		}
		if !piece.IsBlank() && !pivot {
			pivot = true
			continue
		}
		if !piece.IsBlank() && pivot && piece.IsEnemy(pp) {
			result = append(result, Point{X: pp.X, Y: y})
			break
		}
	}
	return
}

func (pp *PiecePoint) MoveJu(m *Map) (result []Point) {
	if pp.P != PieceJu {
		return nil
	}

	for x := pp.X + 1; x < Row; x++ {
		piece := m.GetPiece(x, pp.Y)
		if piece.IsBlank() {
			result = append(result, Point{X: x, Y: pp.Y})
			continue
		}
		if piece.IsEnemy(pp) {
			result = append(result, Point{X: x, Y: pp.Y})
		}
		break
	}

	for x := pp.X - 1; x >= 0; x-- {
		piece := m.GetPiece(x, pp.Y)
		if piece.IsBlank() {
			result = append(result, Point{X: x, Y: pp.Y})
			continue
		}
		if piece.IsEnemy(pp) {
			result = append(result, Point{X: x, Y: pp.Y})
		}
		break
	}

	for y := pp.Y + 1; y < Col; y++ {
		piece := m.GetPiece(pp.X, y)
		if piece.IsBlank() {
			result = append(result, Point{X: pp.X, Y: y})
			continue
		}
		if piece.IsEnemy(pp) {
			result = append(result, Point{X: pp.X, Y: y})
		}
		break
	}

	for y := pp.Y - 1; y >= 0; y-- {
		piece := m.GetPiece(pp.X, y)
		if piece.IsBlank() {
			result = append(result, Point{X: pp.X, Y: y})
			continue
		}
		if piece.IsEnemy(pp) {
			result = append(result, Point{X: pp.X, Y: y})
		}
		break
	}
	return
}

func (pp *PiecePoint) MoveMa(m *Map) (result []Point) {
	if pp.P != PieceMa {
		return nil
	}

	if pp.IsInMapAfterMove(0, 1) && m.GetPiece(pp.X, pp.Y+1).IsEnemy(pp) {
		if pp.IsInMapAfterMove(1, 2) {
			result = append(result, Point{X: pp.X + 1, Y: pp.Y + 2})
		}
		if pp.IsInMapAfterMove(-1, 2) {
			result = append(result, Point{X: pp.X - 1, Y: pp.Y + 2})
		}
	}

	if pp.IsInMapAfterMove(0, -1) && m.GetPiece(pp.X, pp.Y-1).IsEnemy(pp) {
		if pp.IsInMapAfterMove(1, -2) {
			result = append(result, Point{X: pp.X + 1, Y: pp.Y - 2})
		}
		if pp.IsInMapAfterMove(-1, -2) {
			result = append(result, Point{X: pp.X - 1, Y: pp.Y - 2})
		}
	}

	if pp.IsInMapAfterMove(1, 0) && m.GetPiece(pp.X+1, pp.Y).IsEnemy(pp) {
		if pp.IsInMapAfterMove(2, 1) {
			result = append(result, Point{X: pp.X + 2, Y: pp.Y + 1})
		}
		if pp.IsInMapAfterMove(2, -1) {
			result = append(result, Point{X: pp.X + 2, Y: pp.Y - 1})
		}
	}

	if pp.IsInMapAfterMove(-1, 0) && m.GetPiece(pp.X-1, pp.Y).IsEnemy(pp) {
		if pp.IsInMapAfterMove(-2, 1) {
			result = append(result, Point{X: pp.X - 2, Y: pp.Y + 1})
		}
		if pp.IsInMapAfterMove(-2, -1) {
			result = append(result, Point{X: pp.X - 2, Y: pp.Y - 1})
		}
	}
	return
}

func (pp *PiecePoint) MoveXiang(m *Map) (result []Point) {
	if pp.Type() != PieceXiang {
		return nil
	}

	switch pp.Side {
	case SideBlack:
		if pp.IsInMapAfterMove(1, 1) && m.GetPiece(pp.X+1, pp.Y+1).IsEnemy(pp) && pp.IsInMapAfterMove(2, 2) && !pp.IsInRed() {
			result = append(result, Point{X: pp.X + 2, Y: pp.Y + 2})
		}
		if pp.IsInMapAfterMove(1, -1) && m.GetPiece(pp.X+1, pp.Y-1).IsEnemy(pp) && pp.IsInMapAfterMove(2, -2) && !pp.IsInRed() {
			result = append(result, Point{X: pp.X + 2, Y: pp.Y - 2})
		}
		if pp.IsInMapAfterMove(-1, 1) && m.GetPiece(pp.X-1, pp.Y+1).IsEnemy(pp) && pp.IsInMapAfterMove(-2, 2) && !pp.IsInRed() {
			result = append(result, Point{X: pp.X - 2, Y: pp.Y + 2})
		}
		if pp.IsInMapAfterMove(-1, -1) && m.GetPiece(pp.X-1, pp.Y-1).IsEnemy(pp) && pp.IsInMapAfterMove(-2, -2) && !pp.IsInRed() {
			result = append(result, Point{X: pp.X - 2, Y: pp.Y - 2})
		}

	case SideRed:
		if pp.IsInMapAfterMove(1, 1) && m.GetPiece(pp.X+1, pp.Y+1).IsEnemy(pp) && pp.IsInMapAfterMove(2, 2) && pp.IsInRed() {
			result = append(result, Point{X: pp.X + 2, Y: pp.Y + 2})
		}
		if pp.IsInMapAfterMove(1, -1) && m.GetPiece(pp.X+1, pp.Y-1).IsEnemy(pp) && pp.IsInMapAfterMove(2, -2) && pp.IsInRed() {
			result = append(result, Point{X: pp.X + 2, Y: pp.Y - 2})
		}
		if pp.IsInMapAfterMove(-1, 1) && m.GetPiece(pp.X-1, pp.Y+1).IsEnemy(pp) && pp.IsInMapAfterMove(-2, 2) && pp.IsInRed() {
			result = append(result, Point{X: pp.X - 2, Y: pp.Y + 2})
		}
		if pp.IsInMapAfterMove(-1, -1) && m.GetPiece(pp.X-1, pp.Y-1).IsEnemy(pp) && pp.IsInMapAfterMove(-2, -2) && pp.IsInRed() {
			result = append(result, Point{X: pp.X - 2, Y: pp.Y - 2})
		}
	}

	return
}

func (pp *PiecePoint) MoveShi(m *Map) (result []Point) {
	if pp.Type() != PieceShi {
		return nil
	}

	if pp.IsInBasecamp(-1, -1) && m.GetPiece(pp.X-1, pp.Y-1).IsEnemy(pp) {
		result = append(result, Point{X: pp.X - 1, Y: pp.Y - 1})
	}
	if pp.IsInBasecamp(1, 1) && m.GetPiece(pp.X+1, pp.Y+1).IsEnemy(pp) {
		result = append(result, Point{X: pp.X + 1, Y: pp.Y + 1})
	}
	if pp.IsInBasecamp(1, -1) && m.GetPiece(pp.X+1, pp.Y-1).IsEnemy(pp) {
		result = append(result, Point{X: pp.X + 1, Y: pp.Y - 1})
	}
	if pp.IsInBasecamp(-1, 1) && m.GetPiece(pp.X-1, pp.Y+1).IsEnemy(pp) {
		result = append(result, Point{X: pp.X - 1, Y: pp.Y + 1})
	}

	return
}

func (pp *PiecePoint) MoveShuai(m *Map) (result []Point) {
	if pp.Type() != PieceShuai {
		return nil
	}
	if pp.IsInBasecamp(0, -1) && m.GetPiece(pp.X, pp.Y-1).IsEnemy(pp) {
		result = append(result, Point{X: pp.X, Y: pp.Y - 1})
	}
	if pp.IsInBasecamp(0, 1) && m.GetPiece(pp.X, pp.Y+1).IsEnemy(pp) {
		result = append(result, Point{X: pp.X, Y: pp.Y + 1})
	}
	if pp.IsInBasecamp(1, 0) && m.GetPiece(pp.X+1, pp.Y).IsEnemy(pp) {
		result = append(result, Point{X: pp.X + 1, Y: pp.Y})
	}
	if pp.IsInBasecamp(-1, 0) && m.GetPiece(pp.X-1, pp.Y).IsEnemy(pp) {
		result = append(result, Point{X: pp.X - 1, Y: pp.Y})
	}
	return
}

// 移动(x,y)的向量之后是否还在Map上
func (pp PiecePoint) IsInMapAfterMove(x, y int) bool {
	x = pp.X + x
	y = pp.Y + y
	return x >= 0 && x < Row && y >= 0 && y < Col
}

// 该棋子是否在红区域
func (pp PiecePoint) IsInRed() bool {
	return pp.Y <= 4
}

// 该棋子移动(x,y)向量之后是否在帅营
func (pp PiecePoint) IsInBasecamp(x, y int) bool {
	x = pp.X + x
	y = pp.Y + y
	switch pp.Side {
	case SideBlack:
		return y < Col && y >= 7 && x >= 3 && x <= 5
	case SideRed:
		return y >= 0 && y <= 2 && x >= 3 && x <= 5
	}
	return false
}

// pp与p是否为敌军
func (pp *PiecePoint) IsEnemy(p *PiecePoint) bool {
	return pp.Side != p.Side
}

// 移动到(x,y)
func (pp *PiecePoint) MoveTo(x, y int) {
	pp.X = x
	pp.Y = y
}

// func (p PiecePoint) Name() string {
// 	var suffix string
// 	switch p.Side {
// 	case SideRed:
// 		suffix = "r"
// 	case SideBlack:
// 		suffix = "b"
// 	}
// 	switch p.P {
// 	case PieceBing:
// 		return suffix + "兵"
// 	case PiecePao:
// 		return suffix + "炮"
// 	case PieceJu:
// 		return suffix + "車"
// 	case PieceMa:
// 		return suffix + "马"
// 	case PieceXiang:
// 		return suffix + "象"
// 	case PieceShi:
// 		return suffix + "仕"
// 	case PieceShuai:
// 		return suffix + "帅"
// 	}
// 	return ""
// }
