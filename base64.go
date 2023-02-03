package crypto

import (
	"encoding/base64"
	"github.com/vela-ssoc/vela-kit/lua"
)

func Base64Encode(L *lua.LState) int {
	data := L.CheckString(1)
	r := base64.StdEncoding.EncodeToString(lua.S2B(data))
	L.Push(lua.S2L(r))
	return 1
}

func Base64Decode(L *lua.LState) int {
	data := L.CheckString(1)
	r, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.B2L(r))
	}

	return 1
}
