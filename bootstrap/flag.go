package bootstrap

import "flag"

type CommandFlags struct {
	Conf        string
	ConfigType  string
	ConfigHost  string
	ConfigToken string
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{
		Conf:        "",
		ConfigType:  "consul",
		ConfigHost:  "",
		ConfigToken: "",
	}
}

func (f *CommandFlags) Init() {
	flag.StringVar(&f.Conf, "conf", "./configs", "config path")
	flag.StringVar(&f.ConfigHost, "config_host", "", "config server host")
	flag.StringVar(&f.ConfigToken, "config_token", "", "config server token")
}
