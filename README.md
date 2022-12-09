<p align="left">
  <img title="vessl" src='https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/vessl-banner.png' />
</p>

# Vessl

Vessl is a simple container management UI service focused on deploy and maintain Docker containers running in an IOT device based on Linux OS.
The application allows you to manage resources (containers, images, image repository, users, etc.) through a simple GUI.

## Features

- Manage Containers (Start, Stop, Restart, Remove, Configure, Stats, Logs)
- Launch Containers
- Manage COntainers Template Repository
- Get Statistics from Host Environment
- User Management
- Manage Images, Volumes, Networks

Apps             |  App Repository
:-------------------------:|:-------------------------:
![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/apps.png)  |  ![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/apps_repository.png)

App Launcher             |  Host Stats
:-------------------------:|:-------------------------:
![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/app_launcher.png)  |  ![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/host_stats.png)

Images             |  Volumes             |  Networks
:-------------------------:|:-------------------------:|:-------------------------:
![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/images.png)  |  ![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/volumes.png)  |  ![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/networks.png)

System             |  Users             |  Host Settings
:-------------------------:|:-------------------------:|:-------------------------:
![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/system.png)  |  ![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/users.png)  |  ![](https://raw.githubusercontent.com/cadugrillo/vessl-resources/main/screenshots/host_settings.png)

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
- On your web browser navigate to: https://localhost or https://"your-ip-address" 
- default Username: master
- default Password: cgMaster@3306


#### If running on Linux using /etc/network/interfaces for network configuration, you can manage interfaces from Vessl by adding the following line to the run command
  - -v /etc/network:/etc/network
```
docker run -dp 443:443 --name=vessl --restart=always \
  -v vessl-database:/database \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /etc/network:/etc/network \
  --network=vessl-default vessl/vessl:latest
```

#### If running on Linux, you can reboot and shutdown from Vessl by running this container
```
docker run -d --name=vessl-host-control \
  --restart=always --privileged --pid=host \
  --network=vessl-default vessl/host-control:latest
```
