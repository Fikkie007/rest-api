here's the best practice folder structre go, sumber ChatGPT

belajar-golang-rest-api/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── user/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   │
│   ├── product/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   │
│   ├── database/
│   │   └── postgres.go
│   │
│   └── config/
│       └── config.go
│
├── migrations/
│   ├── 001_create_users.sql
│   └── 002_create_products.sql
│
├── go.mod
├── go.sum
└── README.md