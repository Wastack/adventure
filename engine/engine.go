package engine

type GameAction int

type GameActionInfo struct {
	ActionName string
	Story      string
	Target     string
	Secret     string
}

type StoryContent string
type ActionId string

type GameNodeI interface {
	IsGameOver() bool
	IsGameLost() bool
	Next(ActionId) GameNodeI // action ID -> Node
	Actions() map[ActionId]GameActionInfo
	Story() StoryContent
	Name() string
}

type GameDataI interface {
	Start() GameNodeI
	GetNodeByString(string) GameNodeI
	GetActionById(ActionId) *GameActionInfo
}
