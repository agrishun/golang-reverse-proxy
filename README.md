
# Reverse proxy
A simple reverse proxy written in GoLang running within docker image

## Requirements
Docker

## Usage
First you need to clone the repo and build docker image
```
git clone https://github.com/agrishun/reverse-proxy-go.git
cd reverse-proxy-go
docker build -t reverse-proxy .  
```

After that you can run the server
```
docker run \
-e HOSTNAME='https://petstore.swagger.io/v2/' \
-e PARAMS='{"status":"sold"}' \
-e PORT=8080 \
-p 4000:8080 reverse-proxy
```

You can test the proxy by running the following command
```
curl -v http://localhost:4000/pet/findByStatus
```