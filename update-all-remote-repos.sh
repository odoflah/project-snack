#!/bin/bash

HOST=gcr.io
PROJECT_NAME=abiding-cedar-359014
IMAGES=( microservice-app-template-gateway microservice-app-template-auth microservice-app-template-greeting microservice-app-template-frontend )

# Tag images appropriatly 
for i in "${IMAGES[@]}"
do
    docker tag $i $HOST/$PROJECT_NAME/$i
done

# Push images
for i in "${IMAGES[@]}"
do
    docker push $HOST/$PROJECT_NAME/$i
done