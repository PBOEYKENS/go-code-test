package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soltix-dev/go-code-test/core/api/route"
	"github.com/soltix-dev/go-code-test/core/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, app.Gorm, app.EthClient, gin)

	gin.Run(env.ServerAddress)
}
