package main

/*
	usage:
	  - go run main.go --client.foo=baz
	  - TIKTOK_CLIENT_FOO=baz TIKTOK_CLIENT_ECHO=0 go run main.go
	  - go run main.go --config <path to config>
*/

import (
	"fmt"

	"time"

	"github.com/a76yyyy/tiktok/pkg/dlog"
	"github.com/a76yyyy/tiktok/pkg/ttviper"
	_ "github.com/spf13/viper/remote" // enabble viper remote config
)

func main() {

	config := ttviper.ConfigInit("TIKTOK", "userConfig")
	viper := config.Viper

	logger := dlog.InitLog()
	defer logger.Sync()
	logger.Info("logger construction succeeded")

	sugar := logger.Sugar()
	defer sugar.Desugar()
	sugar.Info("sugar consturction succeeded")

	sugar.Infow("Conf", "Global.Source", viper.GetString("global.source"))
	// sugar.Errorf("error")

	// where does it from
	fmt.Printf("Global.Source: '%s'\n", viper.GetString("global.source"))
	fmt.Printf("Global.ChangeMe: '%s'\n", viper.GetString("Global.ChangeMe"))
	// prints 'default(viper)'
	fmt.Printf("viper.GetString(\"Global.Unset\") = '%s'\n", viper.GetString("global.unset"))
	fmt.Printf("Var GlobalUnset = '%s'\n", *ttviper.GlobalUnset)

	// from config file
	fmt.Println("client.servers: ", viper.GetStringSlice("client.servers"))
	fmt.Println("Server.Address: ", viper.GetString("Server.Address"))
	// it can be changed... but when to do that?
	viper.Set("Server.Address", "0.0.0.0")
	// case *insensitive*
	fmt.Println("Server.Address: ", viper.GetString("server.address"))

	// from env
	fmt.Println("client.foo:", viper.GetString("client.foo"))
	fmt.Println("client.echo:", viper.GetBool("client.echo"))

	// block for watch test
	time.Sleep(3600 * time.Second)
}
