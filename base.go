package main

// 棋子
type Piece int

const (
	PieceBing Piece = 1 << iota
	PiecePao        //Piece = 1 << iota
	PieceJu
	PieceMa
	PieceXiang
	PieceShi
	PieceShuai
)
