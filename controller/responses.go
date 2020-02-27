package controller

type StateResponse struct {
	NewState string   `json:"new_state" example:"strtState"`
	IsOver   bool     `json:"is_over" example:"false"`
	IsDeath  bool     `json:"is_death" example:"false"`
	Actions  []string `json:"new_actions" example:"Foo,Bar,Blah"`
}
