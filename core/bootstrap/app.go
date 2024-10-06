package bootstrap

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type Application struct {
	Env        *Env
	Gorm       *gorm.DB
	EthClient  *ethclient.Client
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv(".env")
	app.Gorm = NewGormDatabase(app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseGormDbConnection(app.Gorm)
}
