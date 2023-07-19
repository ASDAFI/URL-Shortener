package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"url-shortener/configs"
	"url-shortener/src/infrastructure/db"
)

var configFile string

func initConfig() {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error loading config file: %s \n", err)
	}
	if err := viper.Unmarshal(&configs.Config); err != nil {
		log.Fatalf("Fatal error marshalng config file: %s \n", err)

	}

	log.Info("Configuration Loaded!")

}

func setupConfig() {
	/*redisProvider, err := cache.CreateRedisCacheProvider(configs.Config.Cache)
	if err != nil {
		log.Fatal("Error while creating redis cache provider: ", err)
	}
	cache.RedisCacheProvider = redisProvider*/

	postgresProvider, err := db.CreateDBProvider(configs.Config.Database)
	if err != nil {
		log.Fatal("an Error has been occured while creating postgres db provider: ", err)
	}
	db.PostgresDBProvider = postgresProvider

}
