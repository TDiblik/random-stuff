alias vim=nvim
alias ls="ls -a"
alias fsize="sudo du -ha --max-dept=1"
alias cclip="xclip -selection clipboard"
alias digall='function _dig_all_records() { types=("ANY" "A" "AAAA" "CNAME" "MX" "NS" "SOA" "SRV" "TXT"); for type in "${types[@]}"; do echo "===== $type ====="; dig +noall +answer +multiline +besteffort +dnssec $1 "$type"; done; }; _dig_all_records'
alias opened-ports="sudo lsof -nP -iTCP -sTCP:LISTEN"
alias tor-start="sudo systemctl start tor"
alias tor-stop="sudo systemctl stop tor"
alias smbmap='python /usr/bin/smbmap.py "$@"'
alias ssh-fix='ssh-add ~/.ssh/id_ed25519'
alias obs='obs --disable-shutdown-check'
alias nmap='sudo nmap'

export GPG_TTY=$(tty)
export PS1="\[\033[38;5;14m\][\W]\[$(tput sgr0)\]\[\033[38;5;39m\]\$(git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/(\1)/' -e 's/^/ /')\[$(tput sgr0)\] \[$(tput bold)\]\[\033[38;5;5m\]>\[$(tput sgr0)\] \[$(tput sgr0)\]"