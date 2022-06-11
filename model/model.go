package model

type StaticServerConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	StaticDir string `json:"static_dir"`
}

type MysqlConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Name      string `json:"db"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parse_time"`
}

type ServerConfig struct {
	Name       string             `json:"name"`
	Tags       map[string]string  `json:"tag"`
	MysqlInfo  MysqlConfig        `json:"mysql"`
	StaticInfo StaticServerConfig `json:"static_server"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
