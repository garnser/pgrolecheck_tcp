// server.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
)

// StartServer initializes the TCP server and handles incoming connections.
func StartServer(cfg Config, logger *log.Logger) {
	listener, err := net.Listen("tcp", ":60065")
	if err != nil {
		logger.Fatalf("Failed to listen on port 60065: %v", err)
	}
	defer listener.Close()
 	logger.Printf("Server is listening on port %s", cfg.ListenPort)

	connStr := fmt.Sprintf("postgres://%s:%s@/%s?host=%s&sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatalf("Failed to open a DB connection: %v", err)
	}
	defer db.Close()

	logger.Println("Server is listening on port 60065")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn, db, logger)
	}
}

// handleConnection processes an individual client connection.
func handleConnection(conn net.Conn, db *sql.DB, logger *log.Logger) {
	defer conn.Close()

	var isInRecovery bool
	err := db.QueryRow("SELECT pg_is_in_recovery();").Scan(&isInRecovery)
	if err != nil {
		logger.Printf("Failed to query database: %v", err)
		return
	}

	if isInRecovery {
		logger.Println("Postgres is in recovery mode, terminating connection.")
		return
	}

	_, err = conn.Write([]byte("Connected\n"))
	if err != nil {
    logger.Printf("Failed to send response: %v", err)
		return
	}

  logger.Println("Response sent to client.")
}
