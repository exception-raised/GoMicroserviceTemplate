# Microservice Template

# Project Setup

├── .env # Environment configuration
├── go.mod/go.sum # Go module files
├── src/
│ ├── cmd/ # Application entry points
│ │ ├── config/ # Configuration setup
│ │ ├── main.go # Main application entry
│ │ └── server/ # Server execution
│ ├── models/ # Data models
│ ├── server/ # HTTP server implementation
│ ├── service/ # Business logic layer
│ ├── sqlc/ # Database queries (sqlc generated)
│ └── storage/ # Data access layer

Default database is setup as postgres, however the project allows for the databases to be hotswappable.
