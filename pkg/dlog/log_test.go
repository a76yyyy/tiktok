package dlog

import (
	"testing"
	"time"
)

func TestInitLogger(t *testing.T) {
	viper := config.Viper

	logger := InitLog(2)
	defer logger.Sync()
	logger.Info("logger construction succeeded")

	// where does it from
	logger.Infof("Global.Source: '%s'", viper.GetString("global.source"))
	logger.Infof("Global.ChangeMe: '%s'", viper.GetString("Global.ChangeMe"))
	// prints 'default(viper)'
	logger.Infof("viper.GetString(\"Global.Unset\") = '%s'", viper.GetString("global.unset"))

	// from config file
	logger.Info("client.servers: ", viper.GetStringSlice("client.servers"))
	logger.Info("Server.Address: ", viper.GetString("Server.Address"))
	// it can be changed... but when to do that?
	viper.Set("Server.Address", "0.0.0.0")
	// case *insensitive*
	logger.Info("Server.Address: ", viper.GetString("server.address"))

	// from env
	logger.Info("client.foo:", viper.GetString("client.foo"))
	logger.Info("client.echo:", viper.GetBool("client.echo"))

	// block for watch test
	time.Sleep(1 * time.Second)
}
