-- lua/set_code.lua
local codeHashKey = KEYS[1]   -- "phone_code"
local cntHashKey = KEYS[2]    -- "cnt"
local fieldPrefix = ARGV[1]   -- "login_sms:12345678922"
local code = ARGV[2]          -- 生成的验证码
local expireSec = ARGV[3]     -- 过期时间（600 秒）

-- 存储验证码（设置过期时间）
redis.call("HSET", codeHashKey, fieldPrefix, code)
redis.call("EXPIRE", codeHashKey, expireSec)

-- 存储错误次数（3次）
redis.call("HSET", cntHashKey, fieldPrefix, "3")
redis.call("EXPIRE", cntHashKey, expireSec)

return 0
