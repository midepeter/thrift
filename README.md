# Thrift-app
This is backend for a financial application wrriten in golang. It entails all the features of a simple finance application

## Technologies Stack
- Golang - gRPC - PostgreSql - Docker - Terraform - Kubernetes - AWS

## Features
- Deposit -- Deposit funds at any time you like.
- Lock -- Lock will check your Current balance and lock an amount specified which you are not allowed to withdraw for a period of time
- Withdraw money -- Withdraw a specified an amount of money and it effected in your database

## Setup
You can use docker to start the application by running:
```
$ docker-compose up
```

You can also start up the application using the makefile:
```
1. create the .env file using the .env.sample as template
2. populate the .env file with your postgres db credentials
3. run `make run`
```
## Project status
The basic functionality of the project is done but I intend to improve on the project structure in the future and write proper test for the application.

## LICENSE
It is free for use and you can also raise issues faced when building or testing the application.
