export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="robbyrussell"
plugins=(git docker docker-compose kubectl golang zsh-autosuggestions zsh-syntax-highlighting)
source $ZSH/oh-my-zsh.sh

# Aliases
alias k="kubectl"
alias d="docker"
alias dc="docker-compose"
alias g="git"

# Go settings
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin

# History settings
HISTSIZE=10000
SAVEHIST=10000
setopt SHARE_HISTORY

# Custom key bindings
bindkey "^[[1;5C" forward-word
bindkey "^[[1;5D" backward-word
