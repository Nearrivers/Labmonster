-- name: ListGames :many
SELECT * from games
ORDER BY name;

-- name: GetOneGame :one
SELECT * FROM games
WHERE id = ? LIMIT 1;

-- name: AddGame :one
INSERT INTO games(
  name, iconPath
) VALUES (
  ?, ?
)
RETURNING *;

-- name: EditGame :exec
UPDATE games SET name = ?, iconPath = ?
WHERE id = ?;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = ?;