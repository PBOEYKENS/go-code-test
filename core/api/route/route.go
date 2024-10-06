package route

import (
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/soltix-dev/go-code-test/core/api/route/collection"
	"github.com/soltix-dev/go-code-test/core/api/route/cors"
	"github.com/soltix-dev/go-code-test/core/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, ethClient *ethclient.Client, gin *gin.Engine) {
	// Create a CORS middleware with permissive settings
	gin.Use(cors.CorsMiddleware())

	publicRouter := gin.Group("")
	// Collection router
	collection.NewGetExplorePageDataRouter(env, timeout, db, publicRouter)
}
