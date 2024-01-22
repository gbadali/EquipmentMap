-- name: GetEquipment :one
SELECT * FROM equipment
WHERE id = ? LIMIT 1;

-- name: ListEquipmentAndParent :many
SELECT e.id, e.name AS equipment_name, p.name as parent_name
FROM equipment e 
JOIN equipment p ON e.parent = p.id
ORDER BY e.name ASC;

-- name: ListEquipment :many
SELECT * FROM equipment
ORDER BY name ASC;

-- name: CreateEquipment :exec
INSERT INTO equipment (name, parent)
VALUES (?, ?)
RETURNING id;

