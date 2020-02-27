package main

import (
	"./docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/Wastack/adventure/engine"
	"github.com/Wastack/adventure/engine/data/yaml"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io/ioutil"
	"log"
	"os"
)

func parse(file_path string) (engine.GameDataI, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return yaml.Parse_yaml(b)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Missing yaml file path as first argument\n")
	}
	file_path := os.Args[1]
	game_data, err := parse(file_path)
	if err != nil {
		log.Fatalf("Error while parsing yaml file: %v", err)
	}

	// TODO
	_ = game_data
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Adventure API"
	docs.SwaggerInfo.Description = "This is a server for the adventure game engine."
	docs.SwaggerInfo.Version = "0.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()

}
