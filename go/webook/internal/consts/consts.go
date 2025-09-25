package consts

// 验证邮箱密码
const (
	EmailRegexPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	PasswordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]).{8,}$`
)

// 数据库配置
const (
	// mysql
	MysqlAddr 		= `127.0.0.1:3306`
	MysqlUser	 	= `root`
	MysqlPassword  	= `root`
	MysqlDatabase		= `webook`

	// redis
	RedisAddr 		= `127.0.0.1:6379`
	RedisUser		= ``
	RedisPassword  	= `redis`

	KeyPairs		= `secret`
)
