package bootstrap

import (
	"log"

	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
)

type EnvInput struct {
	IsLocal             bool   `mapstructure:"IS_LOCAL"`
	IsTestnet           bool   `mapstructure:"IS_TESTNET"`
	IsDocker            bool   `mapstructure:"IS_DOCKER"`
	IsProduction        bool   `mapstructure:"IS_PRODUCTION"`
	LocalServerAddress  string `mapstructure:"LOCAL_SERVER_ADDRESS"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	LocalDomainAddress  string `mapstructure:"LOCAL_DOMAIN_ADDRESS"`
	DockerDomainAddress string `mapstructure:"DOCKER_DOMAIN_ADDRESS"`
	DomainAddress       string `mapstructure:"DOMAIN_ADDRESS"`
	ServerPort          string `mapstructure:"SERVER_PORT"`
	ContextTimeout      int    `mapstructure:"CONTEXT_TIMEOUT"`
	LocalDBHost         string `mapstructure:"LOCAL_DB_HOST"`
	TestnetDBName       string `mapstructure:"TESTNET_DB_NAME"`
	TestnetDBAdminUUID  string `mapstructure:"TESTNET_DB_ADMIN_UUID"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPass              string `mapstructure:"DB_PASS"`
	LocalDBName         string `mapstructure:"LOCAL_DB_NAME"`
	DBName              string `mapstructure:"DB_NAME"`
	DBAdminUUID         string `mapstructure:"DB_ADMIN_UUID"`
}

type Env struct {
	IsLocal        bool
	IsTestnet      bool
	IsDocker       bool
	IsProduction   bool
	ServerAddress  string
	DomainAddress  string
	ServerPort     string
	ContextTimeout int
	DBHost         string
	DBPort         string
	DBUser         string
	DBPass         string
	DBName         string
	DBAdminUUID    uuid.UUID
}

// Handles any conversion of types between envInputs and the Env
func fromEnvInput(envInput *EnvInput) *Env {
	isProduction := envInput.IsProduction
	isLocal := envInput.IsLocal
	isDocker := envInput.IsDocker
	isTestnet := envInput.IsTestnet
	// Assuming production is true
	// domain
	serverAddr := envInput.ServerAddress
	domainAddr := envInput.DomainAddress
	// db
	dbHost := envInput.DBHost
	dbName := envInput.DBName
	dbAdminUuid, err := uuid.FromString(envInput.DBAdminUUID)
	// This can only fail if we are connecting to testnet
	uuidFailed := false
	if err != nil {
		uuidFailed = true
		log.Print("Warning: Mainnet Invalid UUID")
	}

	// If production
	if isProduction {
		isLocal = false
		isTestnet = false
		isDocker = false
		if uuidFailed {
			log.Fatal("Invalid UUID")
		}
		// If production is not true
	} else {
		// Can't we local and docker
		if isLocal {
			serverAddr = envInput.LocalServerAddress
			domainAddr = envInput.LocalDomainAddress
			dbHost = envInput.LocalDBHost
			dbName = envInput.LocalDBName
		} else if isDocker {
			// ServerAddr and dbHost is the same as production which is already set
			domainAddr = envInput.DockerDomainAddress
		}
		// UUID
		if isTestnet {
			dbName = envInput.TestnetDBName
			dbAdminUuid, err = uuid.FromString(envInput.TestnetDBAdminUUID)
			if err != nil {
				log.Fatal("Invalid UUID")
			}
		} else {
			// Ensuring if UUID for mainnet db failed we catch it
			if uuidFailed {
				log.Fatal("Invalid UUID")
			}
		}
	}

	return &Env{
		IsLocal:        isLocal,
		IsTestnet:      isTestnet,
		IsProduction:   isProduction,
		ServerAddress:  serverAddr,
		DomainAddress:  domainAddr,
		ContextTimeout: envInput.ContextTimeout,
		DBHost:         dbHost,
		DBPort:         envInput.DBPort,
		DBUser:         envInput.DBUser,
		DBPass:         envInput.DBPass,
		DBName:         dbName,
		DBAdminUUID:    dbAdminUuid,
	}
}

func NewEnv(path string) *Env {
	envInput := EnvInput{}
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&envInput)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if envInput.IsProduction {
		log.Println("The App is running in production env")
	} else {
		if envInput.IsLocal {
			log.Println("The App is running in local env")
		} else {
			log.Println("The App is running in non-locally env")
		}
		if envInput.IsTestnet {
			log.Println("The App is running in testnet mode")
		} else {
			log.Println("The App is running in mainnet mode")
		}
	}

	env := fromEnvInput(&envInput)

	return env
}
