# DevPod Go Project with Multiple Services

This is a Go project template that demonstrates integration with multiple services:
- MySQL
- Redis
- Kafka
- LocalStack (AWS Services)

## Prerequisites

- Docker and Docker Compose
- DevPod
- VS Code with Remote Containers extension

## Getting Started

1. Copy the environment file:
   ```bash
   cp .env.example .env
   ```

2. Open the project in DevPod/DevContainer:
   - The development container will automatically be built with all necessary tools
   - Your git and SSH configurations will be mounted from the host
   - ZSH shell is configured as the default shell

3. Start the services:
   ```bash
   docker-compose up -d
   ```

4. Run the example:
   ```bash
   go run main.go
   ```

## Project Structure

- `.devcontainer/` - DevContainer configuration
- `docker-compose.yml` - Service definitions
- `main.go` - Example code demonstrating service connections
- `.env` - Environment variables (create from .env.example)

## Services

- MySQL: Running on port 3306
- Redis: Running on port 6379
- Kafka: Running on port 9092 (with Zookeeper on 2181)
- LocalStack: Running on port 4566

## Development

The development container is configured with:
- Go 1.21
- ZSH with Oh My Zsh
- Git configuration from host
- SSH configuration from host
