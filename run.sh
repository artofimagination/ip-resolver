#!/bin/bash

docker-compose down
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker system prune -f
docker-compose up --build --force-recreate -d ip-resolver
status=$?; 
if [[ $status != 0 ]]; then 
  exit $status; 
fi