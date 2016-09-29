package classutil

import "github.com/eandre/lunar-wow/pkg/wow"

const (
	RoleDamager RoleType = "DAMAGER"
	RoleHealer  RoleType = "HEALER"
	RoleTank    RoleType = "TANK"
)

type RoleType string

type SpecInfo struct {
	Name         string
	Class        string
	HasInterrupt bool
	Melee        bool
	Role         RoleType
}

var UnknownSpec = &SpecInfo{
	Name:         "Unknown",
	Class:        "UNKNOWN",
	Melee:        false,
	HasInterrupt: false,
	Role:         RoleDamager,
}

var Specs = map[wow.SpecID]*SpecInfo{
	0:   UnknownSpec,
	62:  &SpecInfo{"Arcane", "MAGE", true, false, RoleDamager},
	63:  &SpecInfo{"Fire", "MAGE", true, false, RoleDamager},
	64:  &SpecInfo{"Frost", "MAGE", true, false, RoleDamager},
	65:  &SpecInfo{"Holy", "PALADIN", false, true, RoleHealer},
	66:  &SpecInfo{"Protection", "PALADIN", true, true, RoleTank},
	70:  &SpecInfo{"Retribution", "PALADIN", true, true, RoleDamager},
	71:  &SpecInfo{"Arms", "WARRIOR", true, true, RoleDamager},
	72:  &SpecInfo{"Fury", "WARRIOR", true, true, RoleDamager},
	73:  &SpecInfo{"Protection", "WARRIOR", true, true, RoleTank},
	102: &SpecInfo{"Balance", "DRUID", false, false, RoleDamager},
	103: &SpecInfo{"Feral", "DRUID", true, true, RoleDamager},
	104: &SpecInfo{"Guardian", "DRUID", true, true, RoleTank},
	105: &SpecInfo{"Restoration", "DRUID", false, false, RoleHealer},
	250: &SpecInfo{"Blood", "DEATHKNIGHT", true, true, RoleTank},
	251: &SpecInfo{"Frost", "DEATHKNIGHT", true, true, RoleDamager},
	252: &SpecInfo{"Unholy", "DEATHKNIGHT", true, true, RoleDamager},
	253: &SpecInfo{"Beast Mastery", "HUNTER", true, false, RoleDamager},
	254: &SpecInfo{"Marksmanship", "HUNTER", true, false, RoleDamager},
	255: &SpecInfo{"Survival", "HUNTER", true, false, RoleDamager},
	256: &SpecInfo{"Discipline", "PRIEST", false, false, RoleHealer},
	257: &SpecInfo{"Holy", "PRIEST", false, false, RoleHealer},
	258: &SpecInfo{"Shadow", "PRIEST", false, false, RoleDamager},
	259: &SpecInfo{"Assassination", "ROGUE", true, true, RoleDamager},
	260: &SpecInfo{"Combat", "ROGUE", true, true, RoleDamager},
	261: &SpecInfo{"Subtlety", "ROGUE", true, true, RoleDamager},
	262: &SpecInfo{"Elemental", "SHAMAN", true, false, RoleDamager},
	263: &SpecInfo{"Enhancement", "SHAMAN", true, true, RoleDamager},
	264: &SpecInfo{"Restoration", "SHAMAN", true, false, RoleHealer},
	265: &SpecInfo{"Affliction", "WARLOCK", false, false, RoleDamager},
	266: &SpecInfo{"Demonology", "WARLOCK", false, false, RoleDamager},
	267: &SpecInfo{"Destruction", "WARLOCK", false, false, RoleDamager},
	268: &SpecInfo{"Brewmaster", "MONK", true, true, RoleTank},
	269: &SpecInfo{"Windwalker", "MONK", true, true, RoleDamager},
	270: &SpecInfo{"Mistweaver", "MONK", true, false, RoleHealer},
	577: &SpecInfo{"Havoc", "DEMONHUNTER", true, true, RoleDamager},
	581: &SpecInfo{"Vengeance", "DEMONHUNTER", true, true, RoleTank},
}
