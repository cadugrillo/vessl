<p align="center">
  <img title="cg-edge-configurator" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-banner.png' />
</p>

CG-EDGE-CONFIGURATOR is a simple container management UI service focused on deploy and maintain Docker containers running in an IOT device based on Debian Linux OS.
The application allows you to manage resources (containers, images, iamge repository, users, etc.) through a simple GUI.  
It contains 7 menu options:


**DASHBOARD - Contains info about the hardware usage.**
<p align="center">
  <img title="cg-edge-dashboard" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-dashboard.png' />
</p>

**APPS - Contains a list of all running containers and options to manage each one of them.**
<p align="center">
  <img title="cg-edge-apps" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-apps.png' />
</p>

**APP REPOSITORY - Contains a list of all available containers that can be installed (the list can be managed using a template JSON file that is lsited at the Settings menu) and options to manage each one of them.**
<p align="center">
  <img title="cg-edge-app-repository" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-app-repository.png' />
</p>

**IMAGES - Contains a list of all available images and options to manage each one of them.**
<p align="center">
  <img title="cg-edge-images" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-images.png' />
</p>

**SETTINGS - Contains all areas that can be configured (Network Settings, App Repository Settings).**
<p align="center">
  <img title="cg-edge-settings" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-settings.png' />
</p>

**USERS - Contains a list of all system users and its access permission.**
<p align="center">
  <img title="cg-edge-users" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-users.png' />
</p>

**SYSTEM - Contains info about the Host System, Containers Engine, Warnings and options for rebbot and shutdown the host system.**
<p align="center">
  <img title="cg-edge-system" src='https://raw.githubusercontent.com/cadugrillo/cg-edge-resources/main/cg-edge-system.png' />
</p>


This App is part of a series of small projects being created during my quest to learn new skills in the modern software development world. It is written mostly using:

- Golang (or Go) programming language. You can find more info at https://go.dev/.
- Angular framework. You can find more info at https://angular.io/.


<h3>This app works in frontend / backend architecture where:</h3>

**cg-edge-conf-api**   - contains an HTTP API endpoint written in Go using Gin framework acting as backend that conects with the Docker Engine Unix socket.
(the backend app can be found at https://hub.docker.com/r/cadugrillo/cg-edge-conf-api).  

**cg-edge-conf** - contains the webpage based on Angular framework acting as the frontend.  
(the frontend app can be found at https://hub.docker.com/r/cadugrillo/cg-edge-conf).


<h3>You can access a running version of the app at:</h3>

**http://cg-edge-aws-01.cadugrillo.com/**

username: cguser  
password: cguse01  


<h1>Where can I find the source code?</h1>

You can fork this git repository.


<h1>How to deploy ?</h1>

The easiest way to deploy is using docker-compose.yml file found in this git repository.

1. Copy the file to your desired folder
2.  From the root of the folder run in a terminal "docker-compose up -d"
3. From your local browser navigate to http://localhost.

