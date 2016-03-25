-- Package declaration
local _acedb = _G["github.com/eandre/groupauras/pkg/ace/acedb"] or {}
_G["github.com/eandre/groupauras/pkg/ace/acedb"] = _acedb

local AceDB = LibStub("AceDB-3.0")

_acedb.New = function(dbName, defaults, profileName)
    if profileName == "" then
        profileName = true
    end
    return AceDB:New(dbName, defaults, profileName)
end