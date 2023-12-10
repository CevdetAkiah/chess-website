#!/bin/bash


# Source the environment script
if [ -z "$ENV_FILE_SOURCED" ]; then
    source env.sh
    export ENV_FILE_SOURCED=true
    echo "env.sh sourced"
fi


# Run Go program
go run *.go

