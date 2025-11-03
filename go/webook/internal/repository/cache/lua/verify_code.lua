local key = KEYS[1]
-- 用户输入的code
local expected = ARGV[1]
local code = redis.call("get", key)
local cntKey = key .. ":cnt"
local cnt = tonumber(redis.call("get", cntKey))
if cnt <= 0 then
    -- 超过验证次数
    return -1
elseif code == expected then
    -- 验证成功,删除验证码和计数
    redi.call("del", key)
    return 0
else
    -- 验证失败,计数减一
    redis.call("decr", cntKey)
    return -2
end