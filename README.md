# Go DevPod Projects Example

This is an example Go project template that demonstrates integration with multiple services utilizing devcontainers and devpod.

The project includes the following services that are common used in a Go project:
- MySQL 8.0
- Redis 7.0
- Kafka (Confluent Platform 7.3.0)
- LocalStack (AWS Services)

## Prerequisites

- Docker and Docker Compose
- DevPod CLI (https://devpod.sh)
- VS Code with Remote Containers extension

## Getting Started

1. Clone this repository and navigate to the project directory

2. Copy the environment file:
   ```bash
   cp .env.example .env
   ```

3. Start DevPod workspace:
   ```bash
   devpod up .
   ```
   This will:
   - Create a new DevPod workspace
   - Build and start the development container
   - Configure the development environment

4. Open the development container in VS Code.

5. Run the example go application:
   ```bash
   go run main.go
   ```

6. If there is any change, to rebuild the devpod workspace:
```bash
devpod up . --recreate
```

## Project Structure

- `.devcontainer/` - DevContainer configuration
- `docker-compose.yml` - Service definitions
- `main.go` - Main application with service integrations and health check
- `.env` - Environment variables (create from .env.example)
- `scripts/` - Utility scripts (e.g., Kafka topic initialization)

## Services

### Core Services
- MySQL: Running on port 3306
- Redis: Running on port 6379
- Kafka: Running on port 9092 (with Zookeeper on 2181)
- LocalStack: Running on port 4566 (supports S3, SQS, SNS)

### Data Management Tools
- Adminer (MySQL Management): http://localhost:8080
- RedisInsight (Redis Management): http://localhost:8001
- Kafdrop (Kafka Management): http://localhost:9000

## Health Check

The application provides a health check endpoint at `/liveness` that verifies connectivity to:
- MySQL database
- Redis connection
- Kafka message publishing

## Development Environment

The development environment comes batteries included with VS Code and essential plugins for productive Go development:

**Core VS Code Extensions:**
- Go (official Go language support)
- Docker (Docker container management)
- Remote - Containers (DevPod integration)
- GitLens (enhanced Git capabilities)
- EditorConfig (consistent editor settings)
- YAML (YAML file support)
- Markdown All in One (Markdown editing)

and some others plugins for better development experience.

The development container is configured with:
- Go 1.23
- ZSH with Oh My Zsh
- Git (with configuration from host)
- SSH (with configuration from host)
- Docker socket mounted for container management
- AWS CLI
- Terraform, Terragrunt, and TFLint
- Helm, Kubectl, and minikube
- jq, yq

## Environment Variables

Key environment variables for the application is configured in the `.env` file (please check .eenv.example for reference).
