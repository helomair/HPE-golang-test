# Simple line bot

## Introduction
This bot recieves messages from Line webhook (from user), save the message & user infos to MongoDB, and reply the same messages back to user.

It also can broadcast messages to all users that add this bot as friend.

## Configuration
**This script will use *sudo* when setup dockers**
```sh
$ bash setup.sh
```

**In folder configs, rename config.json.example to config.json**
```sh
$ mv configs/config.json.example to configs/config.json
```

**Set Line developer infos, ex. :**
```json
"LineInfo": {
    "id": "20....",
    "secret" : "aeaf08b0119....",
    "accesstoken" : "ftgbhOvNEjVcLPKkVc4RAnVpzqn8..."
}
```

**After setup, run server**
```sh
$ go run main.go
```

**Use ngrok to generate a https endpoint, or use others**
```sh
$ ngrok http 8080
```

**Go to Line developers console, setup webhook**
```sh
Webhook endpoint : https://<ngrok-generated-url>/line-message-webhook
```

### What setup.sh do
1. Check Golang installed, if not, install it
```sh
if ! command -v go > /dev/null; then
 ...install golang
fi
```
2. Download go modules
```sh
go mod download
```
3. Check Docker installed, if not, install it
```sh
if ! command -v docker > /dev/null; then
 ...install docker, will use sudo
fi
```
4. Pull MongoDB image and run it
```sh
# container name
MONGODB_CONTAINER="mongo-test" 

# mongodb data folder
mkdir -p mongodb-data 

# bind port 27017, remove when stopped
sudo docker run --name $MONGODB_CONTAINER -v $(pwd)/mongodb-data:/data/db -d -p 27017:27017 --rm mongo:4.4
```

## APIs
```sh
GET "/" : index, only for check server is running

POST "/line-message-webhook" : Line webhook endpoint

POST "/broadcast" : Broadcast message to all users that add this bot to friend
    @param message String
```