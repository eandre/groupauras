-- Package declaration
local _serialize = _G["github.com/eandre/groupauras/pkg/serialize"] or {}
_G["github.com/eandre/groupauras/pkg/serialize"] = _serialize

local _errors = _G["errors"]

local type = type

local deepcopy
_serialize.Do = function(obj)
    if type(obj) ~= "table" then
        return nil, errors.New("can only serialize tables")
    end
    return deepcopy(obj), nil
end

_serialize.Undo = function(from, to)
    if type(from) ~= "table" or type(to) ~= "table" then
        return errors.New("can only deserialize tables")
    end
    return to:_initializeFromTable(from)
end

-- Adapted from http://lua-users.org/wiki/CopyTable
-- (same but without metatable copying)
deepcopy = function(orig)
    local orig_type = type(orig)
    local copy
    if orig_type == 'table' then
        copy = {}
        for orig_key, orig_value in next, orig, nil do
            copy[deepcopy(orig_key)] = deepcopy(orig_value)
        end
    else -- number, string, boolean, etc
        copy = orig
    end
    return copy
end
