#!/bin/bash

servicetag=dachdev/otus-architect:hw24-front-service
docker build -t $servicetag app/.
docker push $servicetag
