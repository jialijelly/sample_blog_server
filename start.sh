#!/bin/bash

if [ "$1" == "run" ]; then
    docker-compose up
	
elif [ "$1" == "test" ]; then
    docker-compose -f docker-compose.test.yml up -d
    sleep 10
    go test ./...
    docker-compose -f docker-compose.test.yml down --remove-orphans
fi

echo "Waiting for keypress to exit..."
read -n 1