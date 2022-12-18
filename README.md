# go-grpc

> This is a school project for the course "Introduction to Go language", in which we had to implement a simple gRPC server and client. The project consists of a gRPC server in which you can interact with the database by using CRUD routes. There is also an optionnal task to add a JWT authentication, a Swagger and Stream using SSE (Server Sent Events).

## Description

This project is a simple gRPC server & client that allows you to Create, Read, Update and Delete users, products and payments by using the given routes.
There are 3 entities:

- User -> Which contains the user's information
- Product -> List of products that the user can buy
- Payment -> Groups the products that the user has bought

## Routes

Note that the database given in this repository only contains 3 users, 3 products and 3 payments. You can add more by using the given routes. The database is in the "database.db" file. You can use [DB Browser for SQLite](https://sqlitebrowser.org/) to view the database and add more data.

### User routes

- Create a user -> POST /user
- Get a user -> GET /user/:id
- Update a user -> PUT /user/:id
- Delete a user -> DELETE /user/:id

### Product routes

- Create a product -> POST /product
- Get a product -> GET /product/:id
- Update a product -> PUT /product/:id
- Delete a product -> DELETE /product/:id

### Payment routes

- Create a payment -> POST /payment
- Get a payment -> GET /payment/:id
- Update a payment -> PUT /payment/:id
- Delete a payment -> DELETE /payment/:id
- Get all payments -> GET /payments
- Get all payments by product -> GET /payments/product/:id

## Optional tasks (bonus points)

- JWT authentication
- Swagger
- SSE (Server Sent Events)
- Documentation

## Installation

To install the project, you need to install the dependencies:

- [Go](https://golang.org/doc/install)
- [Fiber](https://gofiber.io/)

## Starting the server

To launch the server, you need to run the following command:

```csharp
go run main.go
```
