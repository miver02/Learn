
SECRET_KEY = "800043f784d95341c66fabd500335ca3ce7fe4098cfbaedcd05646417d9a3872"
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 30

fake_users_db = {
    "miver": {
        "username": "miver",
        "full_name": "Miver Doe",
        "email": "miver@example.com",
        "hashed_password": "$2b$12$ogevmI2W48JrPN88R6iShuDbx6SVZHDxZK4TOlP1gK/mD3yhEs03G",
        "disabled": False,
    },
    "johndoe": {
        "username": "johndoe",
        "full_name": "John Doe",
        "email": "johndoe@example.com",
        "hashed_password": "$2b$12$EixZaYVK1fsbw1ZfbX3OXePaWxn96p36WQoeG6Lruj3vjPGga31lW",
        "disabled": False,
    },
    "alice": {
        "username": "alice",
        "full_name": "Alice Wonderson",
        "email": "alice@example.com",
        "hashed_password": "fakehashedsecret",
        "disabled": True,
    },
}

