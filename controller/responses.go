package controller

import (
	"github.com/Wastack/adventure/engine"
	"log"
)

type ActionResponse struct {
	State_id string `json:"state_id" example:"state_xy"`
	Text     string `json:"text" example:"Move to state xy where a dragon awaits you"`
}

type StateResponse struct {
	NewState string           `json:"new_state" example:"strtState"`
	IsOver   bool             `json:"is_over" example:"false"`
	IsDeath  bool             `json:"is_death" example:"false"`
	Actions  []ActionResponse `json:"Actions"`
}

func gameDataToResponse(node engine.GameNodeI) StateResponse {
	if node == nil {
		log.Printf("Missing game data")
		return StateResponse{}
	}
	actions := make([]ActionResponse, len(node.Actions()))
	i := 0
	for k, v := range node.Actions() {
		actions[i] = ActionResponse{State_id: k, Text: v.Story}
		i += 1
	}
	log.Printf("DEBUG: state in gameDataToResponse: %s", node)
	return StateResponse{
		NewState: node.Name(),
		IsOver:   node.IsGameOver(),
		IsDeath:  node.IsGameLost(),
		Actions:  actions,
	}
}
