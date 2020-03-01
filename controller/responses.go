package controller

import (
	"github.com/Wastack/adventure/engine"
	"log"
)

type ActionResponse struct {
	Id             string `json:"id" example:"324"`
	Target         string `json:"target_state_id,omitempty" example:"state_xy"`
	Story          string `json:"story,omitempty" example:"The dragon ate you for good"`
	ActionName     string `json:"name" example:"Move to state xy where a dragon awaits you"`
	SecretRequired bool   `json:"secret_required" example:"true"`
}

type StateResponse struct {
	NewState string           `json:"new_state" example:"strtState"`
	IsOver   bool             `json:"is_over,omitempty" example:"false"`
	IsDeath  bool             `json:"is_death,omitempty" example:"false"`
	Actions  []ActionResponse `json:"Actions"`
}

func gameDataToResponse(node engine.GameNodeI) StateResponse {
	if node == nil {
		log.Printf("Missing game data")
		return StateResponse{}
	}
	actions := make([]ActionResponse, len(node.Actions()))
	i := 0
	for id, aPtr := range node.Actions() {
		actions[i] = ActionResponse{Id: string(id), Target: aPtr.Target, Story: aPtr.Story, ActionName: aPtr.ActionName, SecretRequired: aPtr.Secret != ""}
		i += 1
	}
	return StateResponse{
		NewState: node.Name(),
		IsOver:   node.IsGameOver(),
		IsDeath:  node.IsGameLost(),
		Actions:  actions,
	}
}
