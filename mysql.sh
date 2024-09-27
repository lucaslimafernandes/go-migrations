#! /bin/bash
sudo docker run --rm -d \
    -e MYSQL_ROOT_PASSWORD=root_password \
    -e MYSQL_DATABASE=mydb \
    -e MYSQL_USER=myuser \
    -e MYSQL_PASSWORD=user_password \
    -p 3306:3306 \
    -d mysql:9.0.1

