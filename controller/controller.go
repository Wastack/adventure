package controller

import (
	"fmt"
	"github.com/Wastack/adventure/engine"
	"github.com/Wastack/adventure/httputil"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Controller struct {
	model engine.GameDataI
}

func NewController(m engine.GameDataI) *Controller {
	return &Controller{
		model: m,
	}
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
	if c.model.Start() == nil {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("No start point available"))
		return
	}
	ctx.JSON(http.StatusOK, gameDataToResponse(c.model.Start()))
}

// NextState godoc
// @Summary Show next state after action
// @Description Show next state after executing action
// @Accept  json
// @Param state query string true "Current state"
// @Param action query string true "Id of action to execute"
// @Produce  json
// @Success 200 {object} StateResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /adventure [get]
func (c *Controller) NextState(ctx *gin.Context) {
	state_name := ctx.Query("state")
	action_id := ctx.Query("action")

	node := c.model.GetNodeByString(state_name)
	if node == nil {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("No state: %s", state_name))
		return
	}
	new_node := node.Next(engine.ActionId(action_id))
	if new_node == nil {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("State has no action: %s", action_id))
		return
	}

	ctx.JSON(http.StatusOK, gameDataToResponse(new_node))
}

// ResolveSecret godoc
// @Summary Check if secret Ok
// @Description Decides whether a given secret resolves an action
// @Accept  json
// @Param secret query string true "Secret candidate"
// @Param action query string true "Id of action to execute"
// @Produce  json
// @Success 200 {object} SecretResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /adventure/secret [get]
func (c *Controller) ResolveSecret(ctx *gin.Context) {
	secret := ctx.Query("secret")
	action_id := ctx.Query("action")
	action_info := c.model.GetActionById(engine.ActionId(action_id))
	if action_info == nil {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("No such action id: %s", action_id))
		return
	}
	if action_info.Secret == "" {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("Action with id: %s has no secret", action_id))
		return
	}
	is_secret_ok := strings.EqualFold(action_info.Secret, secret)
	ctx.JSON(http.StatusOK, SecretResponse{SecretOk: is_secret_ok})
}
