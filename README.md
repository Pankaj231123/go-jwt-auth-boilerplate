# go-jwt-auth-boilerplate

A small Go API starter built with Gin and environment-based configuration. The project currently boots an HTTP server, loads settings from a `.env` file, and exposes a simple health check route.

## What is implemented

- Gin HTTP server bootstrap in `main.go`
- Environment-based config loading in `config/config.go`
- `GET /ping` health check endpoint
- Project structure prepared for future auth, middleware, database, models, and JWT work

## Project structure

- `main.go` - application entry point
- `config/` - loads environment variables
- `database/` - database layer placeholder
- `handlers/` - request handlers placeholder
- `middleware/` - middleware placeholder
- `models/` - data models placeholder
- `utils/` - utility helpers placeholder
- `test.http` - quick local request example

## Requirements

- Go 1.25 or later
- A `.env` file in the project root

## Environment variables

The app currently reads these values from `.env`:

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

The server starts on the port defined in `PORT`.

## Test the API

Health check:

```http
GET http://localhost:8080/ping
```

Example response:

```json
{
	"message": "pong"
}
```

## Notes

The repository already has folders and names ready for JWT authentication, database integration, rate limiting, and user management, but those features are not implemented yet.
