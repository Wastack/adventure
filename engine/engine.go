package engine

type GameAction int

type GameActionInfo struct {
	Story string
}

type StoryContent string

type GameNodeI interface {
	IsGameOver() bool
	IsGameLost() bool
	Next(string) GameNodeI              // action ID -> Node
	Actions() map[string]GameActionInfo // action ID -> actioninfo
	Story() StoryContent
	Name() string
}

type GameDataI interface {
	Start() GameNodeI
	GetNodeByString(string) GameNodeI
}
