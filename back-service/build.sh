#!/bin/bash

servicetag=dachdev/otus-architect:hw24-back-service
docker build -t $servicetag app/.
docker push $servicetag
