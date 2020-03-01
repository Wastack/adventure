package yaml

import (
	"github.com/Wastack/adventure/engine"
	"github.com/Wastack/adventure/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func findActionByTarget(target string, node engine.GameNodeI) *engine.GameActionInfo {
	for _, v := range node.Actions() {
		if v.Target == target {
			return &v
		}
	}
	return nil
}

func TestYamlEngine(t *testing.T) {
	assert := assert.New(t)
	data, err := Parse_yaml(utils.GetYamlFromFile("testdata/example.yml"), true)
	assert.Nil(err)
	assert.NotNil(data.Start())

	// start node
	assert.Equal("kezdo_pont", data.Start().Name())
	assert.False(data.Start().IsGameLost())
	assert.False(data.Start().IsGameOver())

	mpont_action := findActionByTarget("masodik_pont", data.Start())
	hpont_action := findActionByTarget("harmadik_pont", data.Start())
	assert.NotNil(mpont_action)
	assert.NotNil(hpont_action)
	assert.Equal("masodik_pont", mpont_action.Target)
	assert.Equal("harmadik_pont", hpont_action.Target)
	assert.Equal("Lépés a második pontra", findActionByTarget("masodik_pont", data.Start()).ActionName)

	// masodik_pont
	var node engine.GameNodeI
	for k, v := range data.Start().Actions() {
		if v.Target == "masodik_pont" {
			node = data.Start().Next(k)
			break
		}
	}
	assert.Equal("masodik_pont", node.Name())
	assert.False(data.Start().IsGameLost())
	assert.False(data.Start().IsGameOver())
	kpont_action := findActionByTarget("kezdo_pont", node)
	assert.NotNil(kpont_action)
	assert.Equal("kezdo_pont", kpont_action.Target)
	assert.Equal("Visszalépés az első pontra", findActionByTarget("kezdo_pont", node).ActionName)

	// go back to first node
	for k, v := range node.Actions() {
		if v.Target == "kezdo_pont" {
			node = node.Next(k)
			break
		}
	}
	assert.Equal(node.Name(), "kezdo_pont")

	// assert that graph is connected
	game_data := data.(*GameData)
	assert.Empty(check_story_connected(game_data))
}
