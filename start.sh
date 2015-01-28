#!/bin/bash

docker build -t pewpew .
docker stop pewpew
docker rm pewpew
docker run -v /var/lib/pewpew:/go/src/github.com/avesanen/pewpew/db -p 127.0.0.1:5004:8000 -d --name pewpew pewpew
