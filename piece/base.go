package piece

type Mover interface {
	Type() Piece          // 类型
	CanMove(*Map) []Point // 可以走到的位置
}

type Side int

/*
红：下
黑：上
*/
const (
	SideNil Side = iota
	SideRed
	SideBlack
)
