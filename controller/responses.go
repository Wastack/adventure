package controller

import "github.com/Wastack/adventure/engine"

type StateResponse struct {
	NewState string   `json:"new_state" example:"strtState"`
	IsOver   bool     `json:"is_over" example:"false"`
	IsDeath  bool     `json:"is_death" example:"false"`
	Actions  []string `json:"new_actions" example:"Foo,Bar,Blah"`
}

func gameDataToResponse(node engine.GameNodeI) StateResponse {
	actions := make([]string, len(node.Actions()))
	for i, a := range node.Actions() {
		actions[i] = a
	}
	return StateResponse{
		NewState: node.Name(),
		IsOver:   node.IsGameOver(),
		IsDeath:  node.IsGameLost(),
		Actions:  actions,
	}
}
