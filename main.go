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

// func init() {
// 	chave := make([]byte, 64)
// 	if _, err := rand.Read(chave); err != nil {
// 		log.Fatal(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Carregar()
	fmt.Println(config.StringConexaoBanco)
	fmt.Printf("Listening on port:%d", config.Porta)
	db.ConnectDB()
	r := router.Gerar()
	handler := cors.AllowAll().Handler(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handler))
}
