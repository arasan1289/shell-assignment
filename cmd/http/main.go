package main

import (
	"fmt"
	"os"

	"github.com/arasan1289/shell-test/internal/adapters/config"
	"github.com/arasan1289/shell-test/internal/adapters/handlers/http"
	"github.com/arasan1289/shell-test/internal/adapters/storage/repository"
	"github.com/arasan1289/shell-test/internal/core/domain"
	"github.com/arasan1289/shell-test/internal/core/service"
)

func main() {
	// Initialize config
	config, err := config.New()
	if err != nil {
		fmt.Println("Error initializing config:", err)
		os.Exit(1)
	}

	var visitors domain.Visitors

	// Initialize Handlers
	visitorRepo := repository.NewVisitorRepository(&visitors)
	visitorSvc := service.NewVisitorService(visitorRepo)
	visitorHandler := http.NewVisitorHandler(visitorSvc)

	// Initialize router
	router, err := http.NewRouter(config, *visitorHandler)
	if err != nil {
		fmt.Println("Error Initializing router")
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	err = router.Serve(listenAddr)
	if err != nil {
		fmt.Println("Error starting the HTTP server")
		os.Exit(1)
	}

}
