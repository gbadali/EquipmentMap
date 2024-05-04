package handler

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gbadali/equipmentMap/db"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

// this function returns the breadcrumbs for a given equipment id
func (e EquipmentHandler) breadcrumbs(c echo.Context, id int64) ([]db.Equipment, error) {
	var bread []db.Equipment
	originalID := id
	equip, err := e.Q.GetEquipment(c.Request().Context(), id)
	fmt.Println(equip, equip.Parent)
	fmt.Println(errors.Is(err, sql.ErrNoRows))
	if err != nil {
		err = fmt.Errorf("error getting equipment from DB for breadcrumbs: %v", err)
		return nil, err
	}

	for equip.Parent > 0 {
		id = equip.Parent
		bread = append([]db.Equipment{equip}, bread...)
		if id == originalID {
			err = fmt.Errorf("error getting breadcrumbs: equipment has a circular reference")
			return nil, err
		}
	}
	return bread, nil
}

func (e EquipmentHandler) getChildren(c echo.Context, id int64) ([]db.Equipment, error) {
	list, err := e.Q.ListChildren(c.Request().Context(), id)
	if err != nil {
		err = fmt.Errorf("error getting children from DB: %v", err)
		return nil, err
	}
	return list, nil
}

// this function prevents cycles when the user edits the equipment
func preventCycles(list []db.Equipment, id int64) []db.Equipment {
	return list
}
