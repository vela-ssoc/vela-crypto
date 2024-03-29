package crypto

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/vela"
)

var xEnv vela.Environment

func WithEnv(env vela.Environment) {
	xEnv = env
	kv := lua.NewUserKV()
	kv.Set("md5", lua.NewFunction(MD5))
	kv.Set("sha1", lua.NewFunction(SHA1))
	kv.Set("sha256", lua.NewFunction(SHA256))
	kv.Set("hmac_sha256", lua.NewFunction(HmacSha256))
	kv.Set("base64_encode", lua.NewFunction(Base64Encode))
	kv.Set("base64_decode", lua.NewFunction(Base64Decode))
	kv.Set("hex", lua.NewFunction(Hex))
	xEnv.Set("crypto", lua.NewExport("vela.crypto.export", lua.WithTable(kv)))
}
