FROM mcr.microsoft.com/devcontainers/go:1.23

# Install additional tools that are not available as features
RUN apt-get update && apt-get install -y \
    default-mysql-client \
    redis-tools \
    kcat \
    && rm -rf /var/lib/apt/lists/*

# Install golangci-lint and gotestsum
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2 \
    && go install gotest.tools/gotestsum@v1.11.0

# Fix permissions for Go module cache
RUN mkdir -p /go/pkg/mod && chown -R vscode:vscode /go/pkg/mod

# Switch to non-root user
USER vscode
