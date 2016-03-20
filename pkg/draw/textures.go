package draw

import "github.com/eandre/groupauras/shim/widget"

type textureDef struct {
	Texture   string
	Blend     widget.BlendMode
	TexCoords []float32
}

var textureMap = map[string]*textureDef{
	"diamond":    &textureDef{`Interface\TARGETINGFRAME\UI-RAIDTARGETINGICON_3.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"star":       &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_1.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"circle":     &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_2.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"triangle":   &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_4.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"moon":       &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_5.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"square":     &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_6.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"cross":      &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_7.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"skull":      &textureDef{`Interface\TARGETINGFRAME\UI-RaidTargetingIcon_8.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"cross2":     &textureDef{`Interface\RAIDFRAME\ReadyCheck-NotReady.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"check":      &textureDef{`Interface\RAIDFRAME\ReadyCheck-Ready.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"question":   &textureDef{`Interface\RAIDFRAME\ReadyCheck-Waiting.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"targeting":  &textureDef{`Interface\Minimap\Ping\ping5.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"highlight":  &textureDef{`Interface\AddOns\HudMap\assets\alert_circle`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"radius":     &textureDef{`Interface\AddOns\HudMap\assets\radius`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"radius_lg":  &textureDef{`Interface\AddOns\HudMap\assets\radius_lg`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"timer":      &textureDef{`Interface\AddOns\HudMap\assets\timer`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"glow":       &textureDef{`Interface\GLUES\MODELS\UI_Tauren\gradientCircle`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"party":      &textureDef{`Interface\MINIMAP\PartyRaidBlips`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"tank":       &textureDef{`Interface\AddOns\HudMap\assets\roles`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"dps":        &textureDef{`Interface\AddOns\HudMap\assets\roles`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"healer":     &textureDef{`Interface\AddOns\HudMap\assets\roles`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"fadecircle": &textureDef{`Interface\AddOns\HudMap\assets\fadecircle`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"ring":       &textureDef{`SPELLS\CIRCLE`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"rune1":      &textureDef{`SPELLS\AURARUNE256.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"rune2":      &textureDef{`SPELLS\AURARUNE9.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"rune3":      &textureDef{`SPELLS\AURARUNE_A.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"rune4":      &textureDef{`SPELLS\AURARUNE_B.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"paw":        &textureDef{`SPELLS\Agility_128.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"cyanstar":   &textureDef{`SPELLS\CYANSTARFLASH.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"summon":     &textureDef{`SPELLS\DarkSummon.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"reticle":    &textureDef{`SPELLS\Reticle_128.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"fuzzyring":  &textureDef{`SPELLS\WHITERINGTHIN128.BLP`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"fatring":    &textureDef{`SPELLS\WhiteRingFat128.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
	"swords":     &textureDef{`SPELLS\Strength_128.blp`, widget.BlendBlend, []float32{0, 1, 0, 1}},
}
