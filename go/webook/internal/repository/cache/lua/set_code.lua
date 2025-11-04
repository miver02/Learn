-- 适配 Hash 结构的验证码存储脚本（生成验证码时调用）
local codeHashKey = KEYS[1]  -- phone_code
local cntHashKey = KEYS[2]   -- phone_code:cnt
local field = ARGV[1]        -- biz:phone（如 "login_sms:138xxxx8888"）
local code = ARGV[2]         -- 生成的验证码
local expireSec = 600        -- 过期时间（10分钟）

-- 1. 存储验证码到 Hash 表（HSET 替代 SET）
redis.call("HSET", codeHashKey, field, code)
-- 2. 设置 Hash 表的过期时间（整体过期，而非单个字段）
redis.call("EXPIRE", codeHashKey, expireSec)

-- 3. 存储错误次数（3次）到 Hash 表
redis.call("HSET", cntHashKey, field, 3)
-- 4. 错误次数 Hash 表同步过期时间
redis.call("EXPIRE", cntHashKey, expireSec)

return 0 -- 存储成功