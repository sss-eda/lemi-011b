# LEMI-011B
This repository contains acquisition software for the LEMI-011B magnetometer.

The software consists of a server and client side. The idea is that the client service runs on the logging cluster. It reads data from the serial port and forwards it to the server API. The server, in turn, runs remotely and listens on the API for new data. The data is then persisted by the server side software.

## Docker
### Build
Instructions for building the Docker container images are show below:

- To build the **server** container image:
```bash
$ docker build -t lemi011b-server:latest -f build/docker/server/Dockerfile .
```
- To build the **client** container image:
```bash
$ docker build -t lemi011b-client:latest -f build/docker/client/Dockerfile .
```

### Run
Example instructions for running the Docker containers are show below:

- To run a **server** container with dependencies:
```bash
# Start a timescaledb instance.
$ docker run --name timescale -d -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb:latest-pg12
# Log into the database container and create the database
$ docker exec -it timescale bash
    $ psql -U postgres
    $ CREATE DATABASE lemi011b;
    $ exit
$ exit

# Run the server (development)
$ docker run -d -p 8080:8080 -e TIMESCALE_URL="postgres://postgres:password@192.168.0.1:5432/lemi011b" lemi011b-server
```
- To run a **client** container:
```bash
# Run the client and mount the serial port into the container.
$ docker run --privileged -d -e API_URL="http://192.168.0.1:8080" -v /dev/ttyUSB0:/dev/ttyUSB0 lemi011b-client
```

=======
[![Go Report Card](https://goreportcard.com/badge/github.com/sss-eda/lemi-011b)](https://goreportcard.com/report/github.com/sss-eda/lemi-011b)
[![Docker Build CI - Client](https://github.com/sss-eda/lemi-011b/actions/workflows/client.yml/badge.svg?branch=main)](https://github.com/sss-eda/lemi-011b/actions/workflows/client.yml)
[![Docker Build CI - Server](https://github.com/sss-eda/lemi-011b/actions/workflows/server.yml/badge.svg?branch=main)](https://github.com/sss-eda/lemi-011b/actions/workflows/server.yml)

# lemi-011b
Acquisition software for the LEMI-011B magnetometer.