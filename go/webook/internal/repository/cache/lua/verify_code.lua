-- 验证脚本
local codeKey = KEYS[1]   -- "phone_code:login_sms:12345678922"
local cntKey = KEYS[2]    -- "cnt:login_sms:12345678922"
local blackKey = KEYS[3]    -- "black:login_sms:12345678922"
local inputCode = ARGV[1]     -- 用户输入的验证码 1, 登录成功清除数据 2, 60s过期清除数据

-- 1. 获取存储的验证码
local storedCode = redis.call("GET", codeKey)
if not storedCode then
    return -3  -- 验证码不存在或已过期
end

-- 2. 检查错误次数
local cnt = tonumber(redis.call("GET", cntKey) or "0")
if cnt <= 0 then
    -- 清理数据 只要错误次数用完,就等着过期吧
    return -1  -- 错误次数用尽
end

-- 3. 验证码比对
if storedCode == inputCode then
    -- 验证成功，清理数据
    redis.call("DEL", codeKey)
    redis.call("DEL", cntKey)
    redis.call("DEL", blackKey)
    return 0  -- 验证成功
else
    -- 验证失败，减少错误次数
    redis.call("SET", cntKey, tostring(cnt - 1))
    if cnt == 1 then
        -- 清理数据
        redis.call("DEL", codeKey)
        redis.call("DEL", cntKey)
        return -1  -- 验证码错误,并且错误次数用尽
    end
    return -3  -- 验证码不存在或已过期
end