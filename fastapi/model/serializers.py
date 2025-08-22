
from pydantic import BaseModel

from .model import User, Item



# 嵌套响应模型
class UserResponse(BaseModel):
    data: User
    message: str

# 继承嵌套响应模型(返回的参数需要全部实现父类属性)
class UserpriceResponse(User):
    item: Item
    message: str | None