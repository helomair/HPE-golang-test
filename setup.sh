#!/bin/bash

# Functions
errHandle() {
    if [ $? -ne 0 ]; then 
        echo "Some error occured!"
    fi

    if [ ! -z "$1" ]; then
        echo $1
    fi
}

set -e
trap errHandle EXIT

# Check or install golang
echo "Check Golang installed."

if ! command -v go > /dev/null; then
    echo "Golang installed check failed, install it."

    wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin

    # check again
    echo "Check Golang installed again."
    go version
    if ! command -v go > /dev/null; then
        errHandle "Golang install failed!"
    fi
fi

# Install packages using go mod
echo
echo "Download modules"
go mod download

# Check or install docker
echo
echo "Check Docker installed."

if ! command -v docker > /dev/null; then
    echo "Docker installed check failed, install it."

    sudo apt update
    sudo apt install ca-certificates curl gnupg
    sudo install -m 0755 -d /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    sudo chmod a+r /etc/apt/keyrings/docker.gpg
    echo \
    "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt update
    sudo apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

    # check again
    echo "Check Docker installed again."
    docker -v
    if ! command -v docker > /dev/null; then
        errHandle "Docker install failed!"
    fi
fi

# Setup mongodb
MONGODB_CONTAINER="mongo-test"

echo
echo "Create db volumn folder"
mkdir -p mongodb-data

echo
echo "Docker pull mongo image (4.4)"
sudo docker pull mongo:4.4

echo
echo "Run mongo"
if [ ! "$(sudo docker ps -a -q -f name=$MONGODB_CONTAINER)" ]; then
    if [ "$(sudo docker ps -aq -f status=exited -f name=$MONGODB_CONTAINER)" ]; then
        # cleanup
        sudo docker rm $MONGODB_CONTAINER
    fi
    # run container
    sudo docker run --name $MONGODB_CONTAINER -v $(pwd)/mongodb-data:/data/db -d -p 27017:27017 --rm mongo:4.4
fi

sleep 1

echo
echo "Check docker run success"
sudo docker exec $MONGODB_CONTAINER mongo --eval "print(version())"

errHandle "done"