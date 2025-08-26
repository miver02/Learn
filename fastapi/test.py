from jose import jwt
from datetime import datetime, timedelta

# 配置参数
SECRET_KEY = "800043f784d95341c66fabd500335ca3ce7fe4098cfbaedcd05646417d9a3872"
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 30

def create_jwt_token(data: dict):
    # 复制数据以避免修改原始字典
    to_encode = data.copy()
    
    # 设置令牌过期时间
    expire = datetime.utcnow() + timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES)
    to_encode.update({"exp": expire})
    
    # 生成JWT令牌
    encoded_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
    return encoded_jwt

# 加密"123456"，这里将其作为subject存储在令牌中
token = create_jwt_token(data={"sub": "123"})
print("加密后的JWT令牌：")
print(token)
