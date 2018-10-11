# Exchange Currency

This project is build for self learning process, it uses SOLID principles to structure the project.

# Files

Run `docker-compose up` and it will do the rest

go to browser and enter `localhost:3000/currency`
> For html part it is just use to connect and test with the backend, therefore the code is messy.

# Project Structure
```
- api/
- frontend/
- models/
- pkg/
- sql/
- vendor/

```
### API/
This folder will consists of all of the handler that will be served in main.go
### frontend/
This folder will consists of html files (frontend side)
### models/
Entity of the struct that represent the table from database
### pkg/
in the future if there is a another package it should be there, and inside the package should define the interface for usecase and repository
### sql/
sql for this project, already indexed so the result wont be duplicated and fast
### vendor/
go lang library that is used in this project

# Testing

Only put 2 example of mock test that is done using gomock, it is easy to test the code since we can mock the interface function




# API DOC
**GET** /exchange?from=USD&to=IDR&date=2018-09-03  
To get the currency rate if date is not pass default is today



|Request|Type|
|--|--|
|From|string(3)|
|To|string(3)|
|Date| YYYY-MM-DD (Optional)

Response 
```
{
    "data": {
        "id": 1,
        "from": "",
        "to": "IDR",
        "date": "2018-09-03T00:00:00Z",
        "rate": 13000,
        "week_rate": 0,
        "variance": 0
    }
}
```
week_rate & variance (from one week) is depends on the data, if not found then 0 by default

**POST** /exchange  

to convert the currency rate

|Request|Type|
|--|--|
|From|string(3)*|
|To|string(3)*|
|Date| YYYY-MM-DD*
|Rate| float*

Response 
```
{
    "data": {
        "status": "OK",
        "message": "Data has been successfuly created"
    }
}
```

**POST** /exchange_list  
to add watch list

|Request|Type|
|--|--|
|From|string(3)*|
|To|string(3)*|

Response 
```
{
    "data": {
        "status": "OK",
        "message": "Data has been successfuly created"
    }
}
```
 **GET** /exchange_list
To get the all watch list


Response 
```
{
    "data": [
        {
            "id": 1,
            "from": "IDR",
            "to": "INR"
        },
        {
            "id": 2,
            "from": "USD",
            "to": "EUR"
        },
        {
            "id": 3,
            "from": "USD",
            "to": "IDR"
        }
    ]
}

```
**Delete** /exchange_list/{ID} 
Response
```
{
    "data": {
        "status": "OK",
        "message": "Data has been successfuly deleted"
    }
}
```


