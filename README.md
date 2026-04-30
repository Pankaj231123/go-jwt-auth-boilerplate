# go-jwt-auth-boilerplate

This is a Go API project built with Gin, PostgreSQL, JWT authentication, and simple rate limiting. I used it as a starter for a basic auth system with register, login, and protected routes.

## What I built

- Loaded application settings from a `.env` file in `config/config.go`
- Connected the app to PostgreSQL and auto-migrated the user table in `database/db.go`
- Created a `User` model in `models/user.go`
- Added register and login handlers in `handlers/auth.go`
- Generated and validated JWT tokens in `utils/jwt.go`
- Added JWT protection for the `/me` route in `middleware/auth.go`
- Added in-memory rate limiting for auth routes in `middleware/ratelimit.go`
- Added a simple health check route at `GET /ping`

## Routes

- `GET /ping` - health check
- `POST /register` - create a new user and return a token
- `POST /login` - log in and return a token
- `GET /me` - protected route that returns the current token payload

## Project structure

- `main.go` - application entry point and route setup
- `config/` - environment configuration
- `database/` - database connection and migration
- `handlers/` - request handlers
- `middleware/` - auth and rate limit middleware
- `models/` - database models
- `utils/` - JWT helpers
- `test.http` - sample requests for local testing

## Requirements

- Go 1.25 or later
- PostgreSQL
- A `.env` file in the project root

## Environment variables

The app reads these values from `.env`:

- `PORT`
- `DB_HOST`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_PORT`
- `JWT_SECRET`
- `JWT_EXPIRY`

## Run locally

```bash
go mod tidy
go run main.go
```

## Sample requests

`test.http` contains example requests for:

- `GET /ping`
- `POST /register`
- `POST /login`
- `GET /me`

## Notes

Rate limiting is currently in-memory, so it resets when the server restarts. The auth flow uses bearer tokens in the `Authorization` header.
