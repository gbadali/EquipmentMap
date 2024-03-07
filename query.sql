-- name: GetEquipment :one
SELECT id, name, COALESCE(parent, '') as parent FROM equipment
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

-- name: UpdateEquipment :exec
UPDATE equipment
SET name = ?, parent = ?
WHERE id = ?;

-- name: GetHierarchy :many
WITH RECURSIVE parents AS (
  SELECT id, name, parent
  FROM equipment AS e
  WHERE e.id = ?  -- Replace ? with the given id
  UNION ALL
  SELECT p.id, p.name, p.parent
  FROM equipment p
  INNER JOIN parents c ON p.id = c.parent
)
SELECT id, name, COALESCE(parent, 0) FROM parents
WHERE parent IS NULL OR parent IS NOT NULL
ORDER BY parent;