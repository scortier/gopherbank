# GopherBank: Backend Web Service for Banking Operations

## Introduction

Welcome to GopherBank, a comprehensive backend web service developed in Golang. This project provides APIs for managing bank accounts, recording balance changes, and facilitating money transfers between accounts.

Note: Development is in progress...

**Tech Stack and Libraries Used:**

- **Programming Language**: Golang
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Database Migration Tool**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Containerization**: Docker
- **Testing**: Unit, Mocking
- **Token-Based Authentication**: PASETO and JWT
- **Password Hashing**: Bcrypt
- **Version Control**: Git
- **Continuous Integration**: GitHub Actions
- **Database Visualization**: TablePlus
- **Database Compiler**: [Sqlc](https://github.com/sqlc-dev/sqlc)

## Features

### Database Management

- **Schema Design**: Design and generate SQL schema using dbdiagram.io.
- **Containerized Database Setup**: Set up a database using Docker and Postgres for efficient schema creation.
- **Migration Handling**: Write and run migration scripts in Golang for seamless database updates.
- **CRUD Code Generation**: Generate CRUD Golang code from SQL and compare different approaches.
- **Transaction Management**: Implement clean and effective database transactions in Golang.
- **Deadlock Handling**: Address deadlock scenarios and implement strategies to avoid them.
- **Transaction Isolation Understanding**: Gain a deep understanding of transaction isolation levels and read phenomena in databases.

### API Development

- **RESTful API Implementation**: Develop a RESTful HTTP API using Gin in Go.
- **Configuration Management**: Load configuration from files and environment variables using Viper.
- **Money Transfer API**: Implement the money transfer API with a custom parameter validator.
- **User Management Features**: Add users table with unique and foreign key constraints.
- **Error Handling**: Implement robust error handling mechanisms for database errors.
- **Password Security**: Enhance security by securely storing passwords using Bcrypt.
- **Token-Based Authentication**: Explore the advantages of PASETO over JWT for secure token-based authentication.
- **Token Management and Verification**: Learn to create and verify JWT & PASETO tokens effectively.
- **User Authentication API**: Implement a login user API that returns PASETO or JWT access tokens.
- **Middleware and Authorization**: Implement authentication middleware and define authorization rules effectively.

### Testing

- **Unit Testing**: Write comprehensive unit tests for CRUD operations and APIs in Golang.
- **Effective Mocking**: Utilize mocking techniques for comprehensive API testing in Go.

## Contributing

We welcome contributions to GopherBank! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Clone the forked repository to your local machine.
3. Create a new branch for your feature: `git checkout -b feature-name`.
4. Make the necessary changes and commit them: `git commit -m 'Add some feature'`.
5. Push to the branch: `git push origin feature-name`.
6. Submit a pull request to the `main` branch of the original repository.

<!-- For more details, see our [Contribution Guidelines](CONTRIBUTING.md). -->

## Questions and Contact

If you have any questions or need further assistance, please feel free to reach out:

- Email: onlytoaditya@example.com
- Issue Tracker: [GitHub Issues](https://github.com/scortier/gopherbank/issues)

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as needed.

---
