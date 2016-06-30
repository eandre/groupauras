-- Package declaration
local _acecomm = _G["github.com/eandre/groupauras/pkg/ace/acecomm"] or {}
_G["github.com/eandre/groupauras/pkg/ace/acecomm"] = _acecomm

local AceComm = LibStub("AceComm-3.0")

_acecomm.SendCommMessage = function(prefix, msg, channel, target, prio, f, key)
    AceComm:SendCommMessage(prefix, msg, channel, target, prio, f, key)
end

_acecomm.RegisterComm = function(prefix, handler)
    AceComm:RegisterComm(prefix, handler)
end
