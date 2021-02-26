package web

import (
	"chinachess/battle"
	"chinachess/piece"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type web struct {
	*gin.Engine
	battle       *battle.Battle
	isBackground bool
}

func New(b *battle.Battle, isBackground bool) *web {
	w := &web{
		battle:       b,
		isBackground: isBackground,
	}
	engine := gin.Default()
	// 移动
	engine.POST("move", w.move)
	// 获取某个子可以移动的位置
	engine.GET("/canmove", w.getCanMove)
	// 获取获胜方
	engine.GET("/winside", w.winSide)
	// 获取整体布局
	engine.GET("/map", w.getMap)

	w.Engine = engine
	return w
}

const (
	port = ":9001"
)

var (
	ErrInvalidFrom = errors.New("invalid from point")
	ErrInvalidTo   = errors.New("invalid to point")
	ErrKill        = errors.New("can't kill ally")
)

func (w *web) Listen(addr string) {
	if w.isBackground {
		go log.Fatalln(w.Run(addr))
	} else {
		log.Fatalln(w.Run(addr))
	}
}

func (w *web) move(c *gin.Context) {
	var args moveArgs
	if err := c.ShouldBind(&args); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ppFrom := w.battle.GetPiece(args.From.X, args.From.Y)
	if ppFrom.IsBlank() {
		c.AbortWithError(http.StatusBadRequest, ErrInvalidFrom)
		return
	}
	ppTo := w.battle.GetPiece(args.To.X, args.To.Y)
	if !ppFrom.IsEnemy(ppTo) {
		c.AbortWithError(http.StatusBadRequest, ErrKill)
		return
	}

	if w.battle.Move(ppFrom, piece.Point{X: ppTo.X, Y: ppTo.Y}) {
		w.battle.Switch()
	}
}

type moveArgs struct {
	From piece.Point `json:"from"`
	To   piece.Point `json:"to"`
}

func (w *web) winSide(c *gin.Context) {
	c.JSON(http.StatusOK, w.battle.Winside)
}

func (w *web) getMap(c *gin.Context) {
	c.JSON(http.StatusOK, w.battle.Map)
}

func (w *web) getCanMove(c *gin.Context) {
	var args piece.Point
	if err := c.ShouldBind(&args); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	points := w.battle.GetPiece(args.X, args.Y).CanMove(&w.battle.Map)
	c.JSON(http.StatusOK, points)
}
