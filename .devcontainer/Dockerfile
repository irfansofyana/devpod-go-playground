FROM mcr.microsoft.com/devcontainers/go:1.23

# Install additional tools that are not available as features
RUN apt-get update && apt-get install -y \
    default-mysql-client \
    redis-tools \
    kcat \
    && rm -rf /var/lib/apt/lists/*

# Switch to non-root user
USER vscode
