package ttviper

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Viper *viper.Viper
}

var (
	configVar      string
	isRemoteConfig bool

	GlobalSource = pflag.String("global.source", "default(flag)", "identify the source of configuration")
	//var globalUnset = pflag.String("global.unset", "default(flag)", "this parameter do not appear in config file")
	GlobalUnset = pflag.String("global.unset", "", "this parameter do not appear in config file")
)

/*
	configVar 采用了另一种方式来初始化
	主要是为了强调这个命令行参数的特殊性，这个参数是需要在代码中直接引用的
	其他参数是跟viper绑定的，不会直接使用，而是通过viper.GetType()来获取

	另外一个原因是，plfag.String()返回的是*string，用起来没那么直观
*/

func init() {
	pflag.StringVar(&configVar, "config", "", "Config file path")
	pflag.BoolVar(&isRemoteConfig, "isRemoteConfig", false, "Whether to choose remote config")
}

func (v *Config) SetRemoteConfig(u *url.URL) {
	/*
		这里接受etcd 或 consul 的url

		etcd:
		  url格式为： etcd+http://127.0.0.1:2380/path/to/key.yaml
		  其中：provider=etcd, endpoint=http://127.0,0.1:2380, path=/path/to/key

		consul:
		  url格式为：consul://127.0.0.1:8500/key.json
		  其中：provider=consul, endpoint=127.0,0.1:8500, path=key.json

		TODO: consul 的 key name 可以包含 . 吗？
	*/

	var provider string
	var endpoint string
	var path string

	schemes := strings.SplitN(u.Scheme, "+", 2)
	if len(schemes) < 1 {
		panic(fmt.Errorf("invalid config scheme '%s'", u.Scheme))
	}

	provider = schemes[0]
	switch provider {

	case "etcd":
		if len(schemes) < 2 {
			panic(fmt.Errorf("invalid config scheme '%s'", u.Scheme))
		}
		protocol := schemes[1]
		endpoint = fmt.Sprintf("%s://%s", protocol, u.Host)
		path = u.Path
	case "consul":
		endpoint = u.Host
		path = u.Path[1:]
	default:
		panic(fmt.Errorf("unsupported provider '%s'", provider))
	}

	//  配置文件的后缀
	ext := filepath.Ext(path)
	if ext == "" {
		panic(fmt.Errorf("using remote config, without specifiing file extension"))
	}
	// .yaml ==> yaml
	configType := ext[1:]

	fmt.Printf("Using Remote Config Provider: '%s', Endpoint: '%s', Path: '%s', ConfigType: '%s'\n", provider, endpoint, path, configType)
	if err := v.Viper.AddRemoteProvider(provider, endpoint, path); err != nil {
		panic(fmt.Errorf("error adding remote provider %s", err))
	}

	v.Viper.SetConfigType(configType)

}

func (v *Config) SetDefaultValue() {

	v.Viper.SetDefault("global.unset", "default(viper)")
	/* add more */
}

func (v *Config) WatchRemoteConf() {
	for {
		time.Sleep(time.Second * 5) // delay after each request

		// currently, only tested with etcd support
		err := v.Viper.WatchRemoteConfig()
		if err != nil {
			fmt.Printf("unable to read remote config: %v\n", err)
			continue
		}

		// unmarshal new config into our runtime config struct. you can also use channel
		// to implement a signal to notify the system of the changes
		//runtime_viper.Unmarshal(&runtime_conf)
		fmt.Println("Watching Remote Config")
		fmt.Printf("Global.Source: '%s'\n", v.Viper.GetString("Global.Source"))
		fmt.Printf("Global.ChangeMe: '%s'\n", v.Viper.GetString("Global.ChangeMe"))
	}
}

func (v *Config) ZapLogConfig() []byte {
	log := v.Viper.Sub("Log")
	logConfig, err := json.Marshal(log.AllSettings())
	if err != nil {
		panic(fmt.Errorf("error marshalling log config %s", err))
	}
	return logConfig
}

func (v *Config) InitLogger() *zap.Logger {
	var cfg zap.Config
	if err := json.Unmarshal(v.ZapLogConfig(), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.EncoderConfig.EncodeName = zapcore.FullNameEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func ConfigInit(envPrefix string, cfgName string) Config {
	pflag.Parse()

	v := viper.New()
	config := Config{Viper: v}
	viper := config.Viper
	/*
		viper.BindPFlags 自动绑定了所有命令行参数，如果只需要其中一部分，可以用viper.BingPflag选择性绑定，如
		viper.BindPFlag("global.source", pflag.Lookup("global.source"))
	*/
	viper.BindPFlags(pflag.CommandLine)
	config.SetDefaultValue()

	// read from env
	viper.AutomaticEnv()
	// so that client.foo maps to MYAPP_CLIENT_FOO
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configVar != "" {
		/*
			如果设置了--config参数，尝试从这里解析
			它可能是一个Remote Config，来自etcd或consul
			也可能是一个本地文件
		*/
		u, err := url.Parse(configVar)
		if err != nil {
			panic(fmt.Errorf("error parsing: '%s'", configVar))
		}

		if u.Scheme != "" {
			// 看起来是个remote config
			config.SetRemoteConfig(u)
			isRemoteConfig = true
		} else {
			viper.SetConfigFile(configVar)
		}
	} else {
		/*
			尝试搜索若干默认路径，先后顺序如下:
			- /etc/tiktok/config/userConfig.<ext>
			- ~/.tiktok/userConfig.<ext>
			- ./userConfig.<ext>

			其中<ext> 是 viper所支持的文件类型，如yml，json等
		*/

		viper.SetConfigName(cfgName) //name of config file (without extension)
		viper.AddConfigPath("/etc/tiktok/config")
		viper.AddConfigPath("$HOME/.tiktok/")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("../../config")
	}

	if isRemoteConfig {
		if err := viper.ReadRemoteConfig(); err != nil {
			panic(fmt.Errorf("error reading config: %s", err))
		}
		fmt.Printf("Using Remote Config: '%s'\n", configVar)

		viper.WatchRemoteConfig()

		go config.WatchRemoteConf()

	} else {
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("error reading config: %s", err))
		}
		fmt.Printf("Using configuration file '%s'\n", viper.ConfigFileUsed())

		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			fmt.Printf("Global.Source: '%s'\n", viper.GetString("Global.Source"))
			fmt.Printf("Global.ChangeMe: '%s'\n", viper.GetString("Global.ChangeMe"))
		})

	}

	return config
}
