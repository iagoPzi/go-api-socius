package main

import (
	"api/src/config"
	"api/src/db"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	config.Carregar()
	fmt.Printf("Listening on port:%d", config.Porta)
	db.ConnectDB()
	r := router.Gerar()
	handler := cors.AllowAll().Handler(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handler))
}
