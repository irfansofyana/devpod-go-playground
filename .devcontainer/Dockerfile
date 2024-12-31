FROM mcr.microsoft.com/devcontainers/go:1.21

# Install common development tools and dependencies
RUN apt-get update && apt-get install -y \
    zsh \
    wget \
    git \
    curl \
    unzip \
    jq \
    apt-transport-https \
    ca-certificates \
    gnupg \
    lsb-release \
    default-mysql-client \
    redis-tools \
    kcat \
    && rm -rf /var/lib/apt/lists/*

# Install AWS CLI v2
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf aws awscliv2.zip

# Install Helm
RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" \
    && chmod +x kubectl \
    && mv kubectl /usr/local/bin/

# Install k9s
RUN curl -sS https://webinstall.dev/k9s | bash

# Install Terraform
RUN wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list \
    && apt-get update && apt-get install -y terraform \
    && rm -rf /var/lib/apt/lists/*

# Install yq
RUN wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 -O /usr/bin/yq \
    && chmod +x /usr/bin/yq

# Copy Zsh setup files
COPY install-zsh.sh /tmp/install-zsh.sh
COPY .zshrc /home/vscode/.zshrc

# Run Zsh installation script
RUN chmod +x /tmp/install-zsh.sh && \
    /tmp/install-zsh.sh && \
    rm /tmp/install-zsh.sh

# Switch to non-root user
USER vscode

# Create .ssh directory with correct permissions
RUN mkdir -p /home/vscode/.ssh \
    && chmod 700 /home/vscode/.ssh

# Set Zsh as default shell
ENV SHELL=/bin/zsh
