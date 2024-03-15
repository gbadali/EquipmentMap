package handler

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gbadali/equipmentMap/db"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func (e EquipmentHandler) breadcrumbs(c echo.Context, id int64) ([]db.GetEquipmentRow, error) {
	var bread []db.GetEquipmentRow
	originalID := id
	for {
		equip, err := e.Q.GetEquipment(c.Request().Context(), id)
		if err != nil {
			err = fmt.Errorf("error getting equipment from DB for breadcrumbs: %v", err)
			return nil, err
		}
		if equip.Parent == 0 {
			break
		}
		id = equip.Parent
		bread = append(bread, equip)
		if id == originalID {
			err = fmt.Errorf("error getting breadcrumbs: equipment has a circular reference")
			return nil, err
		}
	}
	return bread, nil
}
