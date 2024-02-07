# Toggl Challenge

This is an API that allows you to do basic operations with a Deck of french cards.

## Structure

```
│
├── cmd/                     # Main applications for this project
│   └── server/              # Entry point of the API server
│       └── main.go          # Initializes and starts the server
|   └── app                  # App builder, used to instantiate the app and its dependencies
│
├── internal/                # Private application and library code
│   ├── api/                 # Route handlers and middleware
│   ├── model/               # Domain models (Deck and Card plain structs/operations)
│   ├── service/             # Business logic
│   └── storage/             # Data storage implementations (focused on extendability)
│
├── pkg/                     # Public library code
│
├── docker-compose.yml       # For easily build and run the application binary in a container
├── Dockerfile               # Dockerfile for containerizing the application
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
└── README.md                # Project overview and setup instructions
```

Most of the packages has an equivalent test file, by exception to the `cmd` package, which is the entry point of the application (and is covered by the tests on `/internal/api`).

## Prerequisites
What things you need to install the software and how to install them, for example:

- [Go](https://go.dev/dl/) - The Go programming language
- [Docker](https://docs.docker.com/get-docker/) - Docker for building and running the application in a container
- [Docker Compose](https://docs.docker.com/compose/install/) (optional) - For easy multi-container orchestration (if applicable)


## How to Run

### Docker compose

1. Build the Docker image:

```sh
docker-compose build
```

2. Run the Docker container:

```sh
docker-compose up
```

The default port where the app is running is `777`, but it's possible to change it in the `docker-compose.yml` file, using the `ports` section + `PORT` env variable.

### Docker (without docker-compose)

1. Build the Docker image:

```sh
docker build -t toggl-challenge .
```

2. Run the Docker container:

```sh
docker run -p 3000:3000 toggl-challenge
```

### Local

1. Run the application:

```sh
go run ./cmd/server/main.go
```

## About me

- [LinkedIn](https://www.linkedin.com/in/ologbonowiwi)
- [GitHub](https://github.com/ologbonowiwi)
- [E-mail](mailto:wm4tos.777@gmail.com)