package battle

import "chinachess/piece"

var Red = &Army{
	Side: piece.SideRed,
}

func resetRed() {
	Red.PieceDead = nil
	Red.Label = map[string]string{"color": "red"}
	Red.PieceNow = []*piece.PiecePoint{
		//
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceBing, Point: piece.Point{X: 0, Y: 3}, Name: "r兵1"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceBing, Point: piece.Point{X: 2, Y: 3}, Name: "r兵2"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceBing, Point: piece.Point{X: 4, Y: 3}, Name: "r兵3"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceBing, Point: piece.Point{X: 6, Y: 3}, Name: "r兵4"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceBing, Point: piece.Point{X: 8, Y: 3}, Name: "r兵5"},

		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PiecePao, Point: piece.Point{X: 1, Y: 2}, Name: "r炮1"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PiecePao, Point: piece.Point{X: 7, Y: 2}, Name: "r炮2"},

		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceJu, Point: piece.Point{X: 0, Y: 0}, Name: "r车1"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceJu, Point: piece.Point{X: 8, Y: 0}, Name: "r车2"},

		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceMa, Point: piece.Point{X: 1, Y: 0}, Name: "r马1"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceMa, Point: piece.Point{X: 7, Y: 0}, Name: "r马2"},

		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceXiang, Point: piece.Point{X: 2, Y: 0}, Name: "r象1"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceXiang, Point: piece.Point{X: 6, Y: 0}, Name: "r象2"},

		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceShi, Point: piece.Point{X: 3, Y: 0}, Name: "r士1"},
		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceShi, Point: piece.Point{X: 5, Y: 0}, Name: "r士2"},

		&piece.PiecePoint{Side: piece.SideRed, Label: Red.Label, P: piece.PieceShuai, Point: piece.Point{X: 4, Y: 0}, Name: "r帅"},
	}
}
