# Link Shortener ![GO][go-badge]

[go-badge]: https://img.shields.io/badge/Go-v1.27.1-blue

> This project involves creating a basic URL shortener using Go. The purpose is to convert lengthy URLs into shorter links.

## System Architecture
A clear example of using the [clean architecture](https://github.com/evrone/go-clean-template) pattern 

## Build & Run (Locally)
### Prerequisites
- go 1.27.1
- docker & docker-compose

Create .env file in root directory and add following values:
```dotenv
GIN_MODE=debug
HTTP_PORT=8000
PG_URL=postgres://user:pass@postgres:5432/linkShortener
```

Use `make run` run app with migrations

## License
📜 This project is licensed under the [MIT License](LICENSE).
