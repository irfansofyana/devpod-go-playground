FROM mcr.microsoft.com/devcontainers/go:1.21

# Install additional OS packages
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    zsh \
    curl \
    git \
    wget \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# Switch to non-root user
USER vscode

# Create .ssh directory with correct permissions
RUN mkdir -p /home/vscode/.ssh \
    && chmod 700 /home/vscode/.ssh

# Install Oh My Zsh only if not already installed
RUN if [ ! -d "$HOME/.oh-my-zsh" ]; then \
    sh -c "$(wget -O- https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended; \
    fi

# Install Zsh plugins if not already installed
RUN if [ ! -d "${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/zsh-autosuggestions" ]; then \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/zsh-autosuggestions; \
    fi && \
    if [ ! -d "${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting" ]; then \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting; \
    fi

# Create or update Zsh configuration
RUN echo 'export ZSH="$HOME/.oh-my-zsh"\n\
ZSH_THEME="robbyrussell"\n\
plugins=(git docker docker-compose kubectl golang zsh-autosuggestions zsh-syntax-highlighting)\n\
source $ZSH/oh-my-zsh.sh\n\
\n\
# Aliases\n\
alias k="kubectl"\n\
alias d="docker"\n\
alias dc="docker-compose"\n\
alias g="git"\n\
\n\
# Go settings\n\
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin\n\
\n\
# History settings\n\
HISTSIZE=10000\n\
SAVEHIST=10000\n\
setopt SHARE_HISTORY\n\
\n\
# Custom key bindings\n\
bindkey "^[[1;5C" forward-word\n\
bindkey "^[[1;5D" backward-word\n\
' > /home/vscode/.zshrc

# Set Zsh as default shell
ENV SHELL=/bin/zsh
