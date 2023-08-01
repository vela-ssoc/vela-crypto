package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/vela-ssoc/vela-kit/lua"
)

func Hex(L *lua.LState) int {
	lv := L.Get(1)
	if lv == nil || lv.Type() == lua.LTNil {
		L.RaiseError("invalid string")
		return 0
	}

	L.Push(lua.S2L(hex.EncodeToString(lua.S2B(lv.String()))))
	return 1
}

func HmacSha256(L *lua.LState) int {
	value := L.Get(1)
	if value == nil || value.Type() == lua.LTNil {
		L.RaiseError("not found value")
		return 0
	}

	secret := L.CheckString(2)
	if len(secret) == 0 {
		L.RaiseError("not found secret")
		return 0
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write(lua.S2B(value.String()))
	L.Push(lua.B2L(h.Sum(nil)))
	return 1
}
