
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