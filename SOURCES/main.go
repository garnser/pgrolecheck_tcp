// main.go
package main

func main() {
	config := LoadConfig()

	logger := SetupLogging(config.LogDest, config.Foreground)

	StartServer(config, logger)
}
