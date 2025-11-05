-- lua/set_code.lua
local codeKey = KEYS[1]   -- "phone_code:login_sms:12345678922"
local cntKey = KEYS[2]    -- "cnt:login_sms:12345678922"
local blackKey = KEYS[3]    -- "black:login_sms:12345678922"
local code = ARGV[1]          -- 生成的验证码 保证60s后发一次
local expireSec = ARGV[2]     -- 过期时间（60 秒）



-- 1, 检测验证码是否存在或过期
--    防止频繁发短信
local exists_code = redis.call("GET", codeKey)
if exists_code then
    return -1  -- 验证码已存在,并且没有过期
end

-- 2, 请求三次都失败,拉进小黑屋一小时
local is_black = tonumber(redis.call("GET", blackKey) or "0")
if is_black >= 3 then
    redis.call("EXPIRE", blackKey, expireSec)
    return -2 -- 黑名单用户,一小时内禁止发送消息
else
    redis.call("SET", blackKey, tostring(is_black+1))
end

-- 3, 存储验证码（设置过期时间）
redis.call("SET", codeKey, code)
redis.call("EXPIRE", codeKey, expireSec)

-- 4, 存储错误次数（3次）
redis.call("SET", cntKey, "3")
redis.call("EXPIRE", cntKey, expireSec)

return 0
