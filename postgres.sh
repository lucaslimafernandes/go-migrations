#! /bin/bash
sudo docker run --rm -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:14-alpine
