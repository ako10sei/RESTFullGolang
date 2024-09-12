package main

import (
	"RESTFullGolang/internal/config"
	"fmt"
)

// config, logger, storage, router, server
func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

}
