
from pydantic import Field
from sqlmodel import SQLModel



# table=True: to emphasize that it is a table model
class Hero(SQLModel, table=True):
    id: int | None = Field(default=None, primary_key=True)
    name: str = Field(index=True)
    age: int | None = Field(default=None, index=True)
    secret_name: str