package engine

type GameAction int

type GameActionInfo struct {
	ActionName string
	Story      string
	ActionId   string
}

type StoryContent string

type GameNodeI interface {
	IsGameOver() bool
	IsGameLost() bool
	Next(string) GameNodeI // action ID -> Node
	Actions() []GameActionInfo
	Story() StoryContent
	Name() string
}

type GameDataI interface {
	Start() GameNodeI
	GetNodeByString(string) GameNodeI
}
