// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: game_queries.sql

package repository

import (
	"context"
)

const addGame = `-- name: AddGame :one
INSERT INTO games(
  name, iconPath
) VALUES (
  ?, ?
)
RETURNING id, name, iconpath
`

type AddGameParams struct {
	Name     string `json:"name"`
	Iconpath string `json:"iconpath"`
}

func (q *Queries) AddGame(ctx context.Context, arg AddGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, addGame, arg.Name, arg.Iconpath)
	var i Game
	err := row.Scan(&i.ID, &i.Name, &i.Iconpath)
	return i, err
}

const deleteGame = `-- name: DeleteGame :exec
DELETE FROM games
WHERE id = ?
`

func (q *Queries) DeleteGame(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteGame, id)
	return err
}

const editGame = `-- name: EditGame :exec
UPDATE games SET name = ?, iconPath = ?
WHERE id = ?
`

type EditGameParams struct {
	Name     string `json:"name"`
	Iconpath string `json:"iconpath"`
	ID       int64  `json:"id"`
}

func (q *Queries) EditGame(ctx context.Context, arg EditGameParams) error {
	_, err := q.db.ExecContext(ctx, editGame, arg.Name, arg.Iconpath, arg.ID)
	return err
}

const getOneGame = `-- name: GetOneGame :one
SELECT id, name, iconpath FROM games
WHERE id = ? LIMIT 1
`

func (q *Queries) GetOneGame(ctx context.Context, id int64) (Game, error) {
	row := q.db.QueryRowContext(ctx, getOneGame, id)
	var i Game
	err := row.Scan(&i.ID, &i.Name, &i.Iconpath)
	return i, err
}

const listGames = `-- name: ListGames :many
SELECT id, name, iconpath from games
ORDER BY name
`

func (q *Queries) ListGames(ctx context.Context) ([]Game, error) {
	rows, err := q.db.QueryContext(ctx, listGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(&i.ID, &i.Name, &i.Iconpath); err != nil {
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
