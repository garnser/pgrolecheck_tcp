// config.go
package main

import (
	"flag"
	"github.com/go-ini/ini"
	"log"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	LogDest    string
	Foreground bool
 	ListenPort string
}

// LoadConfig reads configuration from file and overrides with command line arguments.
func LoadConfig() Config {
	var cfgFile string
	flag.StringVar(&cfgFile, "config", "/etc/pgrolecheck/pgrolecheck.conf", "Path to configuration file")

	dbUser := flag.String("dbuser", "", "Database user")
	dbPassword := flag.String("dbpassword", "", "Database password")
	dbName := flag.String("dbname", "", "Database name")
	dbHost := flag.String("dbhost", "", "Database host")
	logDest := flag.String("logdest", "", "Logging destination (STDOUT, file path, or 'syslog')")
	foreground := flag.Bool("foreground", false, "Run in the foreground instead of daemonizing")
	listenPort := flag.String("listenport", "60065", "Listening port of the service")

	flag.Parse()

	fileCfg := readIniFile(cfgFile)

	// Override INI config with any non-default command line arguments
	if *dbUser != "" {
		fileCfg.DBUser = *dbUser
	}
	if *dbPassword != "" {
		fileCfg.DBPassword = *dbPassword
	}
	if *dbName != "" {
		fileCfg.DBName = *dbName
	}
	if *dbHost != "" {
		fileCfg.DBHost = *dbHost
	}
	if *logDest != "" {
		fileCfg.LogDest = *logDest
	}
	fileCfg.Foreground = *foreground
	fileCfg.ListenPort = *listenPort

	return fileCfg
}

// readIniFile reads the INI configuration file.
func readIniFile(cfgFile string) Config {
	cfg, err := ini.Load(cfgFile)
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}

	return Config{
		DBUser:     cfg.Section("").Key("DBUser").String(),
		DBPassword: cfg.Section("").Key("DBPassword").String(),
		DBName:     cfg.Section("").Key("DBName").String(),
		DBHost:     cfg.Section("").Key("DBHost").String(),
		LogDest:    cfg.Section("").Key("LogDest").String(),
	}
}
