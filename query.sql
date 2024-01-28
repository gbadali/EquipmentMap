-- name: GetEquipment :one
SELECT * FROM equipment
WHERE id = ? LIMIT 1;

-- name: ListEquipmentAndParent :many
SELECT e.id, e.name AS equipment_name, p.name as parent_name
FROM equipment e 
LEFT JOIN equipment p ON e.parent = p.id
ORDER BY e.id ASC;

-- name: ListEquipment :many
SELECT * FROM equipment
ORDER BY id ASC;

-- name: CreateEquipment :exec
INSERT INTO equipment (name, parent)
VALUES (?, ?)
RETURNING id;

-- name: ListChildren :many
SELECT * FROM equipment
WHERE parent = ?;