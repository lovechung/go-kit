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
	flag.StringVar(&f.ConfigType, "config_type", "consul", "config server host")
	flag.StringVar(&f.ConfigHost, "config_host", "139.224.187.162:8500", "config server host")
	flag.StringVar(&f.ConfigToken, "config_token", "f86ff0d5-ee3d-70af-a01a-8e5634bd785c", "config server token")
}
