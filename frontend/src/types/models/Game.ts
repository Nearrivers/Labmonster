import { repository } from "$/models";

export type Game = Omit<repository.Game, "createFrom">
export type AddGame = Omit<repository.AddGameParams, "createFrom">