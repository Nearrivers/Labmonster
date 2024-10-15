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