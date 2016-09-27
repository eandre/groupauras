package classutil

import "github.com/eandre/lunar-wow/pkg/wow"

type SpecInfo struct {
	Name         string
	Class        string
	Melee        bool
	HasInterrupt bool
}

var UnknownSpec = &SpecInfo{
	Name:         "Unknown",
	Class:        "UNKNOWN",
	Melee:        false,
	HasInterrupt: false,
}

var Specs = map[wow.SpecID]*SpecInfo{
	0:   UnknownSpec,
	62:  &SpecInfo{"Arcane", "MAGE", true, false},
	63:  &SpecInfo{"Fire", "MAGE", true, false},
	64:  &SpecInfo{"Frost", "MAGE", true, false},
	65:  &SpecInfo{"Holy", "PALADIN", false, true},
	66:  &SpecInfo{"Protection", "PALADIN", true, true},
	70:  &SpecInfo{"Retribution", "PALADIN", true, true},
	71:  &SpecInfo{"Arms", "WARRIOR", true, true},
	72:  &SpecInfo{"Fury", "WARRIOR", true, true},
	73:  &SpecInfo{"Protection", "WARRIOR", true, true},
	102: &SpecInfo{"Balance", "DRUID", false, false},
	103: &SpecInfo{"Feral", "DRUID", true, true},
	104: &SpecInfo{"Guardian", "DRUID", true, true},
	105: &SpecInfo{"Restoration", "DRUID", false, false},
	250: &SpecInfo{"Blood", "DEATHKNIGHT", true, true},
	251: &SpecInfo{"Frost", "DEATHKNIGHT", true, true},
	252: &SpecInfo{"Unholy", "DEATHKNIGHT", true, true},
	253: &SpecInfo{"Beast Mastery", "HUNTER", true, false},
	254: &SpecInfo{"Marksmanship", "HUNTER", true, false},
	255: &SpecInfo{"Survival", "HUNTER", true, false},
	256: &SpecInfo{"Discipline", "PRIEST", false, false},
	257: &SpecInfo{"Holy", "PRIEST", false, false},
	258: &SpecInfo{"Shadow", "PRIEST", false, false},
	259: &SpecInfo{"Assassination", "ROGUE", true, true},
	260: &SpecInfo{"Combat", "ROGUE", true, true},
	261: &SpecInfo{"Subtlety", "ROGUE", true, true},
	262: &SpecInfo{"Elemental", "SHAMAN", true, false},
	263: &SpecInfo{"Enhancement", "SHAMAN", true, true},
	264: &SpecInfo{"Restoration", "SHAMAN", true, false},
	265: &SpecInfo{"Affliction", "WARLOCK", false, false},
	266: &SpecInfo{"Demonology", "WARLOCK", false, false},
	267: &SpecInfo{"Destruction", "WARLOCK", false, false},
	268: &SpecInfo{"Brewmaster", "MONK", true, true},
	269: &SpecInfo{"Windwalker", "MONK", true, true},
	270: &SpecInfo{"Mistweaver", "MONK", true, false},
	577: &SpecInfo{"Havoc", "DEMONHUNTER", true, true},
	581: &SpecInfo{"Vengeance", "DEMONHUNTER", true, true},
}
