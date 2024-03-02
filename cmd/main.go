package main

import (
	"github.com/gbadali/equipmentMap/db"
	"github.com/gbadali/equipmentMap/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Connect to the database
	db.GetDBConnection()
	q := db.New(db.DB)

	equipHand := handler.NewEquipmentHandler(db.DB, q)

	e.GET("/", equipHand.HandleShowEquipment)
	e.GET("/equipmentList", equipHand.HandleListEquipment)
	e.GET("/equipment/add", equipHand.HandleAddEquipment)
	e.POST("/equipment", equipHand.HandleSaveEquipment)
	e.GET("/equipment/:id", equipHand.HandleShowIndividualEquipment)
	e.GET("/equipment/:id/edit", equipHand.HandleEditEquipment)
	// e.GET("/util/breadcrumbs/:id", equipHand.HandleBreadcrumbs)

	e.Logger.Fatal(e.Start(":3000"))
	// fmt.Println("Listening on :3000")
}
