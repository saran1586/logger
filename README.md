# Employee API with File Logging

A CRUD REST API in Go built with **Gin** and **GORM** on **PostgreSQL**, with a custom logger
that writes request activity to a log file.

## Endpoints

| Method | Path | Action |
|---|---|---|
| `POST` | `/employees` | Create an employee |
| `GET` | `/employees` | List employees |
| `PUT` | `/employees/:id` | Update an employee |
| `DELETE` | `/employees/:id` | Delete an employee |

## Files

- `main.go`: routes and server
- `employee.go`: request handlers
- `database.go`: the `Employee` model and database connection
- `logger.go`: logger setup

## Run

```sh
export DATABASE_URL="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
go run .
```

The server starts on `:8080`.
