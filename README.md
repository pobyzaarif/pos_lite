# POS_LITE
POS_Lite is sample Point Of Sales system build in golang and using hexagonal architecture

## Intro
### Installation
1. `$ cp .env.example .env`
2. `$ go run app/main/main.go`

### endpoint list
- login

POST localhost:4001/v1/user/login
payload
```
{
    "email": "admin@poslite.admin",
    "password": "admin@poslite.admin"
}
```

open the code in your favorite IDE for more info
