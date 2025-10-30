-- 你的验证码在Redis上的key
-- phone_code:login:152......
local key = KEYS[1]
-- 验证次数,我们一个验证码,最多重复三次, 这个记录了验证了几次
-- phone_code:login:152......:cnt
local val = ARGV[1]
-- 过期时间
local ttl = tonumber(redis.call("ttl", key))
if ttl == -1 then
    -- key 存在,但没有过期时间
    return -2
elseif ttl == -2 or ttl < 540 then
    redis.call("set", key, val)
    redis.call("expire", key, 600)
    redis.call("set", cntKey, 3)
    redis.call("expire", cntKey, 600)
    -- 符合预期
    return 0
else
    return -1
end