# POS_LITE
POS_Lite is a sample Point Of Sales system built with golang, using hexagonal architecture and custom log.
This project aims to prove that we can use a custom logger without additional application performance management or monitoring tools to trace, analyze, etc with our lovely apps. just read the log, it's clear enough.

## Installation
1. `$ cp .env.example .env`
2. `$ go run app/main/main.go`

## Other
- [Postman collection](https://www.getpostman.com/collections/1de1fc55c2e35e56d0b3)
- [Sample log](/docs/sample_log.md)
- Sample dashboard after i processed the log with elk
![dashboard](/docs/dashboard.png)
- process log (1 flow of process) discovery by `tracker_id`
![dashboard](/docs/process%20log%20with%20tracker.png)