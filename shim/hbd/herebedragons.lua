-- Package declaration
local _hbd = _G["github.com/eandre/groupauras/shim/hbd"] or {}
_G["github.com/eandre/groupauras/shim/hbd"] = _hbd

local HBD = LibStub("HereBeDragons-1.0")

_hbd.WorldPosFromZone = function(zx, zy, id, level)
    local x, y = HBD:GetWorldCoordinatesFromZone(zx, zy, id, level)
    return x or 0, y or 0
end

_hbd.ZonePosFromWorld = function(wx, wy, id, level, allowOutOfBound)
    local x, y = HBD:GetZoneCoordinatesFromWorld(wx, wy, id, level, allowOutOfBound)
    return x or 0, y or 0
end

_hbd.WorldDistance = function(inst, srcX, srcY, dstX, dstY)
    local dist = HBD:GetWorldDistance(inst, srcX, srcY, dstX, dstY)
    return dist or 0
end

_hbd.WorldDistanceVector = function(inst, srcX, srcY, dstX, dstY)
    local angle, distance = HBD:GetWorldVector(inst, srcX, srcY, dstX, dstY)
    return distance or 0, angle or 0
end

_hbd.UnitWorldPosition = function(unitID)
    local x, y, i = HBD:GetUnitWorldPosition(unitID)
    return x or 0, y or 0, i or 0
end

_hbd.PlayerWorldPosition = function()
    local x, y, i = HBD:GetPlayerWorldPosition()
    return x or 0, y or 0, i or 0
end

_hbd.PlayerZone = function()
    local id, level = HBD:GetPlayerZone()
    return id or 0, level or 0
end