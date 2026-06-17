# Go Backend – User Management REST API

A RESTful API built with **Go** and **Fiber** to manage users with their date of birth. The API calculates each user's age dynamically based on the stored DOB and supports full CRUD operations with pagination.

## Features

* Create, read, update, and delete users
* Dynamic age calculation from date of birth
* Paginated user listing
* Input validation
* Structured logging with Zap
* PostgreSQL database integration using SQLC

## Tech Stack

* Go
* Fiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator

## Project Structure

```text
cmd/
  server/
    main.go

config/
db/
  migrations/
  sqlc/

internal/
  handler/
  repository/
  routes/
  service/
  middleware/
  logger/
```

## Prerequisites

* Go 1.22 or later
* PostgreSQL (local or Neon)
* Git

## Setup

### Clone the repository

```bash
git clone https://github.com/msranjana/go-backend.git
cd go-backend
```

### Create `.env`

```env
DATABASE_URL=postgresql://username:password@host/database?sslmode=require
PORT=3000
```

For a local PostgreSQL instance:

```env
DATABASE_URL=postgres://postgres:password@localhost:5432/userdb?sslmode=disable
PORT=3000
```

### Install dependencies

```bash
go mod tidy
```

### Run database migration

```bash
psql "<DATABASE_URL>" -f db/migrations/001_create_users.sql
```

### Start the server

```bash
go run ./cmd/server
```

The API will be available at:

```
http://localhost:3000
```

## API Endpoints

### Create User

**POST** `/users`

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Get All Users (with Pagination)

**GET** `/users?page=1&limit=10`

Example:

```
GET /users?page=2&limit=5
```

### Get User by ID

**GET** `/users/{id}`

Example:

```
GET /users/1
```

### Update User

**PUT** `/users/{id}`

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

### Delete User

**DELETE** `/users/{id}`

Example:

```
DELETE /users/1
```

## Sample Response

```json
{
  "data": [
    {
      "id": 1,
      "name": "Alice Updated",
      "dob": "1991-03-15",
      "age": 35
    }
  ],
  "total": 1,
  "page": 1,
  "limit": 10,
  "total_pages": 1
}
```

## Running Tests

```bash
go test ./...
```

## Notes

* Age is calculated dynamically from the stored date of birth.
* Keep your `.env` file private and do not commit it to version control.
* Commit a `.env.example` file instead of your actual credentials.
