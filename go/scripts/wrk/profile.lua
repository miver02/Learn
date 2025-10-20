wrk.method="GET"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["User-Agent"] = "Apifox/1.0.0 (https://apifox.com)"
-- 记得修改这个，你在登录页面登录一下，然后复制一个过来这里
wrk.headers["Authorization"]="Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjA5Mzg4NzcsIlVpZCI6MSwiVXNlckFnZW50IjoiQXBpZm94LzEuMC4wIChodHRwczovL2FwaWZveC5jb20pIn0.OswbZZMT_5p05nj8xGBML4qrakvY-7smXXwzx0IdzfTI5CGFkdlvpEzZ1dTj7qtYFgXQ0VvP9-VuRwtRp0LTeg"