package lb

import (
	"io"
	"log/slog"
	"net"
)

type LbConfig struct {
	LbPort  string   `toml:"lb_port"`
	Servers []string `toml:"servers"`
}

func StartLoadBalancer(cfg LbConfig) error {
	// Sets up a socket for lb to listen on,
	// for incoming connections
	address := ":" + cfg.LbPort
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()
	slog.Info("Running lb")

	slog.Info("Listening", slog.String("address", address))

	// Counts the number of total connections,
	// used for round robin load balancing
	robin := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			// TODO What happens if the TCP socket is closed for good?
			// handle different errs, which ones do we break with?
			slog.Error("failed accepting listener", slog.Any("error", err))
			continue
		}
		slog.Info("Accepted connection", slog.Any("address", conn.RemoteAddr()))
		go handleConnection(conn, cfg.Servers[robin%len(cfg.Servers)])
		robin += 1
	}
}

func handleConnection(clientConn net.Conn, serverAddress string) {
	// Ensure the client connection is closed even if
	// the function exits early, e.g. if we fail to
	// connect to the server
	defer clientConn.Close()

	// Connect to the server
	serverConn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		slog.Error("Failed to connect to server", slog.Any("error", err))
		return
	}
	defer serverConn.Close()

	logger := slog.With(
		slog.Any("client_address", clientConn.RemoteAddr()),
		slog.Any("server_address", serverConn.RemoteAddr()))

	txBuffer := make([]byte, 1024)
	rxBuffer := make([]byte, 1024)

	// TODO what happens if one of the goroutines fail

	go func() {
		for {
			// Client -> Server
			n, err := clientConn.Read(txBuffer)
			if err != nil && err != io.EOF {
				logger.Error("Failed reading from client", slog.Any("error", err))
				return
			}
			_, err = serverConn.Write(txBuffer[:n])
			if err != nil {
				logger.Error("Failed writing to server", slog.Any("error", err))
				return
			}
		}
	}()

	go func() {
		for {
			// Server -> Client
			n, err := serverConn.Read(rxBuffer)
			if err != nil && err != io.EOF {
				logger.Error("Failed reading from server", slog.Any("error", err))
				return
			}
			_, err = clientConn.Write(rxBuffer[:n])
			if err != nil {
				logger.Error("Failed writing to client", slog.Any("error", err))
				return
			}
		}
	}()

	// This just blocks the goroutine
	select {}
}
