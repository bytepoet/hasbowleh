package main

import (
	"log"
	"net/http"

	"github.com/bytepoet/hasbowleh/internal/api"
	"github.com/bytepoet/hasbowleh/internal/database"
	"github.com/bytepoet/hasbowleh/internal/v2ray"
	"github.com/bytepoet/hasbowleh/internal/wireguard"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	wg := wireguard.New()
	v2 := v2ray.New()

	router := api.SetupRouter(db, wg, v2)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
