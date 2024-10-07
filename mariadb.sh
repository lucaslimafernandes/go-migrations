#! /bin/bash
sudo docker run --rm \
    --env MARIADB_USER=myuser \
    --env MARIADB_PASSWORD=password \
    --env MARIADB_DATABASE=mydb \
    --env MARIADB_ROOT_PASSWORD=root_password  \
    -d -p 3306:3306 mariadb:11.5.2-ubi
