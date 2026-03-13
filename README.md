lb -  Layer 4 Load Balancer written in Go


## Run

in the parent directory: 

```makefile
make build
```

will produce 3 binaries:
    - lb 
    - test-server
    - test-client

run the lb directory


## General Info
Technologies used:
    - Go 1.25.6
    - TOML
    - Makefile
    - Docker


The load balancer is ```lb.go``` 
I built a test-client and test-server for ease of testing. 



## Build - Docker
I configured an approachable testing enviroment using Docker.

```docker
Docker compose up
```
- Spins up:
    - 3 server Containers
    - 1 load balancer
    - 1 client

- The client will transmit a number 10 times to the server
- The server will increment the recieved number
- The server will respond with the changed number
- The client will recieve the number and print it


## Build - no Docker
A Makefile is available to create binaries of the test-client/server and lb.

1.  Start each server instance, change the respective TOML port to a different port
2. Start the Load Balancer 
3. Start the Client 



