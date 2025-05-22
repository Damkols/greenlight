package main

import(
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env string
}

type application struct {
	config config
	logger *slog.logger
}

func main{
	var cfg config

	flag.IntVar(&cfg.port, "addr", 4000, "API Port Server")
	flag.StringVar(&cfg.env, "environment", "developmemt", "development| staging | production")
	flag.Parse()
}