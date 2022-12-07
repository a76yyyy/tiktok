// Copyright 2022 a76yyyy && CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttviper

/*
	usage:
	  - go run config_test.go --client.foo=baz
	  - TIKTOK_CLIENT_FOO=baz TIKTOK_CLIENT_ECHO=0 go run config_test.go
	  - go run config_test.go --config <path to config>
*/

import (
	"testing"
	"time"
)

func TestConfigInit(t *testing.T) {
	ConfigInit("TIKTOK", "userConfig")
}

func TestInitLogger(t *testing.T) {
	config := ConfigInit("TIKTOK", "logConfig")
	viper := config.Viper

	logger := config.InitLogger()
	defer logger.Sync()
	logger.Info("logger construction succeeded")

	// where does it from
	logger.Infof("Global.Source: '%s'", viper.GetString("global.source"))
	logger.Infof("Global.ChangeMe: '%s'", viper.GetString("Global.ChangeMe"))
	// prints 'default(viper)'
	logger.Infof("viper.GetString(\"Global.Unset\") = '%s'", viper.GetString("global.unset"))
	logger.Infof("Var GlobalUnset = '%s'", *GlobalUnset)

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
