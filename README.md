# Go API E-comm

Build API with golang for e-comm process

## Built With

* [Golang](https://godoc.org/) - open source programming language 
* [MySQL](https://godoc.org/github.com/go-sql-driver/mysql) - Golang Package mysql

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

First you need to prepare the database at you server

#  create database
```
CREATE DATABASE ecommerce;
```

# create table
```
CREATE TABLE employees (
   ID INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   RANK           INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL DEFAULT 25500.00,
   BDAY			  DATE DEFAULT '1900-01-01'
);
```

# insert a record
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');
```

### Installing

After we already set the database now we prepare to running the server and try following routes to test the code

running the server : go to the project folder from your cmd and running

```
go run main.go
```

And now you go to your localhost with your browser

```
http://localhost:8080
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```
