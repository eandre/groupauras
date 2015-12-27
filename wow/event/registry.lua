-- Package declaration
local _event = _G["github.com/eandre/sbm/wow/event"] or {}
_G["github.com/eandre/sbm/wow/event"] = _event

local builtins = _G.lunar_go_builtins

_event.events = {}

_event.Register = function(e, handler)
    local t = _event.events[e]
    if t == nil then
        t = {}
        _event.events[e] = t
    end
    table.insert(t, handler)
end

_event.Trigger = function(e, ...)
    local es = _event.events[e]
    if es == nil then
        return
    end

    for _, handler in ipairs(es) do
        handler(e, ...)
    end
end

