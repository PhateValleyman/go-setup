# bash completion for setup
_setup()
{
    local cur prev opts commands
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    # List of commands supported by the script
    commands="alias defaults functions paths profile version help"

    # First argument suggestions (commands)
    if [[ ${COMP_CWORD} == 1 ]]; then
        opts="-v --version -h --help -? $commands"
        COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
        return 0
    fi

    # No further options to suggest for now
    COMPREPLY=()
}

# Register the completion function for the 'setup' command
complete -F _setup setup
