# Go To Do Demo App for Razorops

If you have forked this repo, then connect with Razorops to create your demo pipeline by following the below button

[![Connect](https://github.com/razorops-public/images/blob/main/connect_with_github.svg)](https://dashboard.razorops.com/get-github-installation-link-for-org)


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
