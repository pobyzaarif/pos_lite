# POS_LITE
POS_Lite is sample Point Of Sales system build in golang and using hexagonal architecture

## INTRO
$ cp .env.example .env
$ go run app/main/main.go

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
