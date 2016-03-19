-- Package declaration
local bridge = _G["github.com/eandre/groupauras/bridge"] or {}
_G["github.com/eandre/groupauras/bridge"] = bridge

local builtins = _G.lunar_go_builtins

local loadstring = loadstring

bridge.Eval = function(str)
    local f, err = loadstring("return " .. str)
    -- TODO setfenv on f
    if err then
        return nil, builtins.create_error(err)
    end
    
    local status, result = pcall(f)
    if not status then
        return nil, builtins.create_error(result)
    elseif type(result) ~= "function" then
        return nil, builtins.create_error("input did not evaluate to function")
    end
    -- TODO setfenv on result
    return result, nil
end