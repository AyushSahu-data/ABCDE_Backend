# Cart Backend

This repository contains the backend code for the Cart application.

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (v1.16 or later)
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your_username/cart-backend.git
cd cart-backend

2. Install Dependencies

go mod tidy

3. Running the application

go run cmd/main.go
The server will start running on localhost: 8080

4. Testing

go test ./...

**## Endpoints**

POST /users: Creates a new user
POST /users/login: Logs in existing user
POST /items: Creates a new item
POST /carts: Creates a new cart
POST /carts/orders: Converts cart to order
GET /users: Lists all users
GET /items: Lists all items
GET /carts: Lists all carts
GET /orders: Lists all orders
