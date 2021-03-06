-- Package declaration
local _acecfg = _G["github.com/eandre/groupauras/pkg/ace/acecfg"] or {}
_G["github.com/eandre/groupauras/pkg/ace/acecfg"] = _acecfg

local AceConfig = LibStub("AceConfig-3.0")
local AceConfigDialog = LibStub("AceConfigDialog-3.0")

_acecfg.RegisterOptionsTable = function(addonName, options, slashCmds)
    AceConfig:RegisterOptionsTable(addonName, options, slashCmds)
end

_acecfg.AddToBlizOptions = function(addonName, name)
    if name == "" then
        name = addonName
    end
    return AceConfigDialog:AddToBlizOptions(addonName, name)
end