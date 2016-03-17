-- Package declaration
local bridge = _G["github.com/eandre/sbm/groupauras/bridge"] or {}
_G["github.com/eandre/sbm/groupauras/bridge"] = bridge

local f = CreateFrame("Frame")

local eventMap = {}

f:SetScript("OnEvent", function(self, event, ...)
    local m = eventMap[event]
    if m == nil then
        return
    end
    local args = {...}
    for listener in pairs(m) do
        listener(event, args)
    end
end)

bridge.RegisterEvent = function(event, listener)
    event = event:upper()
    local m = eventMap[event]
    if m == nil then
        m = {}
        eventMap[event] = m
    end
    m[listener] = true
    f:RegisterEvent(event)
end

bridge.UnregisterEvent = function(event, listener)
    event = event:upper()
    local m = eventMap[event]
    if m ~= nil then
        m[listener] = nil
        
        -- Don't unregister if table is not empty
        for f in pairs(m) do
            return
        end
        f:UnregisterEvent(event)
    end
end

