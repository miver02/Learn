-- 初始化token, 用于存储jwt
local token = nil
-- 首次请求路由地址
local path = "/users/login"
-- 第一次请求方法
mothod = "POST"

-- 请求头设置
wrk.headers["Content-Type"] = "application/json"
wrk.headers["User-Agent"] = ""

-- 发送第一次请求
request = function()
    body = '{"email": "miver@xx.con", "password": "Ba!123456789"}'
    return wrk.format(mothod, path, wrk.headers, body)
end

-- 获取jwt
response = function(status, headers, body)
    if status == 200 and not token then
        token = headers["X-Jwt-Token"]
        path = "/users/profile"
        method = "GET"
        wrk.headers["Authorization"] = string.format("Bearer %s", token)
    end
end
