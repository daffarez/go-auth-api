# Go Auth API

Simple authentication API built with Go, Chi router, PostgreSQL, JWT, and structured logging.  
This project demonstrates idiomatic Go practices for building a clean, modular backend.

## Features

- User registration (`/auth/register`)
  - Email & password validation
  - Password hashing with bcrypt
- User login (`/auth/login`)
  - JWT token generation
- Protected user route (`/user`)
  - JWT bearer authentication
- Structured JSON logging
  - Middleware logs requests
  - Handler logs errors with context
- Modular, idiomatic Go project structure

## Setup PostgreSQL (locally without docker)

1. Install PostgreSQL

```bash
brew install postgresql@14
```

2. Initialize database (skip if already done)

```bash
initdb /opt/homebrew/var/postgresql@14
```

3. Start PostgreSQL

```bash
brew services start postgresql@14

# run manually with
/opt/homebrew/opt/postgresql@14/bin/postgres -D /opt/homebrew/var/postgresql@14
```

4. Create database and role

```bash
# login psql
psql postgres

# create role
CREATE ROLE daffarez LOGIN PASSWORD 'password123' CREATEDB;

# create database
CREATE DATABASE authdb OWNER daffarez;
```

5. Enable table users

```bash
\c authdb

CREATE TABLE users (
  id TEXT PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);
```

## Environment Variables

Env vars can be found in `.env` file

## Run the API

```bash
go run cmd/server/main.go
```
