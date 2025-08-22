from fastapi import FastAPI, Query, Path, Form, Header, Cookie, Request, UploadFile, File, HTTPException,responses
from fastapi.responses import HTMLResponse
import uvicorn
from  typing import Union, List

from model.model import Item, User, Order
from model.serializers import UserResponse, UserpriceResponse



# 创建fastapi对象app
app = FastAPI()

# api主页
@app.get("/")
async def root():
    return {"message": "Hello World"}

# 声明类型
@app.post("/api/users/login/")
async def login(username: str, password: str) -> dict:
    return {
        "username": username,
        "password": password
    }

# 查询参数:在函数中定义而不再路径中的参数就是查询参数
@app.get("/demo/")
async def demo(name: str, age: int) -> dict:
    return {
        "name": name,
        "age": age
    }

# 设置默认值
@app.get('/demo/page/')
async def demo_page(page: str | None = None, size: Union[int, None] = None):
    return {
        'page': page,
        "size": size
    }

# 查询参数效验
@app.get('/demo/page/check/')
async def demo_page_check(token: str | None = Query(
    default=None,
    title="Token",
    max_length=50, 
    min_length=3,
    description="A token for authentication"
)):
    return {
        "message": token
    }

# 路径参数(需要在函数参数中声明)
@app.get('/project/{project_id}')
async def get_project(project_id: int = 1):
    return {
        'project_id': project_id
    }

# 设置路劲参数的效验规则
@app.get('/env/{env_id}')
async def get_env(env_id: int | None = Path(
    default=..., # 效验时,默认值为空用...
    title="环境Id",
    gt=0, # greater than:>
    le=100 # less than or equal:<=
)):
    return {
        'env_id': env_id
    }

# 表单参数
@app.post('/login/')
async def fixed_pwd(
    password: str = Form(max_length=10, min_length=6)
):
    return {
        "password": password
    }

# json参数
@app.post('/items/')
async def create_item(item: Item):
    return {
        "message": "Item created",
        "item": item.dict()
    }

# json嵌套
@app.post('/orders/')
async def create_order(order: Order):
    return order.dict()

# 请求头参数
@app.get('/header/')
async def read_header(
    session_id: str | None = Header(None),
    theme: str = Header("light")
):
    return {
        "session_id": session_id,
        "theme": theme
    }

# Cookie参数
@app.get('/Cookie/')
async def read_Cookie(
    session_id: str | None = Cookie(None),
    theme: str = Cookie("light")
):
    return {
        "session_id": session_id,
        "theme": theme
    }

# 请求对象的获取
"""
@app.get('/item/header/{id}')
def read_request(
    request: Request,
    id: int = Path(ge = 0),
    name: str = Query(default=''),
    ck: str | None = Cookie(default = None)
):
    return request
"""

# 文件上传
@app.post('/upload-file/')
async def upload_file(file: UploadFile):
    return {
        "filename": file.filename,
        "content_type": file.content_type,
        "file_size": await file.read()
    }

# 多文件上传
@app.post('/upload-multiple-files/')
async def upload_multiple_files(files: List[UploadFile]):
    return [
        {
            "filename": file.filename,
            "content_type": file.content_type
        }
        for file in files
    ]

# 多文件上传 + 表单参数
@app.post('/upload/')
async def upload(
    user_id: str = Form(...),
    description: str = Form(...),
    files: List[UploadFile] | None = None,
):
    return {
        "code": 200,
        "message": "success",
        "user_id": user_id,
        "description": description,
        "filenames": [file.filename for file in files]
    }

# 嵌套响应模型
@app.get('/users/{user_id}', response_model=UserResponse)
async def get_user(user_id: int):
    # 假设通过user_id从数据库中获取到了数据如下
    user = {
        "username": "miver",
        "email": "miver@example.com",
        "age": 15
    }

    return {
        "data": user,
        "message": "hello, miver"
    }

# 嵌套+继承响应模型
@app.get('/users/{user_id}/price/', response_model=UserpriceResponse)
async def get_userprice(user_id: int):
# 假设通过user_id从数据库中获取到了数据如下
    user = {
        "username": "miver",
        "email": "miver@example.com",
        "age": 15
    }
    item = {
        "name": "王二",
        "price": 13546.13,
        "is_offer": True
    }
    
    return {
        "username": user['username'],
        "email": user['email'],
        "age": user['age'],
        "item": item,
        "message": "这是嵌套继承响应模型"
    }

# 自定义响应模型
@app.get('/users/demo-response/', response_class=HTMLResponse)
async def demo_response():
    return """
    <html>
    <h1>这是自定义响应模型:HTMLResponse</h1>
    </html>
    """
    




if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)





