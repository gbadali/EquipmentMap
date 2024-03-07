package handler

import (
	"github.com/a-h/templ"
	"github.com/gbadali/equipmentMap/db"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func (e EquipmentHandler) breadcrumbs(c echo.Context, id int64) ([]db.GetHierarchyRow, error) {
	var bread []db.Equipment
	for {
		equip, err := e.Q.GetEquipment(c.Request().Context(), id)
		if err != nil {
			return nil, err
		}
		if equip.Parent.Valid == false {
			break
		}
		id = equip.Parent.Int64
		bread = append(bread, equip)

	}

}
