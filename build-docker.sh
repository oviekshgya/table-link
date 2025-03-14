#!/bin/bash

rebuild() {
    sudo git pull origin
    echo "pull git"

    sudo docker container stop table-link || true
    echo "Stopped container"

    sudo docker container rm table-link || true
    echo "Removed container"

    sudo docker image rm table-link_app:latest || true
    echo "Removed image"

    sudo docker-compose up --build --remove-orphans -d
    echo "Docker compose up complete"

    if [ $? -eq 0 ]; then
        echo "Docker Compose successfully started."
    else
        echo "Error starting Docker Compose."
        exit 1
    fi

}

update() {
    sudo docker-compose down
    echo "Docker Compose down complete"

    sudo docker-compose up -d
    echo "Docker Compose up complete"
}

# Cek parameter
if [ "$1" == "rebuild" ]; then
    rebuild
elif [ "$1" == "update" ]; then
    update
else
    echo "Invalid parameter. Use 'rebuild' or 'update'."
    exit 1
fi
