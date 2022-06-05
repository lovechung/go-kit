package bootstrap

import (
	consulConfig "github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

// NewFileConfigSource 创建一个本地文件配置源
func NewFileConfigSource(filePath string) config.Source {
	return file.NewSource(filePath)
}

// NewRemoteConfigSource 创建一个远程配置源
func NewRemoteConfigSource(configType, configHost, configKey, configToken string) config.Source {
	switch configType {
	case "consul":
		return NewConsulConfigSource(configHost, configKey, configToken)
	}
	return nil
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func NewConsulConfigSource(configHost, configKey, configToken string) config.Source {
	consulClient, err := api.NewClient(&api.Config{
		Address: configHost,
		Token:   configToken,
	})
	if err != nil {
		panic(err)
	}
	consulSource, err := consulConfig.New(consulClient, consulConfig.WithPath(configKey))
	if err != nil {
		panic(err)
	}
	return consulSource
}

// NewConfigProvider 创建一个配置
func NewConfigProvider(configPath, configType, configHost, configToken, configKey string) config.Config {
	var options []config.Option

	// 本地配置
	local := config.WithSource(NewFileConfigSource(configPath))
	options = append(options, local)

	// 远程配置
	remote := NewRemoteConfigSource(configType, configHost, configKey, configToken)
	if remote != nil {
		options = append(options, config.WithSource(remote))
	}

	// 显示定义解析格式为yaml
	decoder := config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
		return yaml.Unmarshal(kv.Value, v)
	})
	options = append(options, decoder)

	return config.New(options...)
}
