# Go To Do App

This a 3 tier **To-Do List** application,  where: 

- Data tier is NoSQL with mongo
- API tier is Golang (exposed to the host on port 8080)

# Getting Started

**1. configure it**

- generate .env file `cp .env.example .env`
- populate env vars `export $(cat .env| xargs)`

**2.a. start it with docker-compose** run `docker-compose up -d` 
- interact directly with API thru localhost:8080 endpoint

# License

MIT License

Copyright (c) 2019 Shubham Chadokar
