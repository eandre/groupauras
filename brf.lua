-- Package declaration
local brf = _G.brf or {}
_G.brf = brf

-- Local declarations
local Blackhand, Start, init, MyFunc

local sbm = _G["sbm"]

local wow = _G["wow"]

Blackhand = {}
brf.Blackhand = Blackhand

Blackhand.Start = function(e)
end

init = function()
	TestWow(5)
	sbm.RegisterEncounter("Blackhand", 1583, function()
		return setmetatable({ ["phase"] = 1 }, {__index=Blackhand})
	end)
end

MyFunc = function()
end
brf.MyFunc = MyFunc

