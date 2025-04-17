# 🧀 CheddarTrack

**CheddarTrack** is a modern, secure, and extensible backend service built with Golang to help users track their income and expenses. It supports categorized transactions, multi-currency support with exchange rate syncing (coming soon), and secure authentication via JWT tokens.

---

## 🚀 Features

- ✅ User registration & login with hashed passwords
- ✅ JWT-based authentication and protected routes
- ✅ Create and list transactions
- ✅ Categorize transactions (e.g. Food, Rent, Salary)
- ✅ Auto-generated Swagger UI for API documentation
- 🔒 Secure, modular, clean architecture
- ⏳ Currency exchange rate syncing with goroutines (soon)
- 🧪 Built with testability and extensibility in mind

---

## 🧱 Tech Stack

| Layer           | Tech                       |
|-----------------|----------------------------|
| Language        | Go 1.21+                   |
| Framework       | Gin (HTTP router)          |
| ORM             | GORM (PostgreSQL)          |
| Auth            | JWT (`golang-jwt/jwt/v5`)  |
| Hashing         | Bcrypt (`x/crypto/bcrypt`) |
| Docs            | Swagger via `swaggo`       |
| Env Config      | `.env` + Go structs        |

---

## 📁 Project Structure

```
cheddartrack/
├── cmd/api/               # Entry point (main.go)
├── internal/
│   ├── config/            # Env config loader
│   ├── db/                # DB connection + migrations
│   ├── handler/           # HTTP route handlers
│   ├── middleware/        # Auth & request middleware
│   ├── model/             # Data models + DTOs
│   ├── repository/        # DB access logic (GORM)
│   ├── service/           # Business logic
├── docs/                  # Swagger docs (auto-generated)
├── .env                   # Environment variables
├── go.mod / go.sum        # Go modules
```

---

## ⚙️ Setup & Run

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/cheddartrack.git
cd cheddartrack
```

### 2. Configure environment

Create a `.env` file:

```env
PORT=8080
ENV=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=cheddartrack
JWT_SECRET=supersecretjwtkey
BASE_CURRENCY=USD
CURRENCY_API_URL=https://api.exchangerate.host
```

### 3. Start PostgreSQL

You can use Docker:

```bash
docker run --name cheddar-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=yourpassword -e POSTGRES_DB=cheddartrack -p 5432:5432 -d postgres
```

### 4. Install dependencies

```bash
go mod tidy
```

### 5. Generate Swagger docs

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init --generalInfo cmd/api/main.go
```

### 6. Run the server

```bash
go run cmd/api/main.go
```

The server will be running at:

```
http://localhost:8080
```

---

## 📄 API Docs

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

---

## 🔐 Auth Flow

### Register a User

```http
POST /register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "mypassword"
}
```

### Login

```http
POST /login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "mypassword"
}
```

Response:

```json
{
  "access_token": "<JWT_TOKEN>",
  "token_type": "bearer",
  "expires_in": 3600
}
```

Use this token in your Authorization header for protected routes:

```
Authorization: Bearer <JWT_TOKEN>
```

---

## 📌 Coming Soon

- 💸 Currency exchange rate integration + historical rates
- 🔁 Recurring transactions
- 💰 Budget tracking & goals
- 📅 Scheduled jobs with goroutines
- ✅ Unit tests & integration tests

---

## 🤝 Contributing

PRs and ideas welcome! Feel free to fork and contribute.

---

## 📜 License

MIT License

---

Happy Tracking! 🧀
