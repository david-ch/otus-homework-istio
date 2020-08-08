#!/bin/bash

modules="front-service back-service"

for module in $modules; do
(
  cd "$module" || exit
  ./build.sh
)
done
