# 📣 E-Commerce Notification Service

A real-time notification microservice built for an e-commerce system. This service consumes messages from RabbitMQ when order status changes and notifies users via WebSocket connections.

---

## 🚀 Features

- Consumes `order.*` events from RabbitMQ
- Sends real-time notifications to clients using WebSocket
- Maintains active WebSocket connections for users
- Built using Fiber and Clean Architecture
- Stores notification logs in MySQL

---

## 🧰 Tech Stack

- **Language**: Golang (Go 1.21+)
- **Framework**: [Fiber](https://gofiber.io/)
- **WebSocket**: [gofiber/websocket](https://github.com/gofiber/websocket)
- **Architecture**: Clean Architecture
- **Database**: MySQL
- **Message Queue**: RabbitMQ (consumer)
- **Containerization**: Docker + Docker Compose

---

## 📁 Project Structure

```md
├───config
├───database
├───internal
│   ├───delivery
│   │   ├───http_handler
│   │   └───ws
│   ├───domain
│   ├───dto
│   ├───infrastructure
│   │   ├───database
│   │   ├───logger
│   │   └───rabbitmq
│   ├───migration
│   ├───repository
│   └───usecase
│       └───broadcaster
├───pkg
├───server
└───tmp
```
## Getting Started
### Prerequisites
- Docker
- Go 1.21+
- <a href="https://github.com/adzi007/ecommerce-order-service" target="_blank">Ecommerce Order Service</a>
- RabbitMQ Container in the same network

## Running Locally (Docker)

1. Clone the project

```bash
git clone https://github.com/adzi007/ecommerce-notification-service.git
cd ecommerce-notification-service
```

2. CD into the ecommerce-notification-service directory and create an .env file or edit from .env.example following with fields bellow

```
DB_HOST=ecommerce-notification-db
DB_PORT=3306
DB_USERNAME=YOUR_DB_USERNAME
DB_PASSWORD=YOUR_DB_PASSWORD
DB_NAME=ecommerce_app
PORT_APP=5003
API_GATEWAY=
URL_PRODUCT_SERVICE=
RABBITMQ_HOST_URL=rabbitmq-local
RABBITMQ_PORT=5672
RABBITMQ_USER=YOUR_RABBITMQ_USER
RABBITMQ_PASSWORD=YOUR_RABBITMQ_PASSWORD
RABBITMQ_VIRTUAL_HOST=ecommerce_development
```

3. Build container

```
docker-compose up --build
```

The App will be running at `http://localhost:5003`

WebSocket endpoint at `ws://localhost:5003/ws/notification/:userId`

## Database Migration
1. Execute migration database
```
docker exec -it ecommerce-notification-service /migrate
```

## API Documentation

<a href="https://www.postman.com/team-ninja-8073/ecommerce-notification-service/overview" target="_blank">Postman Collections</a>