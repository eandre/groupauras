package debug

import (
	"github.com/eandre/groupauras/auras/emerald_nightmare/nythendra"
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/groupauras/pkg/draw"
)

var Point *draw.Point

func onWorldEnter(event string, args []interface{}) {
	Point = draw.NewPoint(&draw.PointCfg{
		Pos:           draw.StaticPlayerPosition(),
		Texture:       "timer",
		SizeYards:     40,
		RotateDegrees: 360,
		RotateSpeed:   5,
	})

	draw.NewPoint(&draw.PointCfg{
		Pos:        draw.PlayerPosition(),
		Texture:    "star",
		SizePixels: 10,
	})

	nythendra.Enable(nil)
}

func init() {
	//wow.RegisterEvent("PLAYER_ENTERING_WORLD", onWorldEnter)
	nythendra.Enable(context.Base)
}
