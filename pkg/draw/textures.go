package draw

import "github.com/eandre/lunar-wow/pkg/widget"

type textureDef struct {
	Texture    string
	Blend      widget.BlendMode
	TexCoords  []float32
	SizeScalar float32
}

var textureMap = map[string]*textureDef{
	"marker_1":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_1.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"marker_2":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_2.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"marker_3":   &textureDef{`Interface\TARGETINGFRAME\UI-RAIDTARGETINGICON_3.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"marker_4":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_4.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"marker_5":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_5.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"marker_6":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_6.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"marker_7":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_7.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"marker_8":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_8.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"cross2":     &textureDef{`Interface\RAIDFRAME\ReadyCheck-NotReady.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"check":      &textureDef{`Interface\RAIDFRAME\ReadyCheck-Ready.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"question":   &textureDef{`Interface\RAIDFRAME\ReadyCheck-Waiting.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"targeting":  &textureDef{`Interface\Minimap\Ping\ping5.blp`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"highlight":  &textureDef{`Interface\AddOns\groupauras\assets\alert_circle`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"radius":     &textureDef{`Interface\AddOns\groupauras\assets\radius`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"radius_lg":  &textureDef{`Interface\AddOns\groupauras\assets\radius_lg`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"timer":      &textureDef{`Interface\AddOns\groupauras\assets\timer`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1.1},
	"glow":       &textureDef{`Interface\GLUES\MODELS\UI_Tauren\gradientCircle`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"party":      &textureDef{`Interface\MINIMAP\PartyRaidBlips`, widget.BlendBlend, []float32{0.525, 0.6, 0.04, 0.2}, 1},
	"tank":       &textureDef{`Interface\AddOns\groupauras\assets\roles`, widget.BlendBlend, []float32{0.5, 0.75, 0, 1}, 1},
	"dps":        &textureDef{`Interface\AddOns\groupauras\assets\roles`, widget.BlendBlend, []float32{0.75, 1, 0, 1}, 1},
	"healer":     &textureDef{`Interface\AddOns\groupauras\assets\roles`, widget.BlendBlend, []float32{0.25, 0.5, 0, 1}, 1},
	"fadecircle": &textureDef{`Interface\AddOns\groupauras\assets\fadecircle`, widget.BlendBlend, []float32{0, 1, 0, 1}, 1},
	"ring":       &textureDef{`SPELLS\CIRCLE`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"rune1":      &textureDef{`SPELLS\AURARUNE256.BLP`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"rune2":      &textureDef{`SPELLS\AURARUNE9.BLP`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"rune3":      &textureDef{`SPELLS\AURARUNE_A.BLP`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"rune4":      &textureDef{`SPELLS\AURARUNE_B.BLP`, widget.BlendAdd, []float32{0.032, 0.959, 0.035, 0.959}, 1},
	"paw":        &textureDef{`SPELLS\Agility_128.blp`, widget.BlendAdd, []float32{0.124, 0.876, 0.091, 0.903}, 1},
	"cyanstar":   &textureDef{`SPELLS\CYANSTARFLASH.BLP`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"summon":     &textureDef{`SPELLS\DarkSummon.blp`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"reticle":    &textureDef{`SPELLS\Reticle_128.blp`, widget.BlendAdd, []float32{0.05, 0.95, 0.05, 0.95}, 1},
	"fuzzyring":  &textureDef{`SPELLS\WHITERINGTHIN128.BLP`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"fatring":    &textureDef{`SPELLS\WhiteRingFat128.blp`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
	"swords":     &textureDef{`SPELLS\Strength_128.blp`, widget.BlendAdd, []float32{0, 1, 0, 1}, 1},
}

func init() {
	textureMap["star"] = textureMap["marker_1"]
	textureMap["circle"] = textureMap["marker_2"]
	textureMap["diamond"] = textureMap["marker_3"]
	textureMap["triangle"] = textureMap["marker_4"]
	textureMap["moon"] = textureMap["marker_5"]
	textureMap["square"] = textureMap["marker_6"]
	textureMap["cross"] = textureMap["marker_7"]
	textureMap["skull"] = textureMap["marker_8"]
}
