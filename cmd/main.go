package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gbadali/equipmentMap/db"
	"github.com/gbadali/equipmentMap/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	turso_url := os.Getenv("TURSO_DATABASE_URL")
	turso_auth := os.Getenv("TURSO_AUTH_TOKEN")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	e := echo.New()

	// Connect to the database
	log.Println("Connecting to the database")
	url := fmt.Sprintf(`%v?authToken=%v`, turso_url, turso_auth)

	turso, err := sql.Open("libsql", url)
	if err != nil {
		log.Printf("failed to opening database: %s: %s", url, err)
		os.Exit(1)
	}
	defer turso.Close()

	q := db.New(turso)

	equipHand := handler.NewEquipmentHandler(db.DB, q)

	e.GET("/", equipHand.HandleShowEquipment)
	e.GET("/equipmentList", equipHand.HandleSelectOptions)
	e.GET("/equipment/add", equipHand.HandleAddEquipment)
	e.POST("/equipment", equipHand.HandleSaveEquipment)
	e.GET("/equipment/:id", equipHand.HandleShowIndividualEquipment)
	e.GET("/equipment/:id/edit", equipHand.HandleEditEquipment)
	e.PUT("/equipment/:id", equipHand.HandleUpdateEquipment)

	e.Logger.Fatal(e.Start(":" + port))
	// fmt.Println("Listening on :3000")
}
