# Go - Backend Development

REST API built with Go to manage users with date of birth. Returns age calculated dynamically.

---

## Tech Stack

- **GoFiber** — HTTP framework
- **PostgreSQL** — database
- **SQLC** — DB access layer
- **Uber Zap** — logging
- **go-playground/validator** — input validation

---

## Prerequisites

- Go 1.22+
- PostgreSQL running locally

---

## Setup

### 1. Clone the repo

```bash
git clone https://github.com/msranjana/go-backend.git
cd go-backend
```

### 2. Create the database

```sql
CREATE DATABASE userdb;
```

### 3. Run migration

```bash
psql postgres://user:password@localhost:5432/userdb?sslmode=disable -f db/migrations/001_create_users.sql
```

### 4. Configure environment

```bash
cp .env.example .env
```

Edit `.env`:

```env
DATABASE_URL=postgres://user:password@localhost:5432/userdb?sslmode=disable
PORT=3000
```

### 5. Install dependencies

```bash
go mod tidy
```

### 6. Run the server

```bash
go run ./cmd/server
```

Server starts at `http://localhost:3000`

---

## API Endpoints

### Create User
```
POST /users
```
```json
{ "name": "Alice", "dob": "1990-05-10" }
```

### Get User by ID
```
GET /users/:id
```
Returns user with dynamically calculated `age`.

### Update User
```
PUT /users/:id
```
```json
{ "name": "Alice Updated", "dob": "1991-03-15" }
```

### Delete User
```
DELETE /users/:id
```
Returns `204 No Content`.

### List All Users
```
GET /users
```
Returns all users with `age`.

---

## Run Tests

```bash
go test ./internal/service/
```
