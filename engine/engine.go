package engine

type GameAction int

type GameActionInfo struct {
}

type StoryContent string

type GameNodeI interface {
	IsGameOver() bool
	IsGameLost() bool
	Next(GameAction) GameNodeI
	Actions() map[GameAction]string
	Story() StoryContent
}

type GameDataI interface {
	Start() GameNodeI
}
