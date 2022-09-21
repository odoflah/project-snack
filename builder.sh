#!/bin/bash

# this script proceed to different actions according to which parameters are entered 
# oncoming - update of token automatically
if [ $# -eq 0 ] # if no argument build and deploy 
then
    echo "build all"
    docker compose -f docker-compose.yaml build
    docker compose -f docker-compose.yaml up &
    curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signup  
    curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signin  
elif [ $1 == "build" ] ; then
    echo "docker build"
    docker compose -f docker-compose.yaml build
elif [ $1 == "up" ] ; then
    echo "docker up"
    docker compose -f docker-compose.yaml up
elif [ $1 == "update" ] ; then
    echo "Update credential"
    curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signup  
    curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signin  
fi