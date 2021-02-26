package battle

import (
	"chinachess/piece"
)

var Map piece.Map

// 重置局
func resetMap() {
	for x := 0; x < 9; x++ {
		for y := 0; y < 10; y++ {
			Map[x][y] = piece.PiecePointNil
		}
	}

	for _, point := range Black.PieceNow {
		// fmt.Printf("黑(%d,%d,%s) ", point.X, point.Y, point.Name())
		Map[point.X][point.Y] = point
	}

	for _, point := range Red.PieceNow {
		// fmt.Printf("红(%d,%d,%s) ", point.X, point.Y, point.Name())
		Map[point.X][point.Y] = point
	}
	// fmt.Println("map = ", Map)
}
