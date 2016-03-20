package draw

import (
	"github.com/eandre/groupauras/shim/bridge"
	"github.com/eandre/groupauras/shim/luamath"
)

var activePoints map[*Point]bool

var updateInterval float32 = 1 / 60.0
var updateTotal float32 = 0

func markPointActive(p *Point) {
	activePoints[p] = true
}

func markPointInactive(p *Point) {
	delete(activePoints, p)
}

func onUpdate(dt float32) {
	updateTotal += dt
	if updateTotal > updateInterval {
		steps := luamath.Floor(updateTotal / updateInterval)
		elapsed := updateInterval * steps
		updateTotal -= elapsed

		for p := range activePoints {
			p.update()
		}
	}
}

func init() {
	bridge.RegisterUpdate(onUpdate)
}
