package main

import "chinachess/piece"

func promptLayout(m piece.Map, label map[string]interface{}) string {
	var val string
	row := len(m)
	col := len(m[0])
	for y := col; y >= 0; y-- {
		for x := 0; x < row; x++ {
			name := m.GetPiece(x, y).Name
			if name == "" {
				name = "+"
				val += name + "----"
			} else {
				val += name + "--"
			}
		}
		val += "\n"
		for y := 0; y < col; y++ {
			val += "|    "
		}
		val += "\n"
	}
	return val
}
