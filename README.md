lb -  Layer 4 Load Balancer written in Go


## Run

1. Build the binaries
```makefile
make build
```
will produce 3 binaries:
    - lb 
    - test-server
    - test-client

2. Run the lb binary with the TOML config (sets port to 8080)
```bash
./bin/lb --config=./example/lb.toml
```

3. You are now ready to load balance! 


## Config
TOML helps set the addresses for lb and the test-client/test-servers
The files are located in the example folder
The current ports are for testing  


## Example
Technologies used:
    - Go 1.25.6
    - TOML
    - Makefile
    - Docker

I configured an approachable testing enviroment using Docker.

```docker
Docker compose up
```
- Builds images
- Spins up containers:
    - 3 servers 
    - 1 lb
    - 1 client

- The client will transmit a number 10 times to a server
- The lb will act as a middle man
- The server will increment the recieved number
- The server will respond with the changed number
- The client will recieve the changed number and print it


