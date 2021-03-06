# Introduction

Many monitoring tools and alert routing systems are already availabile.

Cloud Rover aims to help SREs with the practical Alert Managment, from alert generation to root cause anaysis.

# How to run Cloud Rover as a docker image

## Create Docker Compose file

Create a docker compose file with the below content. 

```
# Use root/example as user/password credentials
version: '3.1'

services:
  cloudrover:
    image: vivaaayi/cloudrover
    restart: always
    environment:
      DD_API_KEY: ${DD_API_KEY:-NOT_SET}
      DD_APP_KEY: ${DD_APP_KEY:-NOT_SET}
      MYSQL_USERNAME: root
      MYSQL_PASSWORD: root
      MYSQL_HOSTNAME: db
    links: 
      - db:db
    ports:
      - 8000:8000

  db:
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root
    ports:
      - 3310:3306
```

## Run the Cloud Rover

```
export DD_API_KEY=REPLACE
export DD_APP_KEY=REPLACE
docker-compose up cloudrover
```

# Create the database

You have to create the Cloud Rover database, when you run the cloudrover for the very first time.

## MySQL WorkBench
1. Connect to the CoudRover with root/root user in port 3310.
2. Execute the scripts from [Database Scripts](https://github.com/vivasaayi/cloudrover/tree/main/database-scripts) folder

# Accessing the Dashboard

You can access the dashboard using http://localhost:8000

# Other Environment Variables

|Name|Description|Default Value|
|---|---|---|
|DD_ALERT_ROVER_SCHDULE|Schedule for Alert Rover|60 seconds|
|DD_TRIGGERED_MONITORS_SCHDULE|Schedule for Triggered Monitors Rover|60 seconds|