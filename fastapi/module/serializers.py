
from typing import Union
from pydantic import BaseModel, Field, EmailStr

# 定义json数据模型 -> /items/
class Item(BaseModel):
    name: str
    price: float
    is_offer: bool | None = None


# 模型嵌套 + 类型效验 -> /orders/
class User(BaseModel):
    username: str = Field(min_length=3, max_length=30)
    email: EmailStr
    age: int | None = Field(None, gt=14, le=65)

class Order(BaseModel):
    id: int
    user: User
    items: list[str]
    total: float


# 嵌套响应模型
class UserResponse(BaseModel):
    data: User
    message: str

# 继承嵌套响应模型(返回的参数需要全部实现父类属性)
class UserpriceResponse(User):
    item: Item
    message: str | None


# 用户密码
class UserInfo(BaseModel):
    username: str
    email: Union[str, None] = None
    full_name: Union[str, None] = None
    disabled: Union[bool, None] = None
    # model_config = {"extra": "forbid"}  # 不能给 UserInfo 实例添加额外属性,，但无法「剥离子类实例自带的字段」。


class UserPwd(UserInfo):
    hashed_password: str

class Token(BaseModel):
    access_token: str
    token_type: str


class TokenData(BaseModel):
    username: str | None = None


