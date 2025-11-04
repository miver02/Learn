-- 适配 Redis Hash 结构的验证码验证脚本
local codeHashKey = KEYS[1]  -- 验证码 Hash 表名（phone_code）
local cntHashKey = KEYS[2]   -- 错误次数 Hash 表名（phone_code:cnt）
local field = ARGV[1]        -- Hash 字段（biz:phone，如 "login_sms:138xxxx8888"）
local inputCode = ARGV[2]    -- 用户输入的验证码

-- 1. 从 Hash 表中获取存储的验证码（HGET 替代 GET）
local code = redis.call("HGET", codeHashKey, field)
if not code then
    return -3  -- 验证码不存在（过期或未生成）
end

-- 2. 从 Hash 表中获取错误次数（默认 0）
local cnt = tonumber(redis.call("HGET", cntHashKey, field) or "0")
if cnt <= 0 then
    return -1  -- 错误次数用尽
end

-- 3. 比对验证码
if code == inputCode then
    -- 验证成功：删除 Hash 表中的对应字段（HDEL 替代 DEL）
    redis.call("HDEL", codeHashKey, field)
    redis.call("HDEL", cntHashKey, field)
    return 0   -- 验证成功
else
    -- 验证失败：错误次数减 1（HINCRBY 替代 DECR，-1 表示减 1）
    redis.call("HINCRBY", cntHashKey, field, -1)
    return -2  -- 验证失败
end