package chatlink

import (
	"github.com/eandre/lunar-wow/pkg/luastrings"
	"github.com/eandre/lunar-wow/pkg/luautil"
	"github.com/eandre/lunar-wow/pkg/widget"
	"github.com/eandre/lunar-wow/pkg/wow"
)

type Listener interface {
	ShowHyperlink(chatFrame widget.Frame, link, text, button string)
	HideHyperlink()
}

var listeners = make(map[string]Listener)

func Register(name string, l Listener) {
	if listeners[name] != nil {
		return
	}
	listeners[name] = l
}

func init() {
	wow.HookSecureFunc("ChatFrame_OnHyperlinkShow", func(chatFrame widget.Frame, link, text, button string) {
		listener := listeners[link]
		for _, l := range listeners {
			if l == listener {
				l.ShowHyperlink(chatFrame, link, text, button)
			} else {
				l.HideHyperlink()
			}
		}
	})

	luautil.HookMethod("ItemRefTooltip", "SetHyperlink", func(orig func(obj widget.Frame, link string, args ...interface{}), obj widget.Frame, link string, args ...interface{}) {
		if !luautil.IsNil(link) {
			for name := range listeners {
				if luastrings.HasPrefix(link, name) {
					return
				}
			}
		}
		orig(obj, link, args...)
	})

	luautil.HookFunc("HandleModifiedItemClick", func(orig func(link string, args ...interface{}), link string, args ...interface{}) {
		for name := range listeners {
			if luastrings.Find(link, "|H"+name+"|h") {
				return
			}
		}
		orig(link, args...)
	})
}
