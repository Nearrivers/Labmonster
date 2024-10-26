package games

import (
	"context"
	"flow-poc/backend/db/repository"
)

type GameRepository struct {
	q *repository.Queries
}

func NewGameRepository(q *repository.Queries) *GameRepository {
	return &GameRepository{
		q,
	}
}

func (gr *GameRepository) AddGame(newGame repository.AddGameParams) (repository.Game, error) {
	ctx := context.Background()

	game, err := gr.q.AddGame(ctx, newGame)
	if err != nil {
		return repository.Game{}, err
	}

	return game, nil
}

func (gr *GameRepository) GetOneGame(id int64) (repository.Game, error) {
	ctx := context.Background()

	game, err := gr.q.GetOneGame(ctx, id)
	if err != nil {
		return repository.Game{}, err
	}

	return game, nil
}

func (gr *GameRepository) ListGames() ([]repository.Game, error) {
	ctx := context.Background()

	games, err := gr.q.ListGames(ctx)
	if err != nil {
		return []repository.Game{}, err
	}

	return games, err
}

func (gr *GameRepository) UpdateGame(editedGame repository.EditGameParams) error {
	ctx := context.Background()

	err := gr.q.EditGame(ctx, editedGame)
	if err != nil {
		return err
	}

	return err
}

func (gr *GameRepository) DeleteGame(id int64) error {
	ctx := context.Background()

	err := gr.q.DeleteGame(ctx, id)
	if err != nil {
		return err
	}

	return err
}
