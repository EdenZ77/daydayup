package main

import "fmt"

/*
参考资料：https://time.geekbang.org/column/article/208572?screen=full
*/

// Color 类型
type Color string

// 定义棋子颜色的常量
const (
	Red   Color = "red"
	Black Color = "black"
)

// ChessPieceUnit 享元类
type ChessPieceUnit struct {
	id    int
	text  string
	color Color
}

// NewChessPieceUnit 创建一个新的 ChessPieceUnit
func NewChessPieceUnit(id int, text string, color Color) *ChessPieceUnit {
	return &ChessPieceUnit{
		id:    id,
		text:  text,
		color: color,
	}
}

// ChessPieceUnitFactory 享元工厂
type ChessPieceUnitFactory struct {
	pieces map[int]*ChessPieceUnit
}

// NewChessPieceUnitFactory 创建一个新的 ChessPieceUnitFactory
func NewChessPieceUnitFactory() *ChessPieceUnitFactory {
	pieces := make(map[int]*ChessPieceUnit)
	pieces[1] = NewChessPieceUnit(1, "车", Black)
	pieces[2] = NewChessPieceUnit(2, "马", Black)
	// ...省略摆放其他棋子的代码...
	return &ChessPieceUnitFactory{pieces: pieces}
}

// GetChessPiece 返回指定 ID 的 ChessPieceUnit
func (f *ChessPieceUnitFactory) GetChessPiece(id int) *ChessPieceUnit {
	return f.pieces[id]
}

// ChessPiece 棋子类
type ChessPiece struct {
	chessPieceUnit *ChessPieceUnit
	positionX      int
	positionY      int
}

// ChessBoard 棋盘类
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

// NewChessBoard 创建一个新的 ChessBoard
func NewChessBoard(factory *ChessPieceUnitFactory) *ChessBoard {
	board := &ChessBoard{chessPieces: make(map[int]*ChessPiece)}
	board.chessPieces[1] = &ChessPiece{chessPieceUnit: factory.GetChessPiece(1), positionX: 0, positionY: 0}
	board.chessPieces[2] = &ChessPiece{chessPieceUnit: factory.GetChessPiece(2), positionX: 1, positionY: 0}
	// ...省略摆放其他棋子的代码...
	return board
}

// Move 移动棋子
func (b *ChessBoard) Move(id, x, y int) {
	if chessPiece, ok := b.chessPieces[id]; ok {
		chessPiece.positionX = x
		chessPiece.positionY = y
	}
}

func main() {
	factory := NewChessPieceUnitFactory()
	board := NewChessBoard(factory)

	fmt.Printf("Initial Position of 1: %d,%d\n", board.chessPieces[1].positionX, board.chessPieces[1].positionY)
	board.Move(1, 2, 3)
	fmt.Printf("Moved Position of 1: %d,%d\n", board.chessPieces[1].positionX, board.chessPieces[1].positionY)
}
