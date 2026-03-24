# lb

lb is a Layer 4 load balancer written in Go.

## Build

Use the Makefile to build the binary, binaries are built in the `bin/` directory.

```sh
$ make build

$ ls bin
lb  test-client  test-server
```

## Run

The binary is run with a TOML config, this can be specified at runtime.

```bash
$ ./bin/lb --help
Usage of ./bin/lb:
  -config string
        the config path (default "/app/lb.toml")
```

The TOML config has the following structure:

```TOML
lb_port = "8080"   # Port LB binds to and is accessible over
servers = [        # List of upstream servers to balance against
  "server-1:9090",
  "server-2:9090",
  "server-3:9090",
]
```

## Example

You can run an example topology using Docker. This topology runs
a custom test client, against three test servers. The client
initiates requests which are load balanced against the servers.

```bash
$ docker compose up
Attaching to client, lb, server-1, server-2, server-3
server-1  | Server running on port :9090
server-3  | Server running on port :9090
server-2  | Server running on port :9090
lb        | 2026/03/24 20:52:34 INFO Running lb
lb        | 2026/03/24 20:52:34 INFO Listening address=:8080
```