#!/bin/bash

if [[ $1 == 'primary' ]]; then 
    shift
    echo "primary"
    [[ -z $CONFIG ]] && export CONFIG=/etc/dsip/config_primary.json
    /opt/dsip/primary.run $@
    VALID="ok"

elif [[ $1 == 'secondary' ]]; then
    shift
    echo "secondary"
    [[ -z $CONFIG ]] && export CONFIG=/etc/dsip/config_secondary.json
    /opt/dsip/secondary.run $@
    VALID="ok"

elif [[ $1 == 'client' ]]; then
    shift
    echo "client"
    /opt/dsip/client.run $@
    VALID="ok"

elif [[ $1 == '-help' ]]; then
    shift
    echo "Welcome to dsip"
    echo "You can run 3 diffrents commands:"
    echo ""
    echo "  primary"
    echo "      this will run dsipe primary node"
    echo ""
    echo "  secondary"
    echo "      this will run dsipe secondary node"
    echo ""
    echo "  client"
    echo "      this will run dsipe client node"
    echo ""
    echo "  to get additional info, you can run each command with flag -help"
    echo "    \$ dsip primary -help"
    echo ""
    echo "  to create alias for dsip, type"
    echo "    \$ dsip alias your_alias"
    VALID="ok"

elif [[ $1 == 'alias' ]]; then
    if [[ -n $2 ]]; then
        grep "^$USER" /etc/passwd | grep zsh > /dev/null
        [[ $? == 0 ]] && echo "alias $2=dsip" >> ~/.zshrc && ZSH="ok"

        grep "^$USER" /etc/passwd | grep bash > /dev/null
        [[ $? == 0 ]] && echo "alias $2=dsip" >> ~/.bashrc && BASH="ok"
        
        if [[ -z $BASH && -z $ZSH ]]; then
            echo "Unsupported shell"
        else
            echo "You may need to reload your shell to use $2 as alias for dsip"
        fi
        
        VALID="ok"
    fi
fi

if [[ -z $VALID ]]; then
    echo "This is not a valid command. To get help, run dsip -help"
    exit 1
fi