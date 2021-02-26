package battle

import "chinachess/piece"

var Black = &Army{
	Side: piece.SideBlack,
}

func resetBlack() {
	Black.PieceDead = nil
	Black.Label = map[string]string{"color": "cyan"}
	Black.PieceNow = []*piece.PiecePoint{
		//
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceBing, Point: piece.Point{X: 0, Y: 6}, Name: "b卒1"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceBing, Point: piece.Point{X: 2, Y: 6}, Name: "b卒2"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceBing, Point: piece.Point{X: 4, Y: 6}, Name: "b卒3"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceBing, Point: piece.Point{X: 6, Y: 6}, Name: "b卒4"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceBing, Point: piece.Point{X: 8, Y: 6}, Name: "b卒5"},

		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PiecePao, Point: piece.Point{X: 1, Y: 7}, Name: "b炮1"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PiecePao, Point: piece.Point{X: 7, Y: 7}, Name: "b炮2"},

		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceJu, Point: piece.Point{X: 0, Y: 9}, Name: "b車1"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceJu, Point: piece.Point{X: 8, Y: 9}, Name: "b車2"},

		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceMa, Point: piece.Point{X: 1, Y: 9}, Name: "b马2"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceMa, Point: piece.Point{X: 7, Y: 9}, Name: "b马2"},

		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceXiang, Point: piece.Point{X: 2, Y: 9}, Name: "b像2"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceXiang, Point: piece.Point{X: 6, Y: 9}, Name: "b像2"},

		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceShi, Point: piece.Point{X: 3, Y: 9}, Name: "b仕2"},
		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceShi, Point: piece.Point{X: 5, Y: 9}, Name: "b仕2"},

		&piece.PiecePoint{Side: piece.SideBlack, Label: Black.Label, P: piece.PieceShuai, Point: piece.Point{X: 4, Y: 9}, Name: "b将"},
	}
}
