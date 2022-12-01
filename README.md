<p align="left">
  <img title="vessl" src='https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/vessl-banner.png' />
</p>

# Vessl

Vessl is a simple container management UI service focused on deploy and maintain Docker containers running in an IOT device based on Linux OS.
The application allows you to manage resources (containers, images, image repository, users, etc.) through a simple GUI.

## Features

- First
- Second
- Third

## Quick Start

- Prerequisites
  - Docker engine

- How to run
```
docker network create vessl-default

docker run -dp 443:443 --name=vessl --restart=always \
  -v vessl-database:/database \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --network=vessl-default vessl/vessl:latest
```

#### On your web browser navigate to: https://localhost or https://"your-ip-address" 

#### default Username: master
#### default Password: cgMaster@3306