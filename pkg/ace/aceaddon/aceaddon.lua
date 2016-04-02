-- Package declaration
local _aceaddon = _G["github.com/eandre/groupauras/pkg/ace/aceaddon"] or {}
_G["github.com/eandre/groupauras/pkg/ace/aceaddon"] = _aceaddon

local AceAddon = LibStub("AceAddon-3.0")

_aceaddon.New = function(name, obj)
    AceAddon:NewAddon(obj, name)
end