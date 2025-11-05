-- 验证脚本
local codeHashKey = KEYS[1]   -- "phone_code"
local cntHashKey = KEYS[2]    -- "cnt"
local fieldPrefix = ARGV[1]   -- "login_sms:12345678922"
local inputCode = ARGV[2]     -- 用户输入的验证码

-- 1. 获取存储的验证码
local storedCode = redis.call("HGET", codeHashKey, fieldPrefix)
if not storedCode then
    return -3  -- 验证码不存在或已过期
end

-- 2. 检查错误次数
local cnt = tonumber(redis.call("HGET", cntHashKey, fieldPrefix) or "0")
if cnt <= 0 then
    -- 清理数据
    redis.call("HDEL", codeHashKey, fieldPrefix)
    redis.call("HDEL", cntHashKey, fieldPrefix)
    return -1  -- 错误次数用尽
end

-- 3. 验证码比对
if storedCode == inputCode then
    -- 验证成功，清理数据
    redis.call("HDEL", codeHashKey, fieldPrefix)
    redis.call("HDEL", cntHashKey, fieldPrefix)
    return 0  -- 验证成功
else
    -- 验证失败，减少错误次数
    redis.call("HSET", cntHashKey, fieldPrefix, tostring(cnt - 1))
    if cnt == 1 then
        -- 清理数据
        redis.call("HDEL", codeHashKey, fieldPrefix)
        redis.call("HDEL", cntHashKey, fieldPrefix)
        return -1  -- 错误次数用尽
    end
    return -3  -- 验证码错误
end