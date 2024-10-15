-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS games (
  id INTEGER PRIMARY KEY,
  name text NOT NULL,
  iconPath text
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE games
-- +goose StatementEnd
