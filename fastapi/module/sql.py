from sqlmodel import SQLModel, Session, create_engine

from settings import sqlite_url, connect_args


# create engine: use it to keep connected to the database
engine = create_engine(sqlite_url, connect_args=connect_args)


# create database and table
def create_db_and_table():
    SQLModel.metadata.create_all(engine)


# get session
def get_session():
    with Session(engine) as session:
        yield session