#crypto
内置hash和加密算法

##crypto.md5(string) string
md5哈希算法
```lua
    local v = crypto.md5("hello")
```

##crypto.SHA256(string) string
SHA256哈希算法
```lua
    local v = crypto.SHA256("helo")
```

## crypto.Crc32(data string , raw bool)  string , err
crc32算法
```lua
    local v , err = crypto.Crc32("helo" , true)  --返回原始算出的原始raw字符
    local v , err = crypto.Crc32("helo" , false) --返回hex后的字符
```

## crypto.SHA1(string) string 
```lua
    local v = crypto.SHA("helo")
```

## crypto.Base64Encode(string) value  
base64加密算法
```lua
    local v = crypto.Base64Encode("helo") -- nil 说明加密失败
```

## crypto.Base64Decode(string) value
base64解密算法
```lua
    local v = crypto.Base64Decode("helo") -- nil说明解密失败
```


