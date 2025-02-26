// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createEquipment = `-- name: CreateEquipment :exec
INSERT INTO equipment (name, parent)
VALUES (?, ?)
RETURNING id
`

type CreateEquipmentParams struct {
	Name   string
	Parent int64
}

func (q *Queries) CreateEquipment(ctx context.Context, arg CreateEquipmentParams) error {
	_, err := q.db.ExecContext(ctx, createEquipment, arg.Name, arg.Parent)
	return err
}

const getEquipment = `-- name: GetEquipment :one
SELECT id, name, parent FROM equipment
WHERE id = ? LIMIT 1
`

func (q *Queries) GetEquipment(ctx context.Context, id int64) (Equipment, error) {
	row := q.db.QueryRowContext(ctx, getEquipment, id)
	var i Equipment
	err := row.Scan(&i.ID, &i.Name, &i.Parent)
	return i, err
}

const getHierarchy = `-- name: GetHierarchy :many
WITH RECURSIVE parents AS (
  SELECT id, name, parent
  FROM equipment AS e
  WHERE e.id = ?  -- Replace ? with the given id
  UNION ALL
  SELECT p.id, p.name, p.parent
  FROM equipment p
  INNER JOIN parents c ON p.id = c.parent
)
SELECT id, name, id, name, parent FROM parents
WHERE parent IS NULL OR parent IS NOT NULL
ORDER BY parent
`

type GetHierarchyRow struct {
	ID     int64
	Name   string
	ID_2   int64
	Name_2 string
	Parent int64
}

func (q *Queries) GetHierarchy(ctx context.Context, id int64) ([]GetHierarchyRow, error) {
	rows, err := q.db.QueryContext(ctx, getHierarchy, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetHierarchyRow
	for rows.Next() {
		var i GetHierarchyRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ID_2,
			&i.Name_2,
			&i.Parent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listChildren = `-- name: ListChildren :many
SELECT id, name, parent FROM equipment
WHERE parent = ?
`

func (q *Queries) ListChildren(ctx context.Context, parent int64) ([]Equipment, error) {
	rows, err := q.db.QueryContext(ctx, listChildren, parent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Equipment
	for rows.Next() {
		var i Equipment
		if err := rows.Scan(&i.ID, &i.Name, &i.Parent); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEquipment = `-- name: ListEquipment :many
SELECT id, name, parent FROM equipment
ORDER BY id ASC
`

func (q *Queries) ListEquipment(ctx context.Context) ([]Equipment, error) {
	rows, err := q.db.QueryContext(ctx, listEquipment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Equipment
	for rows.Next() {
		var i Equipment
		if err := rows.Scan(&i.ID, &i.Name, &i.Parent); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEquipmentAndParent = `-- name: ListEquipmentAndParent :many
SELECT e.id, e.name AS equipment_name, p.name as parent_name
FROM equipment e 
LEFT JOIN equipment p ON e.parent = p.id
ORDER BY e.id ASC
`

type ListEquipmentAndParentRow struct {
	ID            int64
	EquipmentName string
	ParentName    sql.NullString
}

func (q *Queries) ListEquipmentAndParent(ctx context.Context) ([]ListEquipmentAndParentRow, error) {
	rows, err := q.db.QueryContext(ctx, listEquipmentAndParent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEquipmentAndParentRow
	for rows.Next() {
		var i ListEquipmentAndParentRow
		if err := rows.Scan(&i.ID, &i.EquipmentName, &i.ParentName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEquipment = `-- name: UpdateEquipment :exec
UPDATE equipment
SET name = ?, parent = ?
WHERE id = ?
`

type UpdateEquipmentParams struct {
	Name   string
	Parent int64
	ID     int64
}

func (q *Queries) UpdateEquipment(ctx context.Context, arg UpdateEquipmentParams) error {
	_, err := q.db.ExecContext(ctx, updateEquipment, arg.Name, arg.Parent, arg.ID)
	return err
}
