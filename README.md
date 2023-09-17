# GopherBank: Backend Web Service for Banking Operations

Welcome to GopherBank, a comprehensive backend web service developed in Golang. This project provides APIs for managing bank accounts, recording balance changes, and facilitating money transfers between accounts.

**Tech Stack and Libraries Used:**

- **Programming Language**: Golang
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Database Migration Tool**: golang-migrate
- **Containerization**: Docker
- **Testing**: Unit, Mocking
- **Token-Based Authentication**: PASETO and JWT
- **Password Hashing**: Bcrypt
- **Version Control**: Git
- **Continuous Integration**: GitHub Actions

## Introduction

GopherBank is a backend web service developed in Golang, designed to provide functionalities for a simple banking application.

## Features

### Working with Database [Postgres]

This feature set focuses on working with a PostgreSQL database and covers various aspects of database management and integration.

- **Working with Database Schema**: Design DB schema and generate SQL code with dbdiagram.io.
- **Database Setup with Docker and Postgres**: Install & use Docker + Postgres + TablePlus to create DB schema.
- **Database Migration in Golang**: How to write & run database migration in Golang.
- **CRUD Golang Code Generation**: Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc.
- **Unit Testing for Database CRUD**: Write unit tests for database CRUD with random data in Golang.
- **Database Transactions in Golang**: A clean way to implement database transaction in Golang.
- **Handling DB Transaction Locks**: DB transaction lock & How to handle deadlock in Golang.
- **Avoiding Deadlocks**: How to avoid deadlock in DB transaction? Queries order matters!
- **Transaction Isolation Levels**: Deeply understand transaction isolation levels & read phenomena in MySQL & PostgreSQL.
- **Automated Testing with GitHub Actions**: Setup GitHub Actions for Golang + Postgres to run automated tests.

### Building RESTful HTTP JSON API [Gin]

This feature set dives into building a RESTful HTTP JSON API using the Gin framework in Golang.

- **Implementing RESTful HTTP API**: Implement RESTful HTTP API in Go using Gin.
- **Configuration Handling**: Load config from file & environment variables in Go with Viper.
- **Mocking for Testing**: Mock DB for testing HTTP API in Go and achieve 100% coverage.
- **Money Transfer API**: Implement transfer money API with a custom params validator.
- **User Management**: Add users table with unique & foreign key constraints in PostgreSQL.
- **Error Handling**: How to handle DB errors in Golang correctly.
- **Password Security**: How to securely store passwords? Hash password in Go with Bcrypt!
- **Enhanced Unit Testing**: How to write stronger unit tests with a custom gomock matcher.
- **Token-Based Authentication**: Why PASETO is better than JWT for token-based authentication?
- **Token Creation and Verification**: How to create and verify JWT & PASETO token in Golang.
- **User Authentication API**: Implement login user API that returns PASETO or JWT access token in Go.
- **Middleware and Authorization**: Implement authentication middleware and authorization rules in Golang using Gin.

<!-- ## Database Schema

For details about the database schema, refer to [Database Schema](docs/database-schema.md).

## RESTful HTTP JSON API

For details about the RESTful HTTP JSON API, refer to [API Documentation](docs/api-documentation.md).

## Usage

For information on how to use GopherBank, refer to the [Usage Guide](docs/usage.md).

## Contributing

We welcome contributions! See the [Contribution Guidelines](CONTRIBUTING.md) for details. -->

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as needed.

---
