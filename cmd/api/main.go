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