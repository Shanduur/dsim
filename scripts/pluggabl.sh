#!/bin/bash

if [[ $1 == 'primary' ]]; then 
    shift
    echo "primary"
    [[ -z $CONFIG ]] && export CONFIG=/etc/pluggabl/config_primary.json
    /opt/pluggabl/primary.run $@
    VALID="ok"

elif [[ $1 == 'secondary' ]]; then
    shift
    echo "secondary"
    [[ -z $CONFIG ]] && export CONFIG=/etc/pluggabl/config_secondary.json
    /opt/pluggabl/secondary.run $@
    VALID="ok"

elif [[ $1 == 'client' ]]; then
    shift
    echo "client"
    /opt/pluggabl/client.run $@
    VALID="ok"

elif [[ $1 == '-help' ]]; then
    shift
    echo "Welcome to pluggabl"
    echo "You can run 3 diffrents commands:"
    echo ""
    echo "  primary"
    echo "      this will run pluggable primary node"
    echo ""
    echo "  secondary"
    echo "      this will run pluggable secondary node"
    echo ""
    echo "  client"
    echo "      this will run pluggable client node"
    echo ""
    echo "  to get additional info, you can run each command with flag -help"
    echo "    \$ pluggabl primary -help"
    echo ""
    echo "  to create alias for pluggabl, type"
    echo "    \$ pluggabl alias your_alias"
    VALID="ok"

elif [[ $1 == 'alias' ]]; then
    if [[ -n $2 ]]; then
        grep "^$USER" /etc/passwd | grep zsh > /dev/null
        [[ $? == 0 ]] && echo "alias $2=pluggabl" >> ~/.zshrc && ZSH="ok"

        grep "^$USER" /etc/passwd | grep bash > /dev/null
        [[ $? == 0 ]] && echo "alias $2=pluggabl" >> ~/.bashrc && BASH="ok"
        
        if [[ -z $BASH && -z $ZSH ]]; then
            echo "Unsupported shell"
        else
            echo "You may need to reload your shell to use $2 as alias for pluggabl"
        fi
        
        VALID="ok"
    fi
fi

if [[ -z $VALID ]]; then
    echo "This is not a valid command. To get help, run pluggabl -help"
    exit 1
fi