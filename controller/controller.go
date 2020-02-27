package controller

import (
	"github.com/Wastack/adventure/engine"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	model engine.GameDataI
}

func NewController(engine.GameDataI) *Controller {
	return &Controller{}
}

// ShowFirstState godoc
// @Summary Show first state
// @Description Show first state of the adventure game
// @Accept  json
// @Produce  json
// @Success 200 {object} StateResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /adventure/first [get]
func (c *Controller) ShowFirstState(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StateResponse{
		Actions: []string{"Blah", "Foo", "Bar"},
	})
}

// NextState godoc
// @Summary Show next state after action
// @Description Show next state after executing action
// @Accept  json
// @Param state query string true "Current state"
// @Param action query string true "Action to execute"
// @Produce  json
// @Success 200 {object} StateResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /adventure [get]
func (c *Controller) NextState(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StateResponse{
		Actions: []string{"Blah", "Foo", "Bar"},
	})
}
