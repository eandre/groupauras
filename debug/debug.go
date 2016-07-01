package debug

import (
	"github.com/eandre/groupauras/pkg/draw"
	"github.com/eandre/lunar-wow/pkg/wow"
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
		Texture:    "diamond",
		SizePixels: 5,
	})
}

func init() {
	wow.RegisterEvent("PLAYER_ENTERING_WORLD", onWorldEnter)
}
