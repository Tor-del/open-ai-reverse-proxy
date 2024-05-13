# Go Reverse Proxy Server

This repository contains a Go program that sets up a reverse proxy server. The server forwards incoming HTTP requests to a specified target URL and uses an external proxy for the outgoing requests.

## Features

- Configurable via environment variables
- Removes certain headers from incoming requests before forwarding
- Uses an external proxy server for outgoing requests
- Logs request details for debugging purposes

## Environment Variables

The program uses the following environment variables for configuration:

- `PORT`: The port on which the server will listen (default: `8080`).
- `PROXY_URL`: The URL to which the incoming requests will be proxied (default: `https://api.openai.com`).
- `EXT_PROXY_URL`: The URL of the external proxy server (default: `http://10.0.0.8:8080`).

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [Docker](https://www.docker.com/)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/reverse-proxy-server.git
    cd reverse-proxy-server
    ```

2. Build the Go program:
    ```sh
    go build -o reverse-proxy main.go
    ```

### Usage

#### Running Directly

1. Set the necessary environment variables:
    ```sh
    export PORT=8080
    export PROXY_URL=https://api.openai.com
    export EXT_PROXY_URL=http://10.0.0.8:8080
    ```

2. Run the server:
    ```sh
    ./reverse-proxy
    ```

#### Running with Docker

1. Build the Docker image:
    ```sh
    docker build -t reverse-proxy-server .
    ```

2. Run the Docker container:
    ```sh
    docker run -d -p 8080:8080 \
        -e PORT=8080 \
        -e PROXY_URL=https://api.openai.com \
        -e EXT_PROXY_URL=http://10.0.0.8:8080 \
        reverse-proxy-server
    ```

The server will start listening on the specified port and proxy requests to the configured URL.