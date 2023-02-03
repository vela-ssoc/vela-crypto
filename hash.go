package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/vela-ssoc/vela-kit/lua"
	"hash/crc32"
)

func Crc32(L *lua.LState) int {
	h := crc32.NewIEEE()
	s := L.CheckString(1)
	raw := L.CheckBool(2)
	_, err := h.Write([]byte(s))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	var result string
	if !raw {
		result = hex.EncodeToString(h.Sum(nil))
	} else {
		result = string(h.Sum(nil))
	}
	L.Push(lua.LString(result))
	return 1
}

func SHA1(L *lua.LState) int {
	h := sha1.New()
	s := lua.LVAsString(L.Get(1))
	raw := lua.LVAsBool(L.Get(2))
	_, err := h.Write([]byte(s))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	var result string
	if !raw {
		result = hex.EncodeToString(h.Sum(nil))
	} else {
		result = string(h.Sum(nil))
	}
	L.Push(lua.LString(result))
	return 1
}

func MD5(L *lua.LState) int {
	data := L.CheckString(1)
	hash := md5.Sum(lua.S2B(data))
	L.Push(lua.S2L(fmt.Sprintf("%x", hash)))
	return 1
}

func SHA256(L *lua.LState) int {
	data := L.CheckString(1)
	hash := sha256.Sum256(lua.S2B(data))
	L.Push(lua.S2L(fmt.Sprintf("%x", hash)))
	return 1
}
