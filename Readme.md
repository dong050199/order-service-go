# Health Management Service.

Health Management Service provides APIs to websites that manage health and monitoring some users information like weight, fat rates, meal and other information.

>### Where is the data source for this web page.
>You can config your data from devices like smart watch or smart phone. Data will be synchronized with system and then our service will calculate and have some recommendations for you.
>

>### How to use
>You must register your infomation and have acount for access and save records.
Next time you just use the user name or email and password for authentication your information.
>

>### Technologies
>In this project i use golang programing language.
Database MySQL.
> This project provide Swagger UI and 
> front end can use this for integration with backend services.
>

# Prerequisites

Before you continue, ensure you meet the following requirements:

* You have installed docker, docker-compose for run docker-compose file in source code.

# How to run
First of all you need setup environment for running localy.
* Run docker-compose to start MySQL server on docker container.
```
docker-compose start database 
```
* Update go dependency 
```
go mod tidy
```
* Build and run service by make file

```
make build && make run
```
Now all api of service will be displayed in swagger following page
```
localhost:{Port}/swagger/index.html
```
# APIs Provided for this test

* Column page info get infomation for column page with response body in swagger docs.
```
api/v1/page/column
```

* My record page info get infomation for my records page with response body in swagger docs.
```
api/v1/validate/my-record
```

* Top page info get infomation for top page with response body in swagger docs.
```
api/v1/validate/top-page
```

* Loggin or Register for get JWT token
```
api/v1/user/login
api/v1/user/register
```

* And all api for config db.

# Database Schema:

![def]

[def]: db_health_svc.jpg

# TODO List.
* Unit test for all functions on usecases layer.
* Validation input api.
* ................


# Information abour me

>Email: dong01667181618@gmail.com

>Phone number: 0367181618

>LinkeIn: https://www.linkedin.com/in/dong519