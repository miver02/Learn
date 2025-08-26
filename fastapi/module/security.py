from datetime import timedelta, datetime, timezone

from tokenize import Decnumber
from typing import Annotated, Union
from typing_extensions import deprecated

from fastapi import Depends, FastAPI, HTTPException, status
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm
from pydantic import BaseModel
from passlib.context import CryptContext
import jwt

from .serializers import TokenData, UserInfo, UserPwd
from .settings import SECRET_KEY, ALGORITHM, ACCESS_TOKEN_EXPIRE_MINUTES, fake_users_db






"""jwt加密"""
pwd_context = CryptContext(schemes=['bcrypt'], deprecated="auto")

# 提取token的路径
oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token/jwt/")


# 验证用户输入密码与原密码匹配度
def verify_password(plain_password, hashed_password):
    return pwd_context.verify(plain_password, hashed_password)


# 得到hash密码
def get_password_hash(password):
    return pwd_context.hash(password)


# 得到用户信息
def get_user(db, username: str):
    if username in db:
        user_dict = db[username]
        return UserPwd(**user_dict) # pydantic模型默认忽略多余属性
    return None


# 检查是否从伪数据库中查到数据
def authenticate_user(fake_db, username: str, password: str):
    userpwd = get_user(fake_db, username)
    if not userpwd:
        return False
    if not verify_password(password, userpwd.hashed_password):
        return False
 
    return userpwd


# 创建token
def create_access_token(data: dict, expires_delta: timedelta | None = None):
    to_encode = data.copy()
    if expires_delta:
        expire = datetime.now(timezone.utc) + expires_delta
    else:
        expire = datetime.now(timezone.utc) + timedelta(minutes=15)
    to_encode.update({"exp": expire})
    encode_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
    # print(encode_jwt)
    return encode_jwt
                                

# 验证token,查询user
async def get_current_user_jwt(token: Annotated[str, Depends(oauth2_scheme)]):
    credentials_exception = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Invalid authentication credentials",
        headers={"WWW-Authenticate": "Bearer"},
    )
    try:
        payload = jwt.decode(token, SECRET_KEY, algorithms=[ALGORITHM])
        username = payload.get("sub")
        if username is None:
            raise credentials_exception
        token_data = TokenData(username=username)
    except jwt.InvalidTokenError:
        raise credentials_exception
    userpwd = get_user(fake_users_db, username=token_data.username)
    if username is None:
        raise credentials_exception
    return userpwd





"""简化加密"""
def fake_hash_password(password: str):
    return "fakehashed" + password


# 获取用户信息
def fake_decode_token(token):
    # This doesn't provide any security at all
    # Check the next version
    userpwd = get_user(fake_users_db, token)
    return userpwd

# 判断用户是否不存在
async def get_current_user(token: str = Depends(oauth2_scheme)):
    userpwd = fake_decode_token(token)
    if not userpwd:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Invalid authentication credentials",
            headers={"WWW-Authenticate": "Bearer"},
        )
    return userpwd

# 子类当父类用
async def get_current_active_user(current_user: Annotated[UserInfo, Depends(get_current_user_jwt)]): # 声明类型为UserInfo,传递类型为UserPwd,由于是子类,会发生依赖注入,额外保留子类全部字段,
    if current_user.disabled:
        raise HTTPException(status_code=400, detail="Inactive user")
    return current_user


