package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/pkg/util/connect"
	"go.elastic.co/ecszap"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
	"os"
)

type Startup struct {
	Config    *Config
	DBClient  *mongo.Client
	Router    *gin.Engine
	ZapLogger *zap.Logger
}

func NewStartup(config *Config) *Startup {
	return &Startup{
		Config:    config,
		DBClient:  initDatabase(config),
		Router:    initRouter(),
		ZapLogger: initZapLogger(config),
	}
}

func initDatabase(config *Config) *mongo.Client {
	mongoPattern := connect.MongoPattern(config.Database.Host, config.Database.Port)
	clientOptions := options.Client().ApplyURI(mongoPattern)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}

	return client
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())

	return router
}

func initZapLogger(config *Config) *zap.Logger {
	ecsEncoder := ecszap.NewDefaultEncoderConfig()
	ecsEncoderCore := ecszap.NewCore(ecsEncoder, os.Stdout, zap.DebugLevel)

	logger := zap.New(ecsEncoderCore, zap.AddCaller())

	logger = logger.With(zap.Any("app_info", config.App))
	logger = logger.With(zap.Any("server_info", config.Server))

	return logger
}
