package main

import (
	"fmt"
	"net/http"

	"github.com/owaisnadeemdev/roadguard/internal/api/http/server"
	"github.com/owaisnadeemdev/roadguard/internal/config"
)

func main() {
	config.InitializeDB()
	router := server.NewRouter()
	fmt.Println("Server running on PORT", 5000)
	http.ListenAndServe(":5000", router)
}
