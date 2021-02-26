package piece

import "fmt"

const (
	Row = 9
	Col = 10
)

type Map [Row][Col]*PiecePoint

func (m Map) GetPiece(x, y int) *PiecePoint {
	if x < 0 || x >= Row || y < 0 || y >= Col {
		return PiecePointNil
	}
	return m[x][y]
}

func (m *Map) SetPiece(x, y int, pp *PiecePoint) {
	if x < 0 || x >= Row || y < 0 || y >= Col {
		return
	}
	(*m)[x][y] = pp
}

func (m Map) IsEnemy(side Side, x, y int) bool {
	pp := m.GetPiece(x, y)
	if pp.IsBlank() {
		return true
	}
	return side != pp.Side
}

func (m Map) Print() {
	val := m.GetLayout()
	fmt.Println(val)
}

func (m Map) GetLayout() string {
	var val string
	for y := Col; y >= 0; y-- {
		for x := 0; x < Row; x++ {
			name := m.GetPiece(x, y).Name
			if name == "" {
				name = "+"
				val += name + "----"
			} else {
				val += name + "--"
			}
		}
		val += "\n"
		for y := 0; y < Col; y++ {
			val += "|    "
		}
		val += "\n"
	}
	return val
}

func (m Map) Print2() {
	fmt.Println(m)
}
