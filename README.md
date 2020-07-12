# Login Service and OAuth 2.0 Service Workshop
This repository contains source code of Login Provider that created with Go Programming Language and implement Login Service using ORY Kratos (Cloud Native Identity Management) and OAuth 2.0 Service using ORY Hydra (OAuth 2.0 Provider).

# Prerequisite

For run this application, first you need to install `Docker` and `Docker Compose`. This application consists of services such as:

1. Login Provider service (Go source code).
2. ORY Kratos Service (user management service)
3. ORY Hydra Service (OAuth 2.0 Provider)
4. Database Service (postgreqsql server)

Each service above will run as a Docker Service using `docker-compose` command.

# Usage

## A. Run Service
1. Just type this command on terminal (prefer using Linux or Mac OS).
```sh
$ make run
```
2. Open the Login Provider http://127.0.0.1:9000
3. Access http://127.0.0.1:9000/registration for register new user, and http://127.0.0.1:9000/login for login
4. You can create new OAuth 2.0 client on `Credentials` menu.

## B. Remove Service
```sh
$ make down
```

# Contributor

Anggit M Ginanjar - Software Developer <anggit@outlook.com>.