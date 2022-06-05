package bootstrap

import "flag"

type CommandFlags struct {
	Env         string
	Conf        string
	ConfigType  string
	ConfigHost  string
	ConfigToken string
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{
		Env:         "",
		Conf:        "",
		ConfigHost:  "",
		ConfigToken: "",
	}
}

func (f *CommandFlags) Init() {
	flag.StringVar(&f.Env, "env", "dev", "environment")
	flag.StringVar(&f.Conf, "conf", "./configs", "config path")
	flag.StringVar(&f.ConfigType, "config_type", "", "config server host")
	flag.StringVar(&f.ConfigHost, "config_host", "", "config server host")
	flag.StringVar(&f.ConfigToken, "config_token", "", "config server token")
}
