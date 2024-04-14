package handler

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gbadali/equipmentMap/db"
	"github.com/gbadali/equipmentMap/view/equipment"
	"github.com/labstack/echo/v4"
)

type EquipmentHandler struct {
	DB *sql.DB
	Q  *db.Queries
}

// NewEquipmentHandler creates a new EquipmentHandler with the given db and queries
func NewEquipmentHandler(db *sql.DB, q *db.Queries) *EquipmentHandler {
	return &EquipmentHandler{
		DB: db,
		Q:  q,
	}
}

// HandleShowEquipment handles the request to show all equipment in a table
func (e EquipmentHandler) HandleShowEquipment(c echo.Context) error {
	equipmentList, err := e.Q.ListEquipmentAndParent((c.Request().Context()))
	if err != nil {
		return err
	}
	return render(c, equipment.EquipmentList(equipmentList))
}

// HandleSelectOptions handles the request to show all equipment in a select and generates the options
func (e EquipmentHandler) HandleSelectOptions(c echo.Context) error {
	equipmentList, err := e.Q.ListEquipment(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, equipment.EquipmentSelectOptions(equipmentList, ""))
}

// HandleAddEquipment handles the request to show the form to add a new equipment
func (e EquipmentHandler) HandleAddEquipment(c echo.Context) error {
	list, err := e.Q.ListEquipment(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, equipment.EquipmentForm(list))
}

// HandleSaveEquipment handles the post request to save a new equipment
func (e EquipmentHandler) HandleSaveEquipment(c echo.Context) error {
	name := c.FormValue("name")
	parent, err := strconv.ParseInt(c.FormValue("parent"), 10, 64)
	if err != nil {
		return err
	}
	equip := db.CreateEquipmentParams{
		Name:   name,
		Parent: parent,
	}

	list, err := e.Q.ListEquipment(c.Request().Context())
	if err != nil {
		return err
	}

	err = e.Q.CreateEquipment(c.Request().Context(), equip)
	if err != nil {
		return err
	}

	return render(c, equipment.EquipmentForm(list))
}

// HandleShowIndividualEquipment handles the request to show an individual equipment
func (e EquipmentHandler) HandleShowIndividualEquipment(c echo.Context) error {
	isHTMX := c.Request().Header.Get("HX-Request") == "true"
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing id: %v", err)
		slog.Error(err.Error())
		return err
	}

	breadcrumbs, err := e.breadcrumbs(c, id)
	if err != nil {
		err = fmt.Errorf("error getting breadcrumbs: %v", err)
		slog.Error(err.Error())
		return err
	}

	// Create the equipment, parent and children variables
	var equip db.Equipment
	var parent db.Equipment

	equip, err = e.Q.GetEquipment(c.Request().Context(), id)
	if err != nil {
		err = fmt.Errorf("error getting equipment from DB: %v", err)
		slog.Error(err.Error())
		return err
	}
	if equip.Parent != 0 {
		parent, err = e.Q.GetEquipment(c.Request().Context(), equip.Parent)
		if err != nil {
			err = fmt.Errorf("error getting parent equipment from DB: %v", err)
			slog.Error(err.Error())
			return err
		}

	} else {
		parent = db.Equipment{
			ID:     0,
			Name:   "None",
			Parent: 0,
		}
	}
	if isHTMX {
		return render(c, equipment.Equipment(equip, parent))
	}
	return render(c, equipment.EquipmentLayout(equip, parent, breadcrumbs))
}

// HandleEditEquipment handles the request to show the form to edit an equipment
func (e EquipmentHandler) HandleEditEquipment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	// Create the equipment, parent and children variables
	var equip db.Equipment
	var list []db.Equipment
	var parent db.Equipment

	equip, err = e.Q.GetEquipment(c.Request().Context(), id)
	if err != nil {
		return err
	}

	parent, err = e.Q.GetEquipment(c.Request().Context(), equip.Parent)
	if err != nil {
		return err
	}

	list, err = e.Q.ListEquipment(c.Request().Context())
	if err != nil {
		return err
	}

	// Reomve the equipment that is a parent of the current equipment from the list
	// this prevents cyclical references
	// This is wrong because it removes the parents from the list but instead it
	// should remove the children from the list
	newlist := preventCycles(list, id)

	return render(c, equipment.EditEquipment(equip, parent, newlist))

}

func (e EquipmentHandler) HandleUpdateEquipment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	fromData, err := c.FormParams()
	if err != nil {
		return err
	}
	name := fromData.Get("name")
	parentStr := fromData.Get("parent")
	parent, err := strconv.ParseInt(parentStr, 10, 64)
	if err != nil {
		return err
	}

	if name == "" {
		return c.Redirect(
			http.StatusUnprocessableEntity,
			fmt.Sprintf("/equipment/%v/edit", id),
		)
	}
	// Check if the parent is the same as the equipment
	if id == parent {
		parent = 0
		return fmt.Errorf("equipment can't be its own parent")
	}

	err = e.Q.UpdateEquipment(c.Request().Context(), db.UpdateEquipmentParams{
		ID:     id,
		Name:   name,
		Parent: parent,
	})
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/equipment/%v", id))
}
