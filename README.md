# Link Shortener ![GO][go-badge]

[go-badge]: https://img.shields.io/badge/Go-v1.23.0-blue

Learn More about Links Shortener
A clear example of using the [clean architecture](https://github.com/evrone/go-clean-template) pattern 

## Build & Run (Locally)
### Prerequisites
- go 1.23.0
- docker & docker-compose

Create .env file in root directory and add following values:
```dotenv
GIN_MODE=debug
HTTP_PORT=8000
PG_URL=postgres://user:pass@postgres:5432/linkShortener
```

Use `make run` run app with migrations