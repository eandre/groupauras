package transfer

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/groupauras/pkg/ace/acecomm"
	"github.com/eandre/lunar-wow/pkg/luastrings"
)

var prefix = "GroupAuras"

func Send(aura *aura.Aura, channel string, target interface{}, key interface{}, f acecomm.StatusFunc) {
	revision := luastrings.ToString(aura.Revision)
	msg := "AURA " + aura.ID + " " + revision + " " + aura.Sig + " " + aura.Raw
	acecomm.SendCommMessage(prefix, msg, channel, target, "BULK", f, key)
}

type Listener interface {
	OnAuraReceived(*aura.Aura)
}

var listeners []Listener

func RegisterListener(l Listener) {
	listeners = append(listeners, l)
}

func onMessage(prefix, message, distribution, sender string) {
	if !luastrings.HasPrefix(message, "AURA ") {
		return
	}

	parts := luastrings.Split(" ", message, 5)
	if len(parts) < 5 {
		return
	}
	id := parts[1]
	revision := luastrings.ToInt(parts[2])
	sig := parts[3]
	raw := parts[4]

	a := &aura.Aura{
		ID:       id,
		Revision: revision,
		Sig:      sig,
		Raw:      raw,
	}
	// TODO evaluate raw data
	for _, l := range listeners {
		l.OnAuraReceived(a)
	}
}

func init() {
	acecomm.RegisterComm(prefix, onMessage)
}
